package dto

import "github.com/gogf/gf/v2/os/gtime"

type UserPlaylistDto struct {
	Id             string
	PlaylistName   string
	Description    string
	UserId         string
	RegDate        *gtime.Time
	Status         int
	CreatorId      string
	OrigPlaylistId string
	Count          int
	UserInfo	   UserInfo
}

type UserPlaylistItemDto struct {
	Id                    string
	GUID                  string
	ChannelId             string
	FeedId                string
	Title                 string
	HighlightTitle        string
	Link                  string
	PubDate               string
	Author                string
	InputDate             *gtime.Time
	ImageUrl              string
	EnclosureUrl          string
	EnclosureType         string
	EnclosureLength       string
	Duration              string
	Episode               string
	Explicit              string
	Season                string
	EpisodeType           string
	Description           string
	TextDescription       string
	ChannelImageUrl       string
	ChannelTitle          string
	HighlightChannelTitle string
	FeedLink              string
	Count                 int
	TookTime              float64
	HasThumbnail          bool
	Source                string
	ExcludeFeedId         string
	Country               string
	RegDate               string
}
