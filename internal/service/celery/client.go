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
		workerCount *gvar.Var
	)

	g.Log().Line().Info(ctx, "Start init gocelery client")
	redisAddr, _ = g.Cfg().Get(ctx, "redis.default.address")
	workerCount, _ = g.Cfg().Get(ctx, "celery.worker.count")

	redisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://" + redisAddr.String() +"/2")
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
