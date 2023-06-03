package elasticsearch

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
)

func TestQueryFeedChannelFull(t *testing.T) {
	type args struct {
		ctx     context.Context
		keyword string
		offset  int
		limit   int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "query feed channel by keyword",
			args: args{
				ctx:     gctx.New(),
				keyword: "游戏",
				offset:  0,
				limit:   10,
			},
            wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genv.Set("GF_GCFG_FILE", "config.dev.yaml")
			InitES(tt.args.ctx)
			gsElastic = GetClient()
			channelESInfoList, err := gsElastic.QueryFeedChannelFull(tt.args.ctx, tt.args.keyword, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Fatalf("QueryFeedChannelFull() error = %v, wantErr %v", err, tt.wantErr)
			}

            t.Logf("QueryFeedChannelFull() channelEsInfoList size :%d  %+v", len(channelESInfoList), channelESInfoList)
		})
	}
}
