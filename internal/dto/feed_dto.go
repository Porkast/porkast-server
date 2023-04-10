package dto

import "github.com/gogf/gf/v2/os/gtime"

type FeedItem struct {
	Id              string
	ChannelId       string
	Title           string
	Link            string
	PubDate         *gtime.Time
	Author          string
	InputDate       *gtime.Time
	ImageUrl        string
	EnclosureUrl    string
	EnclosureType   string
	EnclosureLength string
	Duration        string
	Episode         string
	Explicit        string
	Season          string
	EpisodeType     string
	Description     string
	ChannelImageUrl string
	ChannelTitle    string
	SourceLink      string
	Count           int
	HasThumbnail    bool
}

type FeedChannel struct {
	Id          string
	Title       string
	ChannelDesc string
	ImageUrl    string
	Link        string
	FeedLink    string
	Launguage   string
	FeedType    string
	Categories  string
	Author      string
	OwnerName   string
	OwnerEmail  string
	Items       []FeedItem
	Count       int
}