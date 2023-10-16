package network

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

func GetHttpClient() (client *gclient.Client) {

	var (
		env string
	)
	env = os.Getenv("env")
	client = g.Client()
	client.SetTimeout(time.Second * 60)
	if env == "dev" {
		client.SetProxy("http://127.0.0.1:51491")
	}

	return
}

func GetContent(ctx context.Context, link string) (resp string) {
	var (
		client *gclient.Client
	)
	client = GetHttpClient()
	r, err := client.SetHeaderMap(getHeaders()).Get(ctx, link)
	defer r.Close()
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return
	}

	resp = r.ReadAllString()
	return
}

func TryGetRSSContent(ctx context.Context, link string) (resp string) {
	var (
		client *gclient.Client
	)
	if link == "" {
		g.Log().Line().Error(ctx, "The request link is empty")
		return
	}
	client = GetHttpClient()
	r, err := client.SetHeaderMap(getHeaders()).Get(ctx, link)
	defer func(resp *gclient.Response) {
		if rec := recover(); rec != nil {
			g.Log().Line().Error(ctx, fmt.Sprintf("Get RSS content by link %s failed: \n%s\n", link, rec))
		}
		if resp != nil {
			resp.Close()
		}
	}(r)

	if err != nil {
		g.Log().Line().Error(ctx, err)
		return
	} else if r == nil {
		g.Log().Line().Errorf(ctx, "The response with url %s is empty", link)
		return
	} else if r != nil && r.StatusCode == 404 {
		g.Log().Line().Errorf(ctx, "The response with url %s is 404", link)
		return
	}

	resp = r.ReadAllString()
	return
}

func GetContentByMobile(ctx context.Context, link string) (resp string) {
	var (
		client *gclient.Client
	)
	client = GetHttpClient()
	resp = client.SetHeaderMap(getMobileHeader()).GetContent(ctx, link)

	return
}

func PostContentByMobile(ctx context.Context, link string, data ...interface{}) (resp string) {
	var (
		err      error
		client   *gclient.Client
		response *gclient.Response
	)
	client = GetHttpClient()
	response, err = client.SetHeaderMap(getMobileHeader()).Post(ctx, link, data)
	defer response.Close()
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Do mobile post by url %s failed: \n%s\n", link, err))
	}

	resp = response.ReadAllString()

	return
}

func getHeaders() map[string]string {
	headers := make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63"
	return headers
}

func getMobileHeader() (headers map[string]string) {
	headers = make(map[string]string)
	headers["accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	headers["user-agent"] = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"
	return
}
