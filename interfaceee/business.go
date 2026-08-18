package interfaceee

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
)

// stockScript lua脚本扣减库存
const stockScript = `
local stock = redis.call('GET', KEYS[1])
if not stock then
    return -1
end
local num = tonumber(stock)
local sub = tonumber(ARGV[1])
if sub > num then 
	return -1 
end
local remain = redis.call('DECRBY', KEYS[1], sub)
return remain
`

var (
	rdb  *redis.Client
	once sync.Once
)

func initRedis() {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		if _, err := rdb.Ping(context.Background()).Result(); err != nil {
			panic(err)
		}
	})
}

func DeStock(key string, num int64) (int64, bool, error) {
	initRedis()

	res, err := rdb.Eval(context.Background(), stockScript, []string{key}, []interface{}{num}).Result()
	if err != nil {
		return 0, false, err
	}

	val, ok := res.(int64)
	if !ok {
		return 0, false, errors.New("invalid result")
	}

	if val == -1 {
		return 0, false, nil
	}

	return val, true, nil
}

func doStock() {
	remain, success, err := DeStock("test", 10)
	if err != nil {
		panic(err)
	}

	if success {
		fmt.Println(remain)
	} else {
		fmt.Println("库存不足")
	}
}
