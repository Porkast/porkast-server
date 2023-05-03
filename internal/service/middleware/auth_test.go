package middleware

import (
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func TestCreateToken(t *testing.T) {
	var (
		ctx        = gctx.New()
		tokenStr   string
		tokenModel TokenModel
		err        error
	)

	tokenModel = TokenModel{
		UserId:         "test-user-id",
		NickName:       "test-nickname",
		Email:          "test@test.com",
		CreateDate:     gtime.Now().String(),
		UpdateDateTime: gtime.Now().String(),
		Role:           "test-roll",
	}

	tokenStr, err = CreateToken(ctx, tokenModel)
	if err != nil {
		t.Fatal("Create token failed : ", err)
	}

	t.Log("Create token success, token is ", tokenStr)
}

func TestParseToken(t *testing.T) {
	var (
		err        error
		tokenStr   = "SYMszb/LzuaMOL9KvTijIofVbg1DL9mcr1daqX9oOMYrNvHHusrI6wiS1vEewIJi7D+UE6M31f63alI14vRDzYoo4hemh0SCjy57SiOld0VVN0k1t7T9phkze+9svSi4pfWGPf/s1Bx4/gVtSmCArLhnS1Z+8GMUlxksICtaT7fqIlageouConpX/yQx8RjhN+a36xWiqxrKi0UzGazuOIPCLXZws3tqia/MFkNIDB3m3/AkBtPru/85m1AHnl86gK32yS7T7JfZxzsNE7hAsg=="
		tokenModel *TokenModel
	)

	tokenModel, err = ParseToken(tokenStr)
    if err != nil {
        t.Fatal("Parse token failed : ", err)
    }

    t.Log("Prase token success : ", gjson.MustEncodeString(tokenModel))
}
