package cache

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestSetCache(t *testing.T) {
	type args struct {
		ctx          context.Context
		key          string
		value        string
		expireSecond int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Set test cache",
			args: args{
				ctx:          gctx.New(),
				key:          "test_key",
				value:        "test_value",
				expireSecond: 24 * 60 * 60,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitCache(tt.args.ctx)
			if err := SetCache(tt.args.ctx, tt.args.key, tt.args.value, tt.args.expireSecond); (err != nil) != tt.wantErr {
				t.Errorf("SetCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
