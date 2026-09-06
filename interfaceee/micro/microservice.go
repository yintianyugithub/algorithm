package micro

import (
	_ "embed"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

/*
*
服务治理
限流算法：计数法、令牌桶、漏桶、滑动窗口
go-zero中的限流算法：周期限流器、令牌桶限流器、限流器
熔断器：阻止级联故障
*
*/
var (
	DistributeLimiter     *redis.Redis
	DistributeLimiterOnce sync.Once
	//go:embed token_buket.lua
	tokenBuketAlgorithm string
	tokenScript         = redis.NewScript(tokenBuketAlgorithm)
	//go:embed window_list.lua
	windowAlgorithm string
	windowScript    = redis.NewScript(windowAlgorithm)
)

func init() {
	DistributeLimiterOnce.Do(func() {
		DistributeLimiter = redis.MustNewRedis(redis.RedisConf{
			Host: "localhost:6379",
			Type: "node",
		})

		if !DistributeLimiter.Ping() {
			logx.Errorf("distribute redis fail")
		} else {
			logx.Infof("Ping distribute redis success")
		}
	})
}

// AtkBuket go-zero令牌桶
func AtkBuket() {
	atkLimiter := limit.NewTokenLimiter(1, 100, DistributeLimiter, "test")
	fmt.Println(atkLimiter.Allow())
}

// GetAtkBuk 令牌桶算法
func GetAtkBuk() {
	res, err := DistributeLimiter.Eval(tokenBuketAlgorithm, []string{"{test}.tokens", "{test}.ms"}, []any{10, 1000, time.Now().Unix(), 1})
	// redis的bool类型，返回 nil
	if errors.Is(err, redis.Nil) {
		logx.Error("rsp 429")
		return
	}
	if err != nil {
		logx.Errorf("%v", err)
	}

	logx.Infof("%v", res)

	res2, err := DistributeLimiter.ScriptRun(tokenScript, []string{"{test}.tokens", "{test}.ms"}, []any{10, 1000, time.Now().Unix(), 1})
	// redis的bool类型，返回 nil
	if errors.Is(err, redis.Nil) {
		logx.Error("rsp 429")
		return
	}

	logx.Infof("%v,%v", res2, err)
}

// WindowAlgorithmByList 滑动窗口限流 基于list实现
func WindowAlgorithmByList() {
	for i := range 11 {
		res, err := DistributeLimiter.ScriptRun(windowScript, []string{"{test}.window"}, []any{10, 100, time.Now().Unix()})
		if errors.Is(err, redis.Nil) {
			logx.Errorf("%d:rsp 429", i)
			continue
		}
		if err != nil {
			logx.Error(err)
		} else {
			logx.Infof("window algorithm_%d: %v", i, res.(int64))
		}
	}
}

// WindowAlgorithmByZSet 滑动窗口限流 基于ZSet实现 10req/s
func WindowAlgorithmByZSet() {
	for range 100 {
		now := time.Now().UnixMilli()
		key := "test"
		_, _ = DistributeLimiter.Zadd(key, now, fmt.Sprintf("%d", now))
		_, _ = DistributeLimiter.Zremrangebyscore(key, 0, now-1000)
		count, _ := DistributeLimiter.Zcard(key)
		if count > 10 {
			logx.Error("429")
		} else {
			logx.Infof("%d", count)
		}
	}
}
