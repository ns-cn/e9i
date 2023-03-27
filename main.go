package main

import (
	"fyne.io/fyne/v2"
	"github.com/ns-cn/e9i/source"
)

func main() {
	playlist = append(playlist, Song{Source: source.NetEase, Name: "You Are Wind", Artists: []Artist{{Name: "阿摩司"}, {Name: "第六大街"}}, ID: "28993079"})
	playlist = append(playlist, Song{Source: source.NetEase, Name: "110819", Artists: []Artist{{Name: "宮内優里"}}, ID: "33166200"})
	playlist = append(playlist, Song{Source: source.NetEase, Name: "111004", Artists: []Artist{{Name: "宮内優里"}}, ID: "33166218"})
	go Loop()
	createApp()
	RefreshSearchList()
	windows.Resize(fyne.NewSize(400, 600))
	windows.ShowAndRun()
}
