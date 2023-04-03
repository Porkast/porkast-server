package controller

import (
	"context"
	"guoshao-fm-web/apiv1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	IndexTpl = cIndex{}
)

type cIndex struct{}

func (h *cIndex) IndexTpl(ctx context.Context, req *apiv1.IndexReq) (res *apiv1.IndexRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
