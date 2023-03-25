package main

import (
	"testing"
)

func TestNeteaseSongURL(t *testing.T) {
	url := NeteaseSongURL("167876")
	t.Log(url)
	if url == "" {
		t.Fail()
	}
}
