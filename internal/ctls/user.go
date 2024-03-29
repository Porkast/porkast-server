package ctls

import (
	"porkast-server/internal/consts"
	"porkast-server/internal/dto"
	"porkast-server/internal/service/middleware"

	userService "porkast-server/internal/service/user"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) LoginTpl(req *ghttp.Request) {

	var tplMap = consts.GetCommonTplMap(req.GetCtx())
	tplMap[consts.LOGIN_EMAIL_INPUT_HINT] = consts.LOGIN_EMAIL_INPUT_HINT_VALUE
	tplMap[consts.LOGIN_EMAIL_INPUT_PLACEHOLDER_HINT] = consts.LOGIN_EMAIL_INPUT_HINT_PLACEHOLDER_VALUE
	tplMap[consts.LOGIN_PASSWORD_PLACEHOLDER_HINT] = consts.LOGIN_PASSWORD_HINT_PLACEHOLDER_VALUE
	tplMap[consts.LOGIN_PASSWORD_HINT] = consts.LOGIN_PASSWORD_HINT_VALUE
	tplMap[consts.LOGIN_BTN_TEXT] = consts.LOGIN_BTN_TEXT_VALUE
	tplMap[consts.TO_REGISTER_TEXT] = consts.TO_REGISTER_TEXT_VALUE
	req.Response.WriteTpl("user/login.html", tplMap)
}

func (ctl *controller) RegisterTpl(req *ghttp.Request) {

	var tplMap = consts.GetCommonTplMap(req.GetCtx())
	tplMap[consts.REGISTER_EMAIL_INPUT_HINT] = consts.REGISTER_EMAIL_INPUT_HINT_VALUE
	tplMap[consts.REGISTER_EMAIL_INPUT_PLACEHOLDER_HINT] = consts.REGISTER_EMAIL_INPUT_HINT_PLACEHOLDER_VALUE
	tplMap[consts.REGISTER_PASSWORD_PLACEHOLDER_HINT] = consts.REGISTER_PASSWORD_HINT_PLACEHOLDER_VALUE
	tplMap[consts.REGISTER_PASSWORD_HINT] = consts.REGISTER_PASSWORD_HINT_VALUE
	tplMap[consts.REGISTER_PASSWORD_CONFIRM_HINT] = consts.REGISTER_PASSWORD_CONFIRM_VALUE
	tplMap[consts.REGISTER_NICKNAME_INPUT_HINT] = consts.REGISTER_NICKNAME_INPUT_HINT_VALUE
	tplMap[consts.REGISTER_NICKNAME_INPUT_PLACEHOLDER_HINT] = consts.REGISTER_NICKNAME_INPUT_PLACEHOLDER_HINT_VALUE
	tplMap[consts.REGISTER_BTN_TEXT] = consts.REGISTER_BTN_TEXT_VALUE
	req.Response.WriteTpl("user/register.html", tplMap)
}

func (ctl *controller) UserInfoTpl(req *ghttp.Request) {

	var tplMap = consts.GetCommonTplMap(req.GetCtx())
	tplMap[consts.NICKANME_TEXT] = g.I18n().T(req.GetCtx(), `{#nickname}`)
	tplMap[consts.ACCOUNT_TEXT] = g.I18n().T(req.GetCtx(), `{#account}`)
	tplMap[consts.REG_DATE_TEXT] = g.I18n().T(req.GetCtx(), `{#reg_date}`)
	req.Response.WriteTpl("user/account_info.html", tplMap)
}


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
		middleware.JsonExit(req, 1, "sync user info failed", nil)
	}
	middleware.JsonExit(req, 0, g.I18n().T(req.GetCtx(), `{#register_sucess}`), syncedUserInfo)
}

func (ctl *controller) GetUserInfo(req *ghttp.Request) {
	var (
		err         error
		userInfoDto dto.UserInfo
		userId      string
	)

	userId = req.Get("userId").String()
	userInfoDto, err = userService.GetUserInfoByUserId(req.GetCtx(), userId)

	if err != nil {
		g.Log().Line().Error(req.GetCtx(), err)
		middleware.JsonExit(req, 1, err.Error(), nil)
	}
	middleware.JsonExit(req, 0, "", userInfoDto)
}
