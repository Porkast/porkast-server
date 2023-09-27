package jobs

import (
	"context"
	"porkast-server/internal/service/cache"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func Test_setItemTotalCountToCache(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Set feed item total count to cache",
			args: args{
				ctx: gctx.New(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.InitCache(tt.args.ctx)
			if err := setItemTotalCountToCache(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("setItemTotalCountToCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setLatestFeedItems(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Set latest item list to cache",
			args: args{
				ctx: gctx.New(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache.InitCache(tt.args.ctx)
			if err := setLatestFeedItems(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("setLatestFeedItems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_updateDailyFeedItemRecord(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
        {
        	name:    "update daily feed item record",
        	args:    args{
        		ctx: gctx.New(),
        	},
        	wantErr: false,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := updateDailyFeedItemRecord(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("updateDailyFeedItemRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
