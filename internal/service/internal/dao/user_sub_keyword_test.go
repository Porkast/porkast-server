package dao

import (
	"context"
	"guoshao-fm-web/internal/consts"
	"guoshao-fm-web/internal/model/entity"
	"reflect"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

func TestGetUserSubKeywordByUserIdAndKeyword(t *testing.T) {
	type args struct {
		ctx     context.Context
		userId  string
		keyword string
	}
	tests := []struct {
		name             string
		args             args
		wantResultEntity entity.UserSubKeyword
		wantErr          bool
	}{
		{
			name: "Get UserSubKeyword by userId and keyword",
			args: args{
				ctx:     gctx.New(),
				userId:  "1t5z27w7h00csfdx7cluc20100do2yyq",
				keyword: `游戏`,
			},
			wantResultEntity: entity.UserSubKeyword{
				Id:          "12tck3g1whucu2v3c0m4dlk100wp0nt2",
				UserId:      "1t5z27w7h00csfdx7cluc20100do2yyq",
				Keyword:     "游戏",
				OrderByDate: 1,
				CreateTime:  gtime.NewFromStr("2023-07-15 15:47:15"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResultEntity, err := GetUserSubKeywordByUserIdAndKeyword(tt.args.ctx, tt.args.userId, tt.args.keyword)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserSubKeywordByUserIdAndKeyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResultEntity, tt.wantResultEntity) {
				t.Errorf("GetUserSubKeywordByUserIdAndKeyword() = %v, want %v", gotResultEntity, tt.wantResultEntity)
			}
		})
	}
}

func TestCreateUserSubKeyword(t *testing.T) {
	type args struct {
		ctx       context.Context
		newEntity entity.UserSubKeyword
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create user sub keyword",
			args: args{
				ctx: gctx.New(),
				newEntity: entity.UserSubKeyword{
					Id:          guid.S(),
					UserId:      "1t5z27w7h00csfdx7cluc20100do2yyq",
					Keyword:     "游戏",
					OrderByDate: 1,
					CreateTime:  gtime.New(),
				},
			},
			wantErr: false,
		},
	}
	g.DB().SetDryRun(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUserSubKeyword(tt.args.ctx, tt.args.newEntity); (err != nil) != tt.wantErr {
				if err.Error() != consts.DB_DATA_ALREADY_EXIST {
					t.Errorf("CreateUserSubKeyword() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
