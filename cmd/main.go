package main

import (
	"flag"
	"github.com/bilibili/kratos/pkg/conf/env"
	"github.com/bilibili/kratos/pkg/naming/etcd"
	"github.com/bilibili/kratos/pkg/net/rpc/warden/resolver"
	"kratos-demo/internal/conf"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bilibili/kratos/pkg/log"
	"kratos-demo/internal/di"
)

func main() {
	// AppID your appid, ensure unique.
	env.AppID = "demo.service"

	flag.Parse()

	//paladin.Init()
	conf.Init()

	log.Init(conf.Conf.Log) // debug flag: log.dir={path}
	defer log.Close()

	// NOTE: 注意这段代码，表示要使用etcd进行服务发现
	resolver.Register(etcd.Builder(conf.Conf.Etcd))

	log.Info("kratos-demo start")

	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("kratos-demo exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
