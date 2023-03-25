package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	songs := Search("海阔天空")
	for _, song := range songs {
		t.Log(song)
	}
	if songs == nil || len(songs) == 0 {
		t.Fail()
	}
}
