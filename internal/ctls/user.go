package ctls

import (
	"porkast-server/internal/dto"
	"porkast-server/internal/service/middleware"

	userService "porkast-server/internal/service/user"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) DoLogin(req *ghttp.Request) {
	var (
		err          error
		reqData      *LoginReqData
		respUserInfo dto.UserInfo
	)
	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	respUserInfo, err = userService.Login(req.GetCtx(), reqData.UserId)
	if err != nil {
		g.Log().Line().Error(req.GetCtx(), err)
		middleware.JsonExit(req, 1, g.I18n().T(req.GetCtx(), `{#username_or_password_not_right}`), nil)
	}
	g.Log().Line().Debug(req.GetCtx(), "do login success :\n", gjson.MustEncodeString(respUserInfo))
	middleware.JsonExit(req, 0, g.I18n().T(req.GetCtx(), `{#login_sucess}`), respUserInfo)
}

func (ctl *controller) DoRegister(req *ghttp.Request) {
	var (
		err         error
		reqData     *RegisterReqData
		userInfoDto dto.UserInfo
	)

	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	userInfoDto = dto.UserInfo{
		Id:       reqData.Id,
		Nickname: reqData.Nickname,
		Password: reqData.Password,
		Email:    reqData.Email,
		Phone:    reqData.Phone,
	}

	userInfoDto, err = userService.Register(req.GetCtx(), userInfoDto)
	if err != nil {
		g.Log().Line().Error(req.GetCtx(), err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}
	middleware.JsonExit(req, 0, g.I18n().T(req.GetCtx(), `{#register_sucess}`), userInfoDto)
}

func (ctl *controller) SyncUserInfo(req *ghttp.Request) {
	var (
		ctx         = req.GetCtx()
		err         error
		reqData     *SyncUserInfoReqData
		userInfoDto dto.UserInfo
	)

	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	userInfoDto = dto.UserInfo{
		Id:       reqData.UserId,
		Nickname: reqData.Nickname,
		Password: reqData.Password,
		Email:    reqData.Email,
		Phone:    reqData.Phone,
		Avatar:   reqData.Avatar,
	}

	syncedUserInfo, err := userService.SyncUserInfo(ctx, userInfoDto)
	if err != nil {
		g.Log().Line().Error(req.GetCtx(), err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}
	middleware.JsonExit(req, 0, g.I18n().T(req.GetCtx(), `{#register_sucess}`), syncedUserInfo)
}

func (ctl *controller) GetUserInfo(req *ghttp.Request) {
	var (
		err         error
		reqData     *GetUserInfoReqData
		userInfoDto dto.UserInfo
	)

	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	userInfoDto, err = userService.GetUserInfoByUserId(req.GetCtx(), reqData.UserId)

	if err != nil {
		g.Log().Line().Error(req.GetCtx(), err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}
	middleware.JsonExit(req, 0, "", userInfoDto)
}
