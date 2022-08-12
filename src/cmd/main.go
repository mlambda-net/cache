package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
)

func main() {
	params := &chart.Params{
		Symbol:   "AAPL",
		Interval: datetime.OneHour,
	}

	iter := chart.Get(params)

	ctx := context.Background()

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

	for iter.Next() {
		bar := iter.Bar()
		key := fmt.Sprintf("AAPL-%d", bar.Timestamp)
		println(key)
		err := rdb.Set(ctx, key, bar.Volume, time.Duration(10)*time.Minute)
		if err != nil {
			println("%v", err.String())
		}

	}

}
