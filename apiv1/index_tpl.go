package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/" tags:"Home Page" method:"get" summary:"Index template request"`
}
type IndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
