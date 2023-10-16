package feed

import (
	"context"
	"testing"
)

func TestSearchPodcastEpisodesFromItunes(t *testing.T) {
	type args struct {
		ctx           context.Context
		keyword       string
		country       string
		excludeFeedId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ctx:           context.TODO(),
				keyword:       "indie hacker",
				country:       "CN",
				excludeFeedId: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItemList, err := SearchPodcastEpisodesFromItunes(tt.args.ctx, tt.args.keyword, tt.args.country, tt.args.excludeFeedId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchPodcastEpisodesFromItunes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(gotItemList) == 0 {
				t.Errorf("SearchPodcastEpisodesFromItunes() gotItemList is empty")
			}
		})
	}
}
