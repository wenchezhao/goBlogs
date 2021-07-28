// @Title main.go
// @Description:
// @Author  zwc
// @Date 2021/7/20
// @Update  zwc  2021/7/20

package main

import (
	"context"
	"fmt"
	"github.com/Wenchuan-Zhao/goBlogs/models"
	"github.com/Wenchuan-Zhao/goBlogs/pkg/logging"
	"github.com/Wenchuan-Zhao/goBlogs/pkg/setting"
	"github.com/Wenchuan-Zhao/goBlogs/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//// linux下的优雅重启（"github.com/fvbock/endless"）
	//setting.Setup()
	//models.Setup()
	//logging.Setup()
	//endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}

	setting.Setup()
	models.Setup()
	logging.Setup()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.ReadTimeout),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
