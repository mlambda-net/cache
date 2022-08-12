package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"redis-cluster-0-svc:6379",
			"redis-cluster-1-svc:6379",
			"redis-cluster-2-svc:6379",
			"redis-cluster-3-svc:6379",
			"redis-cluster-4-svc:6379",
			"redis-cluster-5-svc:6379",
		},
		ReadOnly: true,
	})

	line := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	ctx := context.Background()

	for i := 0; i < 1000000000000; i++ {

		cmd := rdb.Set(ctx, fmt.Sprintf("key-%d", i), line, time.Duration(8)*time.Hour)
		if cmd.Err() != nil {
			println(cmd.Err().Error())
		}

		resp := rdb.Get(ctx, fmt.Sprintf("key-%d", i))
		if resp.Err() != nil {
			println(resp.Err().Error())
		}
	}

}
