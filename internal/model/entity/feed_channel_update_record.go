// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2023-08-23 21:27:04
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FeedChannelUpdateRecord is the golang structure for table feed_channel_update_record.
type FeedChannelUpdateRecord struct {
	ChannelId  string      `json:"channelId"  ` //
	FuncName   string      `json:"funcName"   ` //
	UpdateTime *gtime.Time `json:"updateTime" ` //
}
