module ws/marketApi

go 1.13

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/gin-gonic/gin v1.5.0
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/zhaocong6/goUtils v1.0.3 // indirect
	github.com/zhaocong6/market v0.0.5
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sys v0.0.0-20200327173247-9dae0f8f5775 // indirect
	google.golang.org/genproto v0.0.0-20200330113809-af700f360a68 // indirect
	google.golang.org/grpc v1.29.0-dev.0.20200326222940-e965f2a60b15
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
	market.pd v0.0.0-00010101000000-000000000000
)

replace (
	github.com/marketApi/pkg/setting => ./pkg/setting
	market.pd => ./proto/market
)
