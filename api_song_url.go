package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// apiSongURL 歌曲URL的获取接口
	apiSongURL = "https://api.tangyujun.com/song/url?id=%s&realIP=116.25.146.177"
)

// NeteaseSongURL 获取网易云指定歌曲的url
func NeteaseSongURL(ID string) string {
	resp, err := http.Get(fmt.Sprintf(apiSongURL, ID))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result ApiSongURLResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	if result.Data != nil && len(result.Data) > 0 {
		return result.Data[0].URL
	}
	return ""
}
