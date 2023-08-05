package dao

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestCreateDailyFeedItemRecord(t *testing.T) {
	type args struct {
		ctx       context.Context
		channelId string
		itemId    string
		pubDate   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
        {
        	name:    "create daily feed item record",
        	args:    args{
        		ctx:       gctx.New(),
        		channelId: "105v2e18jj2lf",
        		itemId:    "1ivif5fm8gueb",
        		pubDate:   "2023-05-01",
        	},
        	wantErr: false,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateDailyFeedItemRecord(tt.args.ctx, tt.args.channelId, tt.args.itemId, tt.args.pubDate); (err != nil) != tt.wantErr {
				t.Errorf("CreateDailyFeedItemRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
