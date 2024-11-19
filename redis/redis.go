package redis

import (
	"github.com/alibaba/opentelemetry-go-auto-instrumentation/api"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func newRedisOnEnter(call api.CallContext, conf redis.RedisConf, opts ...redis.Option) {
	if conf.DB > 0 {
		opts = append([]redis.Option{redis.WithDB(conf.DB)}, opts...)
	}
	call.SetParam(1, opts)
}
