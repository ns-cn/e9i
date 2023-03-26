package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/ns-cn/e9i/source"
	"os"
)

func main() {
	fmt.Println("hello")
	playlist = append(playlist, Song{Source: source.NetEase, Name: "有何不可", ID: "167876"})
	playlist = append(playlist, Song{Source: source.NetEase, Name: "有何不可", ID: "1960052435"})
	go Loop()
	createApp()
	RefreshSearchList()
	application.Resize(fyne.NewSize(400, 600))
	application.ShowAndRun()
}

func init() {
	err := os.Setenv("FYNE_FONT", "演示春风楷.ttf")
	if err != nil {
		return
	}
}
