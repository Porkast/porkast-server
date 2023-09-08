package jobs

import (
	"context"
	"guoshao-fm-web/internal/service/celery"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func Test_assignUserSubKeywordUpdateJob(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "assign user sub keyword update job to celery queue",
			args: args{
				ctx: gctx.New(),
			},
		},
	}
	for _, tt := range tests {
		celery.InitCeleryClient(tt.args.ctx)
		t.Run(tt.name, func(t *testing.T) {
			assignUserSubKeywordUpdateJob(tt.args.ctx)
		})
	}
}
