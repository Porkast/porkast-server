package ctls

import (
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/service/middleware"

	userService "guoshao-fm-web/internal/service/user"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (ctl *controller) LoginTpl(req *ghttp.Request) {

	req.Response.WriteTpl("user/login.html", g.Map{
		consts.APP_NAME_KEY:                       consts.APP_NAME,
		consts.LOGIN_EMAIL_INPUT_HINT:             consts.LOGIN_EMAIL_INPUT_HINT_VALUE,
		consts.LOGIN_EMAIL_INPUT_PLACEHOLDER_HINT: consts.LOGIN_EMAIL_INPUT_HINT_PLACEHOLDER_VALUE,
		consts.LOGIN_PASSWORD_PLACEHOLDER_HINT:    consts.LOGIN_PASSWORD_HINT_PLACEHOLDER_VALUE,
		consts.LOGIN_PASSWORD_HINT:                consts.LOGIN_PASSWORD_HINT_VALUE,
		consts.LOGIN_BTN_TEXT:                     consts.LOGIN_BTN_TEXT_VALUE,
		consts.TO_REGISTER_TEXT:                   consts.TO_REGISTER_TEXT_VALUE,
	})
}

func (ctl *controller) DoLogin(req *ghttp.Request) {
	var (
		err          error
		cryptoPwd    string
		reqData      *LoginReqData
		respUserInfo dto.UserInfo
	)
	if err = req.Parse(&reqData); err != nil {
		middleware.JsonExit(req, 1, err.Error())
	}

	cryptoPwd, _ = gmd5.Encrypt(reqData.Password)
	reqData.Password = cryptoPwd
	respUserInfo, err = userService.Login(req.GetCtx(), reqData.Email, reqData.Phone, reqData.Password)
	if err != nil {
		middleware.JsonExit(req, 1, err.Error(), respUserInfo)
	}
	middleware.JsonExit(req, 0, "register success", respUserInfo)
}

func (ctl *controller) RegisterTpl(req *ghttp.Request) {

	req.Response.WriteTpl("user/register.html", g.Map{
		consts.APP_NAME_KEY:                             consts.APP_NAME,
		consts.REGISTER_EMAIL_INPUT_HINT:                consts.REGISTER_EMAIL_INPUT_HINT_VALUE,
		consts.REGISTER_EMAIL_INPUT_PLACEHOLDER_HINT:    consts.REGISTER_EMAIL_INPUT_HINT_PLACEHOLDER_VALUE,
		consts.REGISTER_PASSWORD_PLACEHOLDER_HINT:       consts.REGISTER_PASSWORD_HINT_PLACEHOLDER_VALUE,
		consts.REGISTER_PASSWORD_HINT:                   consts.REGISTER_PASSWORD_HINT_VALUE,
		consts.REGISTER_PASSWORD_CONFIRM_HINT:           consts.REGISTER_PASSWORD_CONFIRM_VALUE,
		consts.REGISTER_NICKNAME_INPUT_HINT:             consts.REGISTER_NICKNAME_INPUT_HINT_VALUE,
		consts.REGISTER_NICKNAME_INPUT_PLACEHOLDER_HINT: consts.REGISTER_NICKNAME_INPUT_PLACEHOLDER_HINT_VALUE,
		consts.REGISTER_BTN_TEXT:                        consts.REGISTER_BTN_TEXT_VALUE,
	})
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
		Nickname: reqData.Nickname,
		Password: reqData.Password,
		Email:    reqData.Email,
		Phone:    reqData.Phone,
	}

	userInfoDto, err = userService.Register(req.GetCtx(), userInfoDto)
	middleware.JsonExit(req, 0, "register success", userInfoDto)
}
