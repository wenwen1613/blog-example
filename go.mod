module github.com/wenwen1613/blog-example

go 1.13

require (
	github.com/astaxie/beego v1.12.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.56.0
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.6 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/net v0.0.0-20200528225125-3c3fba18258b // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	golang.org/x/tools v0.0.0-20200529172331-a64b76657301 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/ini.v1 v1.56.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/wenwen1613/blog-example/conf => /pkg/conf
	github.com/wenwen1613/blog-example/middleware => /middleware
	github.com/wenwen1613/blog-example/models => /models
	github.com/wenwen1613/blog-example/pkg/e => /pkg/e
	github.com/wenwen1613/blog-example/pkg/setting => /pkg/setting
	github.com/wenwen1613/blog-example/pkg/util => /pkg/util
	github.com/wenwen1613/blog-example/routers => /routers
)
