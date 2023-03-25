package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	songs := Search("海阔天空")
	t.Log(songs)
	if songs == nil || len(songs) == 0 {
		t.Fail()
	}
}
