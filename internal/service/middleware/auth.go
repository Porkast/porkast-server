package middleware

import (
	"fmt"
	"os"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
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

var supabaseJWTSecret = "hwfkEIMPtn0gREZUEcV2ZeksWcI/IClvR8TesHgRWkQbnTKTS+VnA0nvczDCjnZDIx5DCJYgxYrzrUy0OrWoSw=="

func AuthToken(req *ghttp.Request) {
	if os.Getenv("env") == "dev" {
		req.Middleware.Next()
	} else {
		tokenString := req.GetHeader("Authorization")
		err := VerifyJWTToken(tokenString)

		if err != nil {
			JsonExit(req, 1, err.Error(), nil)
		}

		req.Middleware.Next()
	}
}

func VerifyJWTToken(tokenString string) (err error) {

	// verify JWT token
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		var hmacSampleSecret = []byte(supabaseJWTSecret)
		return hmacSampleSecret, nil
	})

	fmt.Println(gjson.MustEncodeString(token))
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		err = gerror.New("invalid token")
	}
	return
}
