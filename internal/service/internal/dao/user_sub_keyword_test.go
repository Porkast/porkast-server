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
				Id:          "12tck3g238kcu9gpmnqte0o100dr5yq8",
				UserId:      "1t5z27w7h00csfdx7cluc20100do2yyq",
				Keyword:     "游戏",
				OrderByDate: 0,
				Lang:        "zh",
				CreateTime:  gtime.NewFromStr("2023-07-23 09:59:45"),
				Status:      1,
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

func TestGetUserSubKeywordListByUserIdAndKeyword(t *testing.T) {
	type args struct {
		ctx     context.Context
		userId  string
		keyword string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Get sub feed list by user id and keyword",
			args: args{
				userId:  "1t5z27w7h00csfdx7cluc20100do2yyq",
				keyword: "游戏",
			},
			wantErr: false,
		},
	}
	g.DB().SetDryRun(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDtos, err := GetUserSubKeywordListByUserIdAndKeyword(tt.args.ctx, tt.args.userId, tt.args.keyword)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserSubKeywordListByUserIdAndKeyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(gotDtos) == 0 {
				t.Log("GetUserSubKeywordListByUserIdAndKeyword() sub keyword list is empty")
			}
		})
	}
}

func TestDoSubKeywordByUserIdAndKeyword(t *testing.T) {
	type args struct {
		ctx             context.Context
		newUSKEntity    entity.UserSubKeyword
		newKSEntityList []entity.KeywordSubscription
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sub keyword by user id and keyword",
			args: args{
				ctx: gctx.New(),
				newUSKEntity: entity.UserSubKeyword{
					Id:          guid.S(),
					UserId:      "1t5z27w7h00csfdx7cluc20100do2yyq",
					Keyword:     "游戏",
					OrderByDate: 0,
					CreateTime:  gtime.Now(),
					Lang:        "zh",
					Status:      1,
				},
				newKSEntityList: []entity.KeywordSubscription{
					{
						Id:            guid.S(),
						Keyword:       "游戏",
						FeedChannelId: "o66b2cv6l9qr",
						FeedItemId:    "akeq65sh2r9m4",
						CreateTime:    gtime.Now(),
					},
					{
						Id:            guid.S(),
						Keyword:       "游戏",
						FeedChannelId: "o66b2cv6l9qr",
						FeedItemId:    "4pldl03li7vco",
						CreateTime:    gtime.Now(),
					},
					{
						Id:            guid.S(),
						Keyword:       "游戏",
						FeedChannelId: "atkd2vcbr952a",
						FeedItemId:    "dbk846o14cenn",
						CreateTime:    gtime.Now(),
					},
					{
						Id:            guid.S(),
						Keyword:       "游戏",
						FeedChannelId: "o66b2cv6l9qr",
						FeedItemId:    "2ni4f5ara7h8u",
						CreateTime:    gtime.Now(),
					},
					{
						Id:            guid.S(),
						Keyword:       "游戏",
						FeedChannelId: "7veqaesn5q8o1",
						FeedItemId:    "8qvhn6io2618v",
						CreateTime:    gtime.Now(),
					},
					{
						Id:            guid.S(),
						Keyword:       "游戏",
						FeedChannelId: "d87q7kjoeg8nd",
						FeedItemId:    "ubm550p9m75k",
						CreateTime:    gtime.Now(),
					},
				},
			},
			wantErr: false,
		},
	}
	g.DB().SetDryRun(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DoSubKeywordByUserIdAndKeyword(tt.args.ctx, tt.args.newUSKEntity, tt.args.newKSEntityList); (err != nil) != tt.wantErr {
				if err.Error() != consts.DB_DATA_ALREADY_EXIST {
					t.Errorf("DoSubKeywordByUserIdAndKeyword() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
