package feed

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestGetChannelInfoByChannelId(t *testing.T) {

	type args struct {
		ctx       context.Context
		channelId string
		offset    int
		limit     int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get channel info by channel id",
			args: args{
				ctx:       gctx.New(),
				channelId: "o66b2cv6l9qr",
				offset:    0,
				limit:     10,
			},
			wantErr: false,
		},
		{
			name: "get channel info by channel id without limit and offset",
			args: args{
				ctx:       gctx.New(),
				channelId: "o66b2cv6l9qr",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channelInfo, err := GetChannelInfoByChannelId(tt.args.ctx, tt.args.channelId, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetChannelInfoByChannelId() error = %v, wantErr %v", err, tt.wantErr)
			}
    
            if channelInfo.Count == 0 {
				t.Fatal("GetChannelInfoByChannelId() failed, channel item count is 0")
            }

		})
	}
}
