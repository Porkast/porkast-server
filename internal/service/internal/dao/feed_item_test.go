package dao

import (
	"context"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestGetFeedItemTotalCount(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Get all feed item count",
			args: args{
				ctx: gctx.New(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, err := GetFeedItemTotalCount(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFeedItemTotalCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotCount == 0 {
				t.Log("Total feed item count is ", gotCount)
			}

		})
	}
}
