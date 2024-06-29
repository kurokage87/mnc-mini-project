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

	c := cron.New(cron.WithSeconds())

	// Add a job that runs every 10 minutes
	c.AddFunc("0 */1 * * * *", func() {
		fmt.Println("Start Running Cron Update Status")
		err = service.UpdateStatus(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	})

	// Start the cron scheduler
	c.Start()

	router := routers.Route()
	router.Run(":3000")
}
