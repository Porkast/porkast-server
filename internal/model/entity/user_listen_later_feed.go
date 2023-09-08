package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type UserListenLaterFeed struct {
	Id              string      `json:"id"`              //
	UserId          string      `json:"userId"`          //
	ItemId          string      `json:"itemId"`          //
	ChannelId       string      `json:"channelId"`       //
	RegDate         *gtime.Time `json:"regDate"`         //
	Title           string      `json:"title"`           //
	Link            string      `json:"link"`            //
	PubDate         *gtime.Time `json:"pubDate"`         //
	Author          string      `json:"author"`          //
	InputDate       *gtime.Time `json:"inputDate"`       //
	ImageUrl        string      `json:"imageUrl"`        //
	EnclosureUrl    string      `json:"enclosureUrl"`    //
	EnclosureType   string      `json:"enclosureType"`   //
	EnclosureLength string      `json:"enclosureLength"` //
	Duration        string      `json:"duration"`        //
	Episode         string      `json:"episode"`         //
	Explicit        string      `json:"explicit"`        //
	Season          string      `json:"season"`          //
	EpisodeType     string      `json:"episodeType"`     //
	Description     string      `json:"description"`     //
	ChannelImageUrl string      `json:"channelImageUrl"` //
	ChannelTitle    string      `json:"channelTitle"`    //
	FeedLink        string      `json:"feedLink"`        //
}
