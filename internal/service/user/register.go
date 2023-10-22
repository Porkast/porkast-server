package user

import (
	"context"
	"porkast-server/internal/dto"
)

func Register(ctx context.Context, userInfo dto.UserInfo) (userInfoResp dto.UserInfo, err error) {
	// var (
	// 	userInfoEntity entity.UserInfo
	// 	existEntity    entity.UserInfo
	// 	tokenModel     middleware.TokenModel
	// 	token          string
	// 	cryptoPwd      string
	// )
	// cryptoPwd, _ = gmd5.Encrypt(userInfo.Password)
	// userInfo.Id = guid.S()
	// userInfo.Password = cryptoPwd
	// userInfo.RegDate = gtime.Now()
	// userInfo.UpdateDate = gtime.Now()
	// gconv.Struct(userInfo, &userInfoEntity)
	// userInfoEntity.Password = cryptoPwd

	// existEntity, _ = dao.GetUserInfoByEmailOrPhone(ctx, userInfo.Email, userInfo.Phone)
	// if existEntity.Id != "" {
	// 	return userInfoResp, gerror.New(g.I18n().T(ctx, `{#user_already_exist}`))
	// }

	// err = dao.CreateUserInfo(ctx, userInfoEntity)
	// if err != nil {
	// 	g.Log().Line().Error(ctx, err)
	// 	return dto.UserInfo{}, err
	// }
	// userInfoResp = userInfo
	// userInfoResp.Password = ""
	// tokenModel = middleware.TokenModel{
	// 	UserId:         userInfoResp.Id,
	// 	NickName:       userInfoResp.Nickname,
	// 	Email:          userInfoResp.Email,
	// 	Mobile:         userInfoResp.Phone,
	// 	CreateDate:     userInfoResp.RegDate.String(),
	// 	UpdateDateTime: userInfoResp.UpdateDate.String(),
	// }
	// token, err = middleware.CreateToken(ctx, tokenModel)
	// userInfoResp.Token = token

	return
}
