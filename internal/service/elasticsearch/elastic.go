package elasticsearch

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"
)

type GSElastic struct {
	Client *elastic.Client
}

var gsElastic *GSElastic

func InitES(ctx context.Context) {
	var (
		host     *gvar.Var
		port     *gvar.Var
		username *gvar.Var
		password *gvar.Var
		url      string
	)
	host, _ = g.Cfg().Get(ctx, "elastic.host")
	port, _ = g.Cfg().Get(ctx, "elastic.port")
	username, _ = g.Cfg().Get(ctx, "elastic.username")
	password, _ = g.Cfg().Get(ctx, "elastic.password")
	url = host.String() + ":" + port.String()
	g.Log().Line().Infof(ctx, "Ready to connect to elasticsearch with url %s", url)
	var err error
	gsElastic = &GSElastic{}
	gsElastic.Client, err = elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false), elastic.SetHealthcheck(false), elastic.SetBasicAuth(username.String(), password.String()))
	if err != nil {
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := gsElastic.Client.Ping(url).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	g.Log().Line().Infof(ctx, "Elasticsearch returned with code %d and version %s", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esVersion, err := gsElastic.Client.ElasticsearchVersion(url)
	if err != nil {
		// Handle error
		panic(err)
	}
	g.Log().Line().Infof(ctx, "Elasticsearch version %s", esVersion)
}

func GetClient() *GSElastic {
	return gsElastic
}
