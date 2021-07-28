// @Title cron.go
// @Description:
// @Author  zwc
// @Date 2021/7/28
// @Update  zwc  2021/7/28

package main

//import (
//	"github.com/Wenchuan-Zhao/goBlogs/models"
//	"github.com/robfig/cron"
//	"log"
//	"time"
//)
//
//func main() {
//	log.Println("Starting...")
//
//	c := cron.New()
//	c.AddFunc("* * * * * *", func() {
//		log.Println("Run models.CleanAllTag...")
//		models.CleanAllTag()
//	})
//	c.AddFunc("* * * * * *", func() {
//		log.Println("Run models.CleanAllArticle...")
//		models.CleanAllArticle()
//	})
//
//	c.Start()
//
//	t1 := time.NewTimer(time.Second * 10)
//	for {
//		select {
//		case <-t1.C:
//			t1.Reset(time.Second * 10)
//		}
//	}
//}
