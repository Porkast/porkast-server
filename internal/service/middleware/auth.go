package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/gogf/gf/v2/crypto/gaes"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type TokenModel struct {
	UserId         string `json:"userId"`
	NickName       string `json:"nickname"`
	Email          string `json:"email"`
	Mobile         string `json:"mobile"`
	CreateDate     string `json:"createDate"`
	UpdateDateTime string `json:"updateDateTime"`
	Role           string `json:"role"`
	Token          string `json:"token"`
}

var privateKey = "guoshao-t01-12-1"

func CreateToken(ctx context.Context, tokenData TokenModel) (token string, err error) {

	if jsonToken, err := json.Marshal(tokenData); err != nil {
		g.Log().Line().Error(ctx, "decode token to json error : ", err)
		return "", err
	} else {
		if token, err := gaes.Encrypt(jsonToken, []byte(privateKey)); err != nil {
			g.Log().Line().Error(ctx, "aes encrypt string error: ", err)
			return "", err
		} else {
			encodeToken := gbase64.EncodeToString(token)
			return encodeToken, nil
		}
	}
}

func ParseToken(tokenString string) (*TokenModel, error) {
	decodeToken, _ := gbase64.Decode([]byte(tokenString))
	decResult, err := gaes.Decrypt(decodeToken, []byte(privateKey))
	ctx := context.Background()
	if err != nil {
		g.Log().Line().Error(ctx, "token decrypt error : ", tokenString)
		return nil, err
	}
	tokenModel := new(TokenModel)
	if err := json.Unmarshal(decResult, tokenModel); err != nil {
		g.Log().Line().Error(ctx, "token string decode to json error , token: ", decResult, " ,error : ", err)
		return nil, err
	}
	return tokenModel, nil
}

func validateToken(cxt context.Context, authString string) (token, uid string, err error) {
	authorizationArray := strings.Split(authString, "@@")
	if len(authorizationArray) < 2 {
		g.Log().Line().Error(cxt, "Token or uid is null")
		return "", "", errors.New("Token or uid is null")
	}
	token = authorizationArray[0]
	uid = authorizationArray[1]
	if len(token) < 0 || len(uid) < 0 {
		g.Log().Line().Error(cxt, "AuthToken or uid is null")
		return "", "", errors.New("AuthToken or uid is null")
	}
	return token, uid, nil
}

func AuthToken(req *ghttp.Request) {
	authorization := req.GetHeader("Authorization")
	token, uid, err := validateToken(req.GetCtx(), authorization)
	if err != nil {
		JsonExit(req, 1, err.Error())
	}

	tokenModel, err := ParseToken(token)
	if err != nil {
		JsonExit(req, 1, "AuthToken Invalid")
	}

	if tokenModel.UserId != uid {
		g.Log().Line().Info(req.Request.Context(), "token invalid tokenModel : ", tokenModel, " ,uid : ", uid)
		JsonExit(req, 1, "AuthToken Invalid")
	}
	req.Middleware.Next()
}
