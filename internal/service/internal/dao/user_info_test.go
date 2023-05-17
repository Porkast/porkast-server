package dao

import (
	"context"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestGetUserInfoByEmailOrPhone(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
		phone string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get user info by phone",
			args: args{
				ctx:   gctx.New(),
				phone: "18801731480",
			},
			wantErr: false,
		},
		{
			name: "get user info by email",
			args: args{
				ctx:   gctx.New(),
				phone: "chenjunqian0810@foxmail.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetUserInfoByEmailOrPhone(tt.args.ctx, tt.args.email, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfoByEmailOrPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
