package dto

import "github.com/gogf/gf/v2/os/gtime"

type UserSubscriptionFeedDetailDto struct {
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

type UserSubscriptionDto struct {
	Id            string
	UserId        string
	CreateTime    *gtime.Time
	Status        int
	Keyword       string
	OrderByDate   int
	Lang          string
	Country       string
	ExcludeFeedId string
	Source        string
	RefId         string
	Type          string
}
