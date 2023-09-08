package user

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/model/entity"
	"guoshao-fm-web/internal/service/internal/dao"
	"guoshao-fm-web/internal/service/middleware"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

func Register(ctx context.Context, userInfo dto.UserInfo) (userInfoResp dto.UserInfo, err error) {
	var (
		userInfoEntity entity.UserInfo
		existEntity    entity.UserInfo
		tokenModel     middleware.TokenModel
		token          string
		cryptoPwd      string
	)
	cryptoPwd, _ = gmd5.Encrypt(userInfo.Password)
	userInfo.Id = guid.S()
	userInfo.Password = cryptoPwd
	userInfo.RegDate = gtime.Now()
	userInfo.UpdateDate = gtime.Now()
	gconv.Struct(userInfo, &userInfoEntity)
	userInfoEntity.Password = cryptoPwd

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
	tokenModel = middleware.TokenModel{
		UserId:         userInfoResp.Id,
		NickName:       userInfoResp.Nickname,
		Email:          userInfoResp.Email,
		Mobile:         userInfoResp.Phone,
		CreateDate:     userInfoResp.RegDate.String(),
		UpdateDateTime: userInfoResp.UpdateDate.String(),
	}
	token, err = middleware.CreateToken(ctx, tokenModel)
	userInfoResp.Token = token

	return
}
