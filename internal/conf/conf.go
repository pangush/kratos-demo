package conf

import (
	"fmt"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"go.etcd.io/etcd/clientv3"
)

var (
	// Conf config
	Conf = &Config{}
)

type Config struct {
	Log *log.Config
	Server *bm.ServerConfig
	Client *redis.Config
	Remote *Remote
	Etcd *clientv3.Config
}

type Remote struct {
	Driver string
}

// Init init conf
func Init() error {
	if err := paladin.Init(); err != nil {
		panic(err)
	}

	if err := paladin.Get("remote.toml").UnmarshalTOML(&Conf); err != nil {
		// 不存在时，将会为nil使用默认配置
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	if Conf.Remote.Driver == "" {
		return local()
	}

	//return remote()
	return nil
}

//func remote() (err error) {
//	if client, err = apoll; err != nil {
//		return
//	}
//	if err = load(); err != nil {
//		return
//	}
//	go func() {
//		for range client.Event() {
//			log.Info("config reload")
//			if load() != nil {
//				log.Error("config reload error (%v)", err)
//			}
//		}
//	}()
//	return
//}

func local() (err error) {
	if err := paladin.Get("etcd.toml").UnmarshalTOML(&Conf); err != nil {
		// 不存在时，将会为nil使用默认配置
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	if err := paladin.Get("redis.toml").UnmarshalTOML(&Conf); err != nil {
		// 不存在时，将会为nil使用默认配置
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	if err := paladin.Get("log.toml").UnmarshalTOML(&Conf); err != nil {
		// 不存在时，将会为nil使用默认配置
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	if err := paladin.Get("http.toml").UnmarshalTOML(&Conf); err != nil {
		// 不存在时，将会为nil使用默认配置
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	fmt.Printf("%+v", Conf.Server, Conf.Log)
	fmt.Printf("%+v", Conf.Client, Conf.Log)

	return
}

func init() {

}