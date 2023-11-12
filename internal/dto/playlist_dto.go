package dto

import "github.com/gogf/gf/v2/os/gtime"

type UserPlaylistDto struct {
	Id             string
	PlaylistName   string
	Description    []byte
	UserId         string
	RegDate        *gtime.Time
	Status         int
	CreatorId      string
	OrigPlaylistId string
}
