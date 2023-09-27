package user

import (
	"context"
	"porkast-server/internal/dto"
	"porkast-server/internal/model/entity"
	"porkast-server/internal/service/internal/dao"
	"porkast-server/internal/service/middleware"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func Login(ctx context.Context, email, phone, password string) (userInfoDto dto.UserInfo, err error) {
	var (
		userInfoEntity entity.UserInfo
		tokenModel     middleware.TokenModel
		token          string
	)

	if email == "" {
		userInfoEntity, err = dao.GetUserInfoByPhoneAndPassword(ctx, phone, password)
	} else if phone == "" {
		userInfoEntity, err = dao.GetUserInfoByEmailAndPassword(ctx, email, password)
	}

	if userInfoEntity.Id == "" {
		err = gerror.New("user not exist")
		return userInfoDto, err
	}

	g.Log().Line().Debug(ctx, "query user info : \n", gjson.MustEncodeString(userInfoEntity))
	gconv.Struct(userInfoEntity, &userInfoDto)
	tokenModel = middleware.TokenModel{
		UserId:         userInfoDto.Id,
		NickName:       userInfoDto.Nickname,
		Email:          userInfoDto.Email,
		Mobile:         userInfoDto.Phone,
		CreateDate:     userInfoDto.RegDate.String(),
		UpdateDateTime: userInfoDto.UpdateDate.String(),
	}
	token, err = middleware.CreateToken(ctx, tokenModel)
	userInfoDto.Password = ""
	userInfoDto.Token = token

	return
}
