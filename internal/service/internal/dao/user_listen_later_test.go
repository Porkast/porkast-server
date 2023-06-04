package dao

import (
	"context"
	"guoshao-fm-web/internal/model/entity"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

func TestGetListenLaterByUserIdAndFeedId(t *testing.T) {
	type args struct {
		ctx       context.Context
		userId    string
		channelId string
		itemId    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get listen later by user_id and feed_id",
			args: args{
				userId:    "1t5z27w7h00csfdx7cluc20100do2yyq",
				channelId: "8k4vnjjtcmjqi",
				itemId:    "100r8jf600hcm",
			},
			wantErr: false,
		},
		{
			name: "get listen later without user_id",
			args: args{
				channelId: "8k4vnjjtcmjqi",
				itemId:    "100r8jf600hcm",
			},
			wantErr: true,
		},
		{
			name: "get listen later without channel_id",
			args: args{
				userId: "1t5z27w7h00csfdx7cluc20100do2yyq",
				itemId: "100r8jf600hcm",
			},
			wantErr: true,
		},
		{
			name: "get listen later without item_id",
			args: args{
				userId:    "1t5z27w7h00csfdx7cluc20100do2yyq",
				channelId: "8k4vnjjtcmjqi",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				err error
			)
			if _, err = GetListenLaterByUserIdAndFeedId(tt.args.ctx, tt.args.userId, tt.args.channelId, tt.args.itemId); (err != nil) != tt.wantErr {
				t.Errorf("GetListenLaterByUserIdAndFeedId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateListenLaterByUserIdAndFeedId(t *testing.T) {
	type args struct {
		ctx       context.Context
		newEntity entity.UserListenLater
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Create Listen Later By UserId And FeedId",
			args: args{
				newEntity: entity.UserListenLater{
					Id:        guid.S(),
					UserId:    "1t5z27w7h00csfdx7cluc20100do2yyq",
					ChannelId: "8k4vnjjtcmjqi",
					ItemId:    "100r8jf600hcm",
					RegDate:   gtime.Now(),
				},
			},
			wantErr: true,
		},
		{
			name: "Create Listen Later without RegDate",
			args: args{
				newEntity: entity.UserListenLater{
					Id:        guid.S(),
					UserId:    "1t5z27w7h00csfdx7cluc20100do2yyq",
					ChannelId: "8k4vnjjtcmjqi",
					ItemId:    "100r8jf600hcm",
				},
			},
			wantErr: true,
		},
		{
			name: "Create Listen Later without user_id",
			args: args{
				newEntity: entity.UserListenLater{
					Id:        guid.S(),
					ChannelId: "8k4vnjjtcmjqi",
					ItemId:    "100r8jf600hcm",
					RegDate:   gtime.New(),
				},
			},
			wantErr: true,
		},
	}

	g.DB().SetDryRun(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateListenLaterByUserIdAndFeedId(tt.args.ctx, tt.args.newEntity); (err != nil) != tt.wantErr {
				t.Errorf("CreateListenLaterByUserIdAndFeedId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetListenLaterListByUserId(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get listen later list",
			args: args{
				ctx:    gctx.New(),
				userId: "1t5z27w7h00csfdx7cluc20100do2yyq",
				offset: 0,
				limit:  10,
			},
			wantErr: false,
		},
		{
			name: "get listen later list without user id",
			args: args{
				ctx:    gctx.New(),
				offset: 0,
				limit:  10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultEntity, err := GetListenLaterListByUserId(tt.args.ctx, tt.args.userId, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListenLaterListByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if !tt.wantErr {
				if len(resultEntity) == 0 {
					t.Fatal("get listen later feed list is empty")
				}
			}

		})
	}
}

func TestGetTotalListenLaterCountByUserId(t *testing.T) {

	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get user listen later total count by user id",
			args: args{
				ctx:    gctx.New(),
				userId: "1t5z27w7h00csfdx7cluc20100do2yyq",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalCount, err := GetTotalListenLaterCountByUserId(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetTotalListenLaterCountByUserId() error = %v, wantErr %v", err, tt.wantErr)
			}

			t.Logf("the total count is %d", totalCount)
		})
	}
}
