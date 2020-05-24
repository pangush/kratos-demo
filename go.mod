module kratos-demo

go 1.12

require (
	github.com/bilibili/kratos v0.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-redsync/redsync v1.4.2
	github.com/gogo/protobuf v1.3.0
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/wire v0.3.0
	github.com/prometheus/common v0.6.0
	github.com/stretchr/testify v1.4.0
	github.com/tidwall/buntdb v1.1.2
	go.etcd.io/etcd v0.0.0-20190917205325-a14579fbfb1a
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03
	google.golang.org/grpc v1.24.0
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.39.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190627132806-fd42eb6b336f
	golang.org/x/image => github.com/golang/image v0.0.0-20190703141733-d6a02ce849c9
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190607214518-6fa95d984e88
	golang.org/x/net => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190609082536-301114b31cce
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190612232758-d4e310b4a8a5
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.0.0-20191224000733-4ba822d8138d
	google.golang.org/appengine => github.com/golang/appengine v1.6.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc => github.com/grpc/grpc-go v1.16.0
)
