package main

import (
	"github.com/robfig/cron"
	"github.com/wenwen1613/blog-example/models"
	"log"
	"time"
)

func main() {
	log.Println("Starting....")

	c := cron.New()

	_ = c.AddFunc("* * * * * *", func() {
		log.Println("clean all tags")
		models.CleanAllTag()
	})
	_ = c.AddFunc("* * * * * *", func() {
		log.Println("clean all article")
		models.CleanAllArticle()
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
