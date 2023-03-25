package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func play(url string) {
	// Download the MP3 file
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
	defer resp.Body.Close()

	// Decode the MP3 file
	mp3Decoder, err := mp3.NewDecoder(resp.Body)
	if resp.StatusCode != 200 || err != nil {
		log.Fatal(err)
	}
	// Usually 44100 or 48000. Other values might cause distortions in Oto
	samplingRate := 44100

	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
	// 1 is mono sound, and 2 is stereo (most speakers are stereo).
	numOfChannels := 2

	// Bytes used by a channel to represent one sample. Either 1 or 2 (usually 2).
	audioBitDepth := 2

	// Remember that you should **not** create more than one context
	otoCtx, err := oto.NewContext(samplingRate, numOfChannels, audioBitDepth, int(mp3Decoder.Length()))
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	// Create a new 'player' that will handle our sound. Paused by default.
	player := otoCtx.NewPlayer()
	// Play the audio
	buf := make([]byte, 32)
	for {
		_, err := mp3Decoder.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if _, err := player.Write(buf); err != nil {
			log.Fatal(err)
		}
	}
}
