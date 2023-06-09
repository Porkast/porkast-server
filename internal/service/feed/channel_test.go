package feed

import (
	"context"
	"guoshao-fm-web/internal/service/elasticsearch"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
)

func TestGetChannelInfoByChannelId(t *testing.T) {

	type args struct {
		ctx       context.Context
		channelId string
		offset    int
		limit     int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get channel info by channel id",
			args: args{
				ctx:       gctx.New(),
				channelId: "o66b2cv6l9qr",
				offset:    0,
				limit:     10,
			},
			wantErr: false,
		},
		{
			name: "get channel info by channel id without limit and offset",
			args: args{
				ctx:       gctx.New(),
				channelId: "o66b2cv6l9qr",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channelInfo, err := GetChannelInfoByChannelId(tt.args.ctx, tt.args.channelId, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetChannelInfoByChannelId() error = %v, wantErr %v", err, tt.wantErr)
			}

			if channelInfo.Count == 0 {
				t.Fatal("GetChannelInfoByChannelId() failed, channel item count is 0")
			}

		})
	}
}

func TestQueryFeedChannelByKeyword(t *testing.T) {
	type args struct {
		ctx          context.Context
		searchParams SearchParams
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get channel info by channel id",
			args: args{
				ctx: gctx.New(),
				searchParams: SearchParams{
					Keyword:    "游戏",
					Page:       0,
					Size:       10,
					SortByDate: 1,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genv.Set("GF_GCFG_FILE", "config.dev.yaml")
			elasticsearch.InitES(tt.args.ctx)
			channelInfoList, err := QueryFeedChannelByKeyword(tt.args.ctx, tt.args.searchParams)
			if (err != nil) != tt.wantErr {
				t.Fatalf("QueryFeedChannelByKeyword() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(channelInfoList) == 0 {
				t.Fatalf("QueryFeedChannelByKeyword() len = 0")
			}

			t.Logf("QueryFeedChannelByKeyword() len %d, channel list :\n %+v", len(channelInfoList), channelInfoList)

		})
	}
}

func Test_formatItemShownotes(t *testing.T) {
	type args struct {
		shownotes string
	}

	tests := []struct {
		name string
		args args
	}{

		{

			name: "format shownotes without error",
			args: args{
				shownotes: `这是最好的时代，这是最坏的时代；这是智慧的时代，这是愚蠢的时代；这是信仰的时期，这是怀疑的时期。<br><br>狄更斯老先生写在「双城记」里的这段话精炼的概括了我们这个充满了被科技所包裹的「人文关怀」的世界的本质。原速也罢，加速也罢，都是我们面对生活的态度。<br><br><br><strong># 内容提要</strong><br>08:29 ·  「硅谷的工程师们的傲慢」<br>14:49 ·  本次的讨论缘起于网易的一次改版<br>21:26 ·  从主播的角度说说倍速<br>01:27:54 ·  讨论倍速播放离不开大家收听音频内容的整体习惯<br>36:10 ·  二号主播表示播客应该跟其他音频内容有很多本质区别<br><br><strong># 参考链接</strong><br><ol><li>EQ 是指音频播放时的「<a href="https://www.newvfx.com/forums/topic/68933">均衡器</a>」 6:01</li><li>网易云音乐小幅更新过的<a href="https://s.anw.red/x/netease-podcast-ui.jpg">播客专属播放界面</a> 15:00</li><li>播客播放器 Castro 的收费订阅<a href="https://castro.fm/#castro-plus">功能列表</a> 16:53</li><li>日剧《<a href="https://movie.douban.com/subject/10491666/">胜者即是正义</a>》 18:41</li><li>一段典型的 <a href="https://s.anw.red/x/leon-vs-jj-voice-wave.jpg">Leon vs. JJ 的声音波形</a> 22:16</li><li>一个典型的形象:《勇士》里开着老车<a href="https://s.anw.red/x/old-man-warrior.jpg">听着《白鲸记》有声书的老爹</a> 39:10</li><li>2012 年那时出版的 J·K·罗琳的书《<a href="https://book.douban.com/subject/19898714/">偶发空缺</a>》 40:11</li><li>我台<a href="https://public.flourish.studio/visualisation/75904/">上次听众调查时的数据</a> 42:22</li></ol><br><strong># 会员计划</strong><br>在<a href='http://r.anw.red/rss'>本台官网(Anyway.FM)</a> 注册会员即可 14 天试用 X 轴播放器和催更功能~ 开启独特的播客互动体验，Pro 会员更可加入听众群参与节目讨(hua)论(shui)~<br/><img src='https://s.anw.red/anyway.fm/rss-member.jpg' alt='Anyway.Member'/>`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatItemShownotes(tt.args.shownotes)
			if result == tt.args.shownotes {
				t.Fatalf("formatItemShownotes() not format shownotes")
			}
			t.Log(result)
		})
	}

}
