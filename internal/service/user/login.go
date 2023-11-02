package user

import (
	"context"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao"

	"github.com/gogf/gf/v2/os/gtime"
)

func Login(ctx context.Context, userId string) (userInfoDto dto.UserInfo, err error) {
	var (
		userInfoEntity entity.UserInfo
	)

	userInfoEntity, err = dao.GetUserInfoByUserId(ctx, userId)
	if err != nil {
		return
	}

	userInfoEntity.UpdateDate = gtime.Now()

	err = dao.UpdateUserInfoByUserId(ctx, userId, userInfoEntity)

	return
}
