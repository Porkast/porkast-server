package dto

import "github.com/gogf/gf/v2/os/gtime"

type UserListenLater struct {
	Id        string      `json:"id"        ` //
	UserId    string      `json:"userId"    ` //
	ItemId    string      `json:"itemId"    ` //
	ChannelId string      `json:"channelId" ` //
	RegDate   *gtime.Time `json:"regDate"   ` //
	ItemInfo  FeedItem    `json:"itemInfo"  ` //
}
