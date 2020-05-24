module github.com/wenwen1613/blog-example

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.56.0
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/unknwon/com v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/wenwen1613/blog-example/conf => /pkg/conf
	github.com/wenwen1613/blog-example/middleware => /middleware
	github.com/wenwen1613/blog-example/models => /models
	github.com/wenwen1613/blog-example/pkg/e => /pkg/e
	github.com/wenwen1613/blog-example/pkg/setting => /pkg/setting
	github.com/wenwen1613/blog-example/routers => /routers
)