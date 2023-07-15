package feed

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/utility"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
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
				channelId: "o66b2cv6l9qr",
				itemId:    "akeq65sh2r9m4",
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
				if userListenLaterDto.Id == "" {
					t.Fatal("the item info is empty")
				}
			}

		})
	}

}

func TestGetListenLaterRSSByUserId(t *testing.T) {
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		args    args
		wantRss string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:    gctx.New(),
				userId: "1yx7pq47d00csom5aepfi481002dsdjg",
			},
			wantErr: false,
		},
	}
    i18nPath := utility.GetProjectAbsRootPath()+"resource/i18n"
    g.I18n().SetPath(i18nPath)
	g.I18n().SetLanguage("zh-CN")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            t.Log("i18nPath : ", i18nPath)
			gotRss, err := GetListenLaterRSSByUserId(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListenLaterRSSByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotRss)
		})
	}
}
