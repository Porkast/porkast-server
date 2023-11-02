package user

import (
	"context"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func GetUserInfoByUserId(ctx context.Context, userId string) (userInfoDto dto.UserInfo, err error) {
	var (
		userInfoEntity entity.UserInfo
	)

	if userId == "" {
		err = gerror.New(gcode.CodeMissingParameter.Message())
		return
	}

	userInfoEntity, err = dao.GetUserInfoByUserId(ctx, userId)
	if err != nil {
		return
	}

	gconv.Struct(userInfoEntity, &userInfoDto)
	userInfoDto.Password = ""

	return
}

func SyncUserInfo(ctx context.Context, userInfoDto dto.UserInfo) (syncedUserInfo dto.UserInfo ,err error) {
	var (
		userInfoEntity entity.UserInfo
	)

	userInfoEntity, _ = dao.GetUserInfoByUserId(ctx, userInfoDto.Id)
	if userInfoEntity.Id == "" {
		userInfoEntity = entity.UserInfo{
			Id:       userInfoDto.Id,
			Nickname: userInfoDto.Nickname,
			Password: userInfoDto.Password,
			Email:    userInfoDto.Email,
			Phone:    userInfoDto.Phone,
			Avatar:   userInfoDto.Avatar,
		}
		userInfoEntity.RegDate = gtime.Now()
		userInfoEntity.UpdateDate = gtime.Now()
		err = dao.CreateUserInfo(ctx, userInfoEntity)
	} else {
		userInfoEntity.UpdateDate = gtime.Now()
		err = dao.UpdateUserInfoByUserId(ctx, userInfoDto.Id, userInfoEntity)
	}

	if err == nil {
		gconv.Struct(userInfoEntity, &syncedUserInfo)
	}

	return
}
