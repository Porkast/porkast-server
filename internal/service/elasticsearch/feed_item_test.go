package elasticsearch

import (
	"guoshao-fm-web/internal/model/entity"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
)

func TestESClient_QueryFeedItemFull(t *testing.T) {
	var (
		ctx        = gctx.New()
		err        error
		gsElastic  *GSElastic
		resultList []entity.FeedItemESData
	)
	genv.Set("GF_GCFG_FILE", "config.dev.yaml")
	InitES(ctx)
	gsElastic = GetClient()
	resultList, err = gsElastic.QueryFeedItemFull(ctx, "推荐", 0, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(resultList) == 0 {
		t.Fatal("Elasticsearch query feed item is empty")
	}
}
