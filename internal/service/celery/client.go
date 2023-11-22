package celery

import (
	"context"

	"github.com/gocelery/gocelery"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gomodule/redigo/redis"
)

var goceleryClient *gocelery.CeleryClient

func GetClient() *gocelery.CeleryClient {
	return goceleryClient
}

func InitCeleryClient(ctx context.Context) {
	var (
		err         error
		redisPool   *redis.Pool
		redisAddr   *gvar.Var
		redisPass   *gvar.Var
		workerCount *gvar.Var
		brokeUrl    string
	)

	g.Log().Line().Info(ctx, "Start init gocelery client")
	redisAddr, _ = g.Cfg().Get(ctx, "redis.default.address")
	redisPass, _ = g.Cfg().Get(ctx, "redis.default.pass")
	workerCount, _ = g.Cfg().Get(ctx, "celery.worker.count")

	brokeUrl = ":" + redisPass.String() + "@" + redisAddr.String() + "/2"

	redisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://" + brokeUrl)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	// initialize celery client
	goceleryClient, err = gocelery.NewCeleryClient(
		gocelery.NewRedisBroker(redisPool),
		&gocelery.RedisCeleryBackend{Pool: redisPool},
		workerCount.Int(),
	)

	if err != nil {
		g.Log().Line().Fatal(ctx, "Failed to create new celery client : ", err)
	}

}
