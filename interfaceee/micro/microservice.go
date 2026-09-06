package micro

import (
	_ "embed"
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
	TokenBuketAlgorithm string
	tokenScript         = redis.NewScript(TokenBuketAlgorithm)
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
			logx.Infof("distribute redis success")
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
	res, err := DistributeLimiter.Eval(TokenBuketAlgorithm, []string{"{test}.tokens", "{test}.ms"}, []any{10, 1000, time.Now().Unix(), 1})
	if err != nil {
		logx.Errorf("%v", err)
	}
	logx.Infof("%v", res)

	res2, err := DistributeLimiter.ScriptRun(tokenScript, []string{"{test}.tokens", "{test}.ms"}, []any{10, 1000, time.Now().Unix(), 1})
	logx.Infof("%v,%v", res2, err)
}
