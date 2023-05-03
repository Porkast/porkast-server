package user

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestRegister(t *testing.T) {
	type args struct {
		ctx      context.Context
		userInfo dto.UserInfo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "create user info",
			args: args{
				ctx: gctx.New(),
				userInfo: dto.UserInfo{
					Nickname: "test nickname",
					Email:    "test@test.com",
					Password: "testpassword",
				},
			},
		},
	}
	g.DB().SetDryRun(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserInfoResp, err := Register(tt.args.ctx, tt.args.userInfo)
			if err != nil && err.Error() != "user exist" {
				t.Fatal(err)
			}
			t.Log("create user info success : ", gjson.MustEncodeString(gotUserInfoResp))
		})
	}
}
