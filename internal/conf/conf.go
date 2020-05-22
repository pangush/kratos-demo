package conf

import (
	"fmt"
	"kratos-demo/pkg/jwtauth"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"go.etcd.io/etcd/clientv3"
)

var (
	// Conf config
	Conf = &Config{}
)

// Config .
type Config struct {
	Log    *log.Config
	Server *bm.ServerConfig
	Redis  *redis.Config
	Etcd   *clientv3.Config
	MySQL  *sql.Config
	Jwt    *jwtauth.Config
	App    *App
}

// App .
type App struct {
	AppName     string
	CorsDomains string
	AppCode		int
}

// Init init conf
func Init() error {
	if err := paladin.Init(); err != nil {
		panic(err)
	}

	return local()
}

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

	if err := paladin.Get("jwt.toml").UnmarshalTOML(&Conf); err != nil {
		// 不存在时，将会为nil使用默认配置
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	if err := paladin.Get("app.toml").UnmarshalTOML(&Conf); err != nil {
		// 不存在时，将会为nil使用默认配置
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	fmt.Printf("%+v", Conf.App)

	return
}
