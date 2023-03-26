package main

import (
	"fmt"
	"github.com/ns-cn/e9i/command"
	"github.com/ns-cn/e9i/loop"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var playlist = make([]Song, 0)
var current = 0
var currentSong = Song{}
var loopType = loop.LOOP_QUEUE
var actionChan = make(chan int)
var OnProcessing = false
var paused = false

func Loop() {
	for {
		select {
		case action := <-actionChan:
			if action == command.PLAYNEXT {
				if loopType != loop.NONE {
					refreshNextIndex(1)
					go Play(playlist[current])
					RefreshSearchList()
				}
			}
		}
	}
}

func refreshNextIndex(offset int) {
	size := len(playlist)
	if loopType == loop.LOOP_QUEUE {
		current = (current + offset + size) % size
	} else if loopType == loop.LOOP_RANDOM {
		current = rand.Intn(size)
	}
}

func Play(song Song) {
	currentSong = song
	url := song.GetPath()
	fmt.Println(url)
	// Download the MP3 file
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Decode the MP3 file
	streamer, format, err := mp3.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sr := format.SampleRate * 2
	OnProcessing = true
	paused = false
	speaker.Init(sr, sr.N(time.Second/10))
	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	// Play the audio
	done := make(chan bool)
	speaker.Clear()
	speaker.Play(beep.Seq(resampled, beep.Callback(func() {
		actionChan <- command.PLAYNEXT
		OnProcessing = false
		done <- true
	})))
	<-done
}
