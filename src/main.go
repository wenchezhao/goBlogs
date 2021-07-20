// @Title main.go
// @Description:
// @Author  zwc
// @Date 2021/7/20
// @Update  zwc  2021/7/20

package main

import (
	"fmt"
	"github.com/Wenchuan-Zhao/goBlogs/pkg/setting"
	"github.com/Wenchuan-Zhao/goBlogs/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
