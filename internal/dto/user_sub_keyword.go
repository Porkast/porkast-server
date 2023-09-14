package dto

import "github.com/gogf/gf/v2/os/gtime"

type UserSubKeywordFeedDetailDto struct {
	Id              string
	UserId          string
	Keyword         string
	OrderByDate     int
	ItemId          string
	ChannelId       string
	CreateDate      string
	Title           string
	Link            string
	PubDate         string
	Author          string
	InputDate       string
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
	TextDescription string
	ChannelImageUrl string
	ChannelTitle    string
	ChannelAuthor   string
	FeedLink        string
	Count           int
	HasThumbnail    bool
}

type UserSubKeywordDto struct {
	Id          string
	UserId      string
	CreaterName string
	Keyword     string
	OrderByDate int
	CreateTime  *gtime.Time
	Lang        string
	Status      int
}
