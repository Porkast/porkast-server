package feed

import (
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/service/elasticsearch"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
)

func TestSearchFeedItemsByKeyword(t *testing.T) {
	var (
		ctx         = gctx.New()
		err         error
		itemDtoList []dto.FeedItem
		keyword     = "推荐"
		from        = 0
		size        = 10
	)

	genv.Set("GF_GCFG_FILE", "config.dev.yaml")
	elasticsearch.InitES(ctx)

	itemDtoList, err = SearchFeedItemsByKeyword(ctx, keyword, from, size)
	if err != nil {
		t.Fatal(err)
	}

	if len(itemDtoList) == 0 {
		t.Fatal("The search result is empty")
	}

}
