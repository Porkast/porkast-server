package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func SetI18nLang(req *ghttp.Request) {
	g.I18n().SetLanguage("zh-CN")
	req.Middleware.Next()
}
