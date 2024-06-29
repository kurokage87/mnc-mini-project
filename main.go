package main

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"mnc/config"
	"mnc/routers"
	"mnc/service"
)

func main() {
	//init redis
	err := config.InitRedis()
	if err != nil {
		log.Fatal(err)
	}

	//init database connection
	err = config.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}

	cronJob := cron.New()
	cronJob.AddFunc("*/1 * * * * *", func() {
		fmt.Println("Start Running Cron")
		fmt.Println("Start Running Cron Update Status Transaction")
		err := service.UpdateStatus(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	cronJob.Start()

	router := routers.Route()
	router.Run(":3000")
}
