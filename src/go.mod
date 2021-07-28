module github.com/Wenchuan-Zhao/goBlogs

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.7.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/robfig/cron v1.2.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/Wenchuan-Zhao/goBlogs/conf => ./pkg/conf
	github.com/Wenchuan-Zhao/goBlogs/middleware => ./middleware
	github.com/Wenchuan-Zhao/goBlogs/models => ./models
	github.com/Wenchuan-Zhao/goBlogs/pkg/e => ./pkg/e
	github.com/Wenchuan-Zhao/goBlogs/pkg/setting => ./pkg/setting
	github.com/Wenchuan-Zhao/goBlogs/pkg/util => ./pkg/util
	github.com/Wenchuan-Zhao/goBlogs/routers => ./routers
)
