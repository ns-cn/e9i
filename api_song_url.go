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

// Generated by https://quicktype.io

type ApiSongURLResult struct {
	Data []ApiSongUrlData `json:"data"`
	Code int64            `json:"code"`
}

type ApiSongUrlData struct {
	ID                     int64                            `json:"id"`
	URL                    string                           `json:"url"`
	Br                     int64                            `json:"br"`
	Size                   int64                            `json:"size"`
	Md5                    string                           `json:"md5"`
	Code                   int64                            `json:"code"`
	Expi                   int64                            `json:"expi"`
	Type                   string                           `json:"type"`
	Gain                   float64                          `json:"gain"`
	Peak                   float64                          `json:"peak"`
	Fee                    int64                            `json:"fee"`
	Uf                     interface{}                      `json:"uf"`
	Payed                  int64                            `json:"payed"`
	Flag                   int64                            `json:"flag"`
	CanExtend              bool                             `json:"canExtend"`
	FreeTrialInfo          interface{}                      `json:"freeTrialInfo"`
	Level                  string                           `json:"level"`
	EncodeType             string                           `json:"encodeType"`
	FreeTrialPrivilege     ApiSongUrlFreeTrialPrivilege     `json:"freeTrialPrivilege"`
	FreeTimeTrialPrivilege ApiSongUrlFreeTimeTrialPrivilege `json:"freeTimeTrialPrivilege"`
	URLSource              int64                            `json:"urlSource"`
	RightSource            int64                            `json:"rightSource"`
	PodcastCtrp            interface{}                      `json:"podcastCtrp"`
	EffectTypes            interface{}                      `json:"effectTypes"`
	Time                   int64                            `json:"time"`
}

type ApiSongUrlFreeTimeTrialPrivilege struct {
	ResConsumable  bool  `json:"resConsumable"`
	UserConsumable bool  `json:"userConsumable"`
	Type           int64 `json:"type"`
	RemainTime     int64 `json:"remainTime"`
}

type ApiSongUrlFreeTrialPrivilege struct {
	ResConsumable      bool        `json:"resConsumable"`
	UserConsumable     bool        `json:"userConsumable"`
	ListenType         interface{} `json:"listenType"`
	CannotListenReason interface{} `json:"cannotListenReason"`
}
