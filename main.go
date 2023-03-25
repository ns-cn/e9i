package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	song := Song{Source: NetEase, ID: "167876"}
	play(song.GetPath())
}
