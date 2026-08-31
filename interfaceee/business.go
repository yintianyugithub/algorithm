package interfaceee

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
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

	res, err := rdb.Eval(context.Background(), stockScript, []string{key}, []any{num}).Result()
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

var (
	recDb     *redis.ClusterClient
	recDbOnce sync.Once
	recMap    = map[string]string{
		"":           "127.0.0.1:6379",
		"10.244.0.3": "127.0.0.1:6379",
		"10.244.0.5": "127.0.0.1:6380",
		"10.244.0.6": "127.0.0.1:6381",
	}
)

func init() {
	RedisClusterDb(context.TODO())
}

func RedisClusterDb(ctx context.Context) {
	recDbOnce.Do(func() {
		if len(recMap) == 0 {
			logx.WithContext(ctx).Infof("redis cluster is empty，node number:%n", len(recMap))
			return
		}

		recDb = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{
				recMap[""],
			},
			ClientName:     "",
			NewClient:      nil,
			MaxRedirects:   0,
			ReadOnly:       false,
			RouteByLatency: false,
			RouteRandomly:  false,
			ClusterSlots: func(ctx context.Context) ([]redis.ClusterSlot, error) {

				// 连接任意 Redis 节点获取 slots
				tmp := redis.NewClient(&redis.Options{
					Addr: recMap[""],
				})

				slots, err := tmp.ClusterSlots(ctx).Result()
				if err != nil {
					return nil, err
				}

				for i := range slots {
					for j := range slots[i].Nodes {
						addr := slots[i].Nodes[j].Addr
						// Redis 返回 Pod IP
						for podIp, ip := range recMap {
							if strings.HasPrefix(addr, podIp) {
								slots[i].Nodes[j].Addr = ip
							}
						}
					}
				}

				return slots, nil
			},
			Dialer:                       nil,
			OnConnect:                    nil,
			Protocol:                     0,
			Username:                     "",
			Password:                     "",
			CredentialsProvider:          nil,
			CredentialsProviderContext:   nil,
			StreamingCredentialsProvider: nil,
			MaxRetries:                   0,
			MinRetryBackoff:              0,
			MaxRetryBackoff:              0,
			DialTimeout:                  0,
			DialerRetries:                0,
			DialerRetryTimeout:           0,
			DialerRetryBackoff:           nil,
			ReadTimeout:                  0,
			WriteTimeout:                 0,
			ContextTimeoutEnabled:        false,
			MaxConcurrentDials:           0,
			PoolFIFO:                     false,
			PoolSize:                     0,
			PoolTimeout:                  0,
			MinIdleConns:                 0,
			MaxIdleConns:                 0,
			MaxActiveConns:               0,
			ConnMaxIdleTime:              0,
			ConnMaxLifetime:              0,
			ConnMaxLifetimeJitter:        0,
			ReadBufferSize:               0,
			WriteBufferSize:              0,
			PipelineReadBufferSize:       0,
			PipelineWriteBufferSize:      0,
			PipelinePoolSize:             0,
			AutoPipelineOptions:          nil,
			TLSConfig:                    nil,
			DisableRoutingPolicies:       false,
			DisableIndentity:             false,
			DisableIdentity:              false,
			IdentitySuffix:               "",
			UnstableResp3:                false,
			PushNotificationProcessor:    nil,
			FailingTimeoutSeconds:        0,
			MaintNotificationsConfig:     nil,
			ShardPicker:                  nil,
			ClusterStateReloadInterval:   0,
		})

		if res, err := recDb.Ping(ctx).Result(); err != nil {
			panic(err)
		} else {
			if res, err := recDb.Set(ctx, "test", 1000000, -1).Result(); err != nil {
				panic(err)
			} else {
				logx.WithContext(ctx).Infof("set success:%s", res)
			}

			val, err2 := recDb.Get(ctx, "test").Result()
			logx.WithContext(ctx).Infof("get test,val:%s,err:%+v", val, err2)
			fmt.Println("Redis Ping:", res)
		}
	})
}

// UserLogin 用户登录
func UserLogin(ctx context.Context, uid int64) {
	recDb.SAdd(ctx, fmt.Sprintf("uid:%d{1}", uid), fmt.Sprintf("%d|%d", time.Now().Unix(), time.Now().Unix()+int64(rand.IntN(1000000000))))
}

// CalculateLoginTime 计算用户登录时长
func CalculateLoginTime(ctx context.Context, uid int64) {
	res, cursor, err := recDb.SScan(ctx, fmt.Sprintf("uid:%d{1}", uid), 0, "", -1).Result()
	if err != nil {
		logx.WithContext(ctx).Infof("redis scan err:%v", err)
		return
	}

	logx.WithContext(ctx).Infof("redis scan cursor:%v", cursor)

	if len(res) == 0 {
		logx.WithContext(ctx).Infof("redis scan res:%v", res)
		return
	}

	// 获取用户登录时长
	loginTimeItems := make([][]int64, 0)

	for _, v := range res {
		s, e, ok := strings.Cut(v, "|")
		if !ok {
			continue
		}

		loginTimeItems = append(loginTimeItems, []int64{str2Int64(s), str2Int64(e)})
	}

	// 用户登录时长合并区间
	sort.Slice(loginTimeItems, func(i, j int) bool {
		return loginTimeItems[i][0] < loginTimeItems[j][0]
	})

	mergeLoginTimes := make([][]int64, 0, len(loginTimeItems))

	for _, v := range loginTimeItems {
		if len(mergeLoginTimes) == 0 || mergeLoginTimes[len(mergeLoginTimes)-1][1] < v[0] {
			mergeLoginTimes = append(mergeLoginTimes, v)
		} else {
			mergeLoginTimes[len(mergeLoginTimes)-1][1] = max(v[0], loginTimeItems[len(mergeLoginTimes)-1][1])
		}
	}

	// 计算时长
	loginTime := int64(0)
	for _, v := range mergeLoginTimes {
		loginTime += v[1] - v[0]
	}

	logx.WithContext(ctx).Infof("用户:%d登录时长%d秒", uid, loginTime)
}
