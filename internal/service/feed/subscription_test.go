package feed

import (
	"context"
	"guoshao-fm-web/internal/consts"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestSubFeedByKeyword(t *testing.T) {
	type args struct {
		ctx        context.Context
		userId     string
		keyword    string
		sortByDate int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sub feed by keyword sort by date",
			args: args{
				ctx:        gctx.New(),
				userId:     "1t5z27w7h00csfdx7cluc20100do2yyq",
				keyword:    `游戏`,
				sortByDate: 1,
			},
			wantErr: false,
		},
	}
	g.DB().SetDryRun(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SubFeedByKeyword(tt.args.ctx, tt.args.userId, tt.args.keyword, tt.args.sortByDate); (err != nil) != tt.wantErr {
				if err.Error() != consts.DB_DATA_ALREADY_EXIST {
					t.Errorf("SubFeedByKeyword() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
