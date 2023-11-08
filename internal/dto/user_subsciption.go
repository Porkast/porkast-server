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
	Id            string      `json:"id"            ` //
	UserId        string      `json:"userId"        ` //
	CreateTime    *gtime.Time `json:"createTime"    ` //
	Status        int         `json:"status"        ` //
	Keyword       string      `json:"keyword"       ` //
	OrderByDate   int         `json:"orderByDate"   ` //
	Lang          string      `json:"lang"          ` // feed language
	Country       string      `json:"country"       ` //
	ExcludeFeedId string      `json:"excludeFeedId" ` //
	Source        string      `json:"source"        ` //
	RefId         string      `json:"refId"         ` // subscription type id, like listenlater, playlist id
	Type          string      `json:"type"          ` // suscription type, like searchKeyword, playlist, listenlater
}
