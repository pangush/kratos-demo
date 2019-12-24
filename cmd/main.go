package main

import (
	"flag"
	"kratos-demo/internal/conf"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bilibili/kratos/pkg/log"
	"kratos-demo/internal/di"
)

func main() {
	flag.Parse()

	//paladin.Init()
	conf.Init()

	log.Init(conf.Conf.Log) // debug flag: log.dir={path}
	defer log.Close()

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
