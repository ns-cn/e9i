package main

import "github.com/ns-cn/e9i/source"

const (
	apiNeteaseURL = "https://music.163.com/song/media/outer/url?id=%s.mp3"
)

// Song 歌曲
type Song struct {
	// Source 数据来源
	Source string
	// ID ID
	ID string
	// Name 歌曲名称
	Name string
	// Artists 歌手
	Artists []Artist
	// Album 专辑
	Album Album
}

// GetPath 获取可用于播放的URL
func (song Song) GetPath() string {
	if song.Source == source.NetEase {
		return NeteaseSongURL(song.ID)
	}
	panic("不受支持的数据源")
}

func (song Song) GetArtistDisplay() string {
	result := ""
	artistSize := len(song.Artists)
	if artistSize > 0 {
		for _, artist := range song.Artists {
			result += artist.Name + ","
		}
	}
	if len(result) > 1 {
		return result[:len(result)-1]
	} else {
		return ""
	}
}

// Artist 歌手模型
type Artist struct {
	Source string
	ID     string
	Name   string
}

// Album 专辑
type Album struct {
	Source string
	ID     string
	Name   string
}
