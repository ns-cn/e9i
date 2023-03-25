package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// apiSearch 搜索的接口URL
	apiSearh = "https://api.tangyujun.com/search?keywords=%s"
)

// Search 搜索关键词
func Search(keyword string) []Song {
	resp, err := http.Get(fmt.Sprintf(apiSearh, keyword))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result ApiSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result.GetSongs()
}
