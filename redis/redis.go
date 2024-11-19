package redis

import (
	"fmt"

	"github.com/alibaba/opentelemetry-go-auto-instrumentation/api"
	"github.com/zeromicro/go-zero/core/errorx"
)

// NewRedis returns a Redis with given options.
func NewRedis(conf RedisConf, opts ...Option) (*Redis, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	if conf.Type == ClusterType {
		opts = append([]Option{Cluster()}, opts...)
	}
	if len(conf.Pass) > 0 {
		opts = append([]Option{WithPass(conf.Pass)}, opts...)
	}
	if conf.Tls {
		opts = append([]Option{WithTLS()}, opts...)
	}
	if conf.DB > 0 {
		opts = append([]Option{WithDB(conf.DB)}, opts...)
	}

	rds := newRedis(conf.Host, opts...)
	if !conf.NonBlock {
		if err := rds.checkConnection(conf.PingTimeout); err != nil {
			return nil, errorx.Wrap(err, fmt.Sprintf("redis connect error, addr: %s", conf.Host))
		}
	}

	return rds, nil
}

func newRedisOnEnter(call api.CallContext, conf RedisConf, opts ...Option) {
	call.SetSkipCall(true)
	rdb, err := NewRedis(conf, opts...)
	if err != nil {
		call.SetReturn(nil, err)
	}
	call.SetReturn(rdb, nil)
}
