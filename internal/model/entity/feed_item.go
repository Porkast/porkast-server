// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2023-05-02 05:52:15
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FeedItem is the golang structure for table feed_item.
type FeedItem struct {
	Id              string      `json:"id"              ` //
	ChannelId       string      `json:"channelId"       ` //
	Title           string      `json:"title"           ` //
	Link            string      `json:"link"            ` //
	PubDate         *gtime.Time `json:"pubDate"         ` //
	Author          string      `json:"author"          ` //
	InputDate       *gtime.Time `json:"inputDate"       ` //
	ImageUrl        string      `json:"imageUrl"        ` //
	EnclosureUrl    string      `json:"enclosureUrl"    ` //
	EnclosureType   string      `json:"enclosureType"   ` //
	EnclosureLength string      `json:"enclosureLength" ` //
	Duration        string      `json:"duration"        ` //
	Episode         string      `json:"episode"         ` //
	Explicit        string      `json:"explicit"        ` //
	Season          string      `json:"season"          ` //
	EpisodeType     string      `json:"episodeType"     ` //
	Description     string      `json:"description"     ` //
}
