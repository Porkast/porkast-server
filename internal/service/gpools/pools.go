package gpools

import "github.com/gogf/gf/v2/os/grpool"

var pool *grpool.Pool

func GetPool() *grpool.Pool {
	if pool == nil {
		pool = grpool.New(20)
	}

	return pool
}
