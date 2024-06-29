package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"mnc/model"
	"time"
)

var RedisClient *redis.Client
var DB *gorm.DB

func InitRedis() (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",                   // host redis
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", // password redis
		DB:       0,                                  // redis database
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("error connect to redis : %v", err)
	}
	fmt.Println("Connected to Redis")
	return
}

func InitDatabase() (err error) {
	dsn := "host=localhost user=postgres password=1 dbname=mnc-test port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	fmt.Println("Connected to Database")
	// Create the extension for UUID generation if it doesn't exist
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	fmt.Println("Successfully create extension")

	// Migrate the schema
	err = DB.AutoMigrate(&model.User{}, &model.Balance{}, &model.Transaction{})
	if err != nil {
		log.Fatal("failed migration schema")
	}
	fmt.Println("Migration schema successfully")
	return
}
