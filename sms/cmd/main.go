package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sms/config"
	"sms/src/router"
	"sms/src/service"
	"sms/src/store"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var cfg *config.Config

func initDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.MySqlUrl))
	if err != nil {
		panic(err)
	}

	dbx, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = dbx.Ping()
	if err != nil {
		panic("ping db error" + err.Error())
	}

	dbx.SetConnMaxIdleTime(5 * time.Minute)
	dbx.SetMaxIdleConns(25)
	dbx.SetMaxOpenConns(25)
	if cfg.Env != "prod" {
		db = db.Debug()
	}

	return db
}

func initRedis() *redis.Client {
	rdOtps, err := redis.ParseURL(cfg.RedisUrl)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(rdOtps)
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic("ping redis error" + err.Error())
	}

	return redisClient
}

func main() {
	var err error
	cfg, err = config.Init()
	if err != nil {
		panic(err)
	}

	if cfg.Env != "prod" {
		b, _ := json.MarshalIndent(cfg, "", "\t")
		fmt.Println(string(b))
	}

	fmt.Println("ping redis success")
	db := initDb()
	redisClient := initRedis()
	userStore := store.NewUseStore(db)
	svc := service.NewService(cfg, userStore, redisClient)
	engine := gin.Default()
	router.InitGin(engine, svc)
	go func() {
		engine.Run()
	}()

	// handler graceful shutdown
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	for {
		select {
		case <-osSig:
			{
				fmt.Println("shutting down")
				time.Sleep(25 * time.Second)
				dbx, _ := db.DB()
				dbx.Close()
				redisClient.Close()
				os.Exit(0)
			}
		}
	}
}
