package main

import (
	"fmt"

	"github.com/alicebob/miniredis/v2"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {

	r1, err := miniredis.Run()
	logx.Must(err)
	defer r1.Close()

	cfg := fmt.Sprintf(`Host: %s
DB: 1`, r1.Addr())

	var c redis.RedisConf
	err = conf.LoadConfigFromYamlBytes([]byte(cfg), &c)
	logx.Must(err)
	fmt.Println(c)

	rdb := c.NewRedis()
	rdb.Set("test", "test")
	fmt.Println(rdb.Get("test"))

	fmt.Println(r1.DB(1).Get("test"))
}
