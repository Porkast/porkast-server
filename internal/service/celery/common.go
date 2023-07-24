package celery

import (
	"context"
	"guoshao-fm-web/internal/service/cache"
	"time"

	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/container/gvar"
)

func RandInt(min int, max int) int {
	return grand.N(min, max)
}

func GetRandomStartTime() (startTime time.Duration) {
	var (
		randomInt int
	)

	randomInt = RandInt(5, 20)

	startTime = time.Second * time.Duration(randomInt)

	return
}

func IsJobStarted(ctx context.Context, key string) (isStart bool) {
	var (
		valueVal *gvar.Var
		err      error
	)

	valueVal, err = cache.GetCache(ctx, key)
	if err != nil {
		isStart = true
	} else if !valueVal.IsEmpty() {
		isStart = true
	} else {
		isStart = false
	}

	return
}

func JobIsStarted(ctx context.Context, key string) {
	cache.SetCache(ctx, key, key, int(60*60))
}
