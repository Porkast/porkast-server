package user

import (
	"context"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func Register(ctx context.Context, userInfo dto.UserInfo) (userInfoResp dto.UserInfo, err error) {
	var (
		userInfoEntity entity.UserInfo
		existEntity    entity.UserInfo
	)
	userInfo.RegDate = gtime.Now()
	userInfo.UpdateDate = gtime.Now()
	gconv.Struct(userInfo, &userInfoEntity)

	existEntity, _ = dao.GetUserInfoByEmailOrPhone(ctx, userInfo.Email, userInfo.Phone)
	if existEntity.Id != "" {
		return userInfoResp, gerror.New(g.I18n().T(ctx, `{#user_already_exist}`))
	}

	err = dao.CreateUserInfo(ctx, userInfoEntity)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return dto.UserInfo{}, err
	}
	userInfoResp = userInfo
	userInfoResp.Password = ""

	return
}
