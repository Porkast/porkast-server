package workers

import (
	"guoshao-fm-web/internal/service/elasticsearch"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestUpdateUserSubkeyword(t *testing.T) {
	type args struct {
		keyword     string
		lang        string
		orderByDate int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "update user sub keyword worker",
			args: args{
				keyword:     "游戏",
				lang:        "zh",
				orderByDate: 0,
			},
		},
	}

	elasticsearch.InitES(gctx.New())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateUserSubkeyword(tt.args.keyword, tt.args.lang, tt.args.orderByDate)
		})
	}
}
