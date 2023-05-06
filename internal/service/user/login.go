package user

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

func Login(ctx context.Context, email, phone, password string) (userInfoDto dto.UserInfo, err error) {
	var (
		userInfoEntity entity.UserInfo
	)

	if email == "" {
		userInfoEntity, err = dao.GetUserInfoByPhoneAndPassword(ctx, phone, password)
	} else if phone == "" {
		userInfoEntity, err = dao.GetUserInfoByEmailAndPassword(ctx, email, password)
	}

	gconv.Struct(userInfoEntity, userInfoDto)
    userInfoDto.Password = ""

	return
}
