package ctls

import (
	"context"
	"guoshao-fm-web/internal/service/elasticsearch"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func Test_genKeywordSubEntity(t *testing.T) {
	type args struct {
		ctx        context.Context
		userId     string
		keyword    string
		lang       string
		sortByDate int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Generate ks entity",
			args: args{
				ctx:        gctx.New(),
				userId:     "1t5z27w7h00csfdx7cluc20100do2yyq",
				keyword:    "游戏",
				lang:       "zh",
				sortByDate: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elasticsearch.InitES(tt.args.ctx)
			gotKsEntityList, err := genKeywordSubEntity(tt.args.ctx, tt.args.userId, tt.args.keyword, tt.args.lang, tt.args.sortByDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("genKeywordSubEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(gotKsEntityList) <= 0 {
				t.Error("genKeywordSubEntity() result list size is 0")
				return
			} else {
                t.Log(gotKsEntityList)
            }
		})
	}
}
