package feed

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
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
			name: "get listen later info by user id and feed id",
			args: args{
				userId:    "1t5z27w7h00csfdx7cluc20100do2yyq",
				channelId: "8k4vnjjtcmjqi",
				itemId:    "100r8jf600hcm",
			},
			wantErr: false,
		},
	}
	g.DB().SetDryRun(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				err                error
				userListenLaterDto dto.UserListenLater
			)
			if userListenLaterDto, err = GetListenLaterByUserIdAndFeedId(tt.args.ctx, tt.args.userId, tt.args.channelId, tt.args.itemId); (err != nil) != tt.wantErr {
				t.Fatalf("GetListenLaterByUserIdAndFeedId() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if userListenLaterDto.ItemInfo.Id == "" {
					t.Fatal("the item info is empty")
				}
			}

		})
	}

}
