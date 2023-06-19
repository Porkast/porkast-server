package feed

import (
	"context"
	"guoshao-fm-web/internal/dto"
	"guoshao-fm-web/internal/service/elasticsearch"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
)

func TestSearchFeedItemsByKeyword(t *testing.T) {
	var (
		ctx          = gctx.New()
		err          error
		itemDtoList  []dto.FeedItem
		keyword      = "推荐"
		from         = 0
		size         = 10
		searchParams SearchParams
	)

	genv.Set("GF_GCFG_FILE", "config.dev.yaml")
	elasticsearch.InitES(ctx)
	searchParams = SearchParams{
		Keyword:    keyword,
		Page:       from,
		Size:       size,
		SortByDate: 1,
	}

	itemDtoList, err = SearchFeedItemsByKeyword(ctx, searchParams)
	if err != nil {
		t.Fatal(err)
	}

	if len(itemDtoList) == 0 {
		t.Fatal("The search result is empty")
	}

}

func TestGetPubFeedItemsByDate(t *testing.T) {
	type args struct {
		ctx  context.Context
		date string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Get feed item by date",
			args: args{
				ctx:  gctx.New(),
				date: "2023-05-08",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItemList, err := GetPubFeedItemsByDate(tt.args.ctx, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPubFeedItemsByDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(gotItemList) == 0 {
				t.Fatal("the feed item list is empty")
			}
		})
	}
}
