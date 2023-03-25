package main

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
	if song.Source == NetEase {
		return NeteaseSongURL(song.ID)
	}
	panic("不受支持的数据源")
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
