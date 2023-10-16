// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2023-10-16 22:30:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserSubKeyword is the golang structure for table user_sub_keyword.
type UserSubKeyword struct {
	Id            string      `json:"id"            ` //
	UserId        string      `json:"userId"        ` //
	Keyword       string      `json:"keyword"       ` //
	OrderByDate   int         `json:"orderByDate"   ` //
	CreateTime    *gtime.Time `json:"createTime"    ` //
	Lang          string      `json:"lang"          ` // feed language
	Country       string      `json:"country"       ` //
	ExcludeFeedId string      `json:"excludeFeedId" ` //
	Status        int         `json:"status"        ` //
}
