package user

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func TestLogin(t *testing.T) {
	type args struct {
		ctx      context.Context
		email    string
		phone    string
		password string
	}
	tests := []struct {
		name            string
		args            args
		wantUserInfoDto dto.UserInfo
		wantErr         bool
	}{
		{
			name: "Do Login",
			args: args{
				ctx:      gctx.New(),
				email:    "test@test.com",
				phone:    "",
				password: "",
			},
			wantUserInfoDto: dto.UserInfo{
				Id:         "",
				Username:   "",
				Nickname:   "",
				Token:      "",
				Password:   "",
				Email:      "",
				Phone:      "",
				RegDate:    gtime.New(),
				UpdateDate: gtime.New(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserInfoDto, err := Login(tt.args.ctx, tt.args.email, tt.args.phone, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
            t.Log(gotUserInfoDto)
		})
	}
}
