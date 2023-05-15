package user

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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
