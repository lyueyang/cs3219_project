package main

import (
	"cs3219_project/handlers"
	"cs3219_project/storage"
	"github.com/gin-gonic/gin"

	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "users"
)

func main() {
	db := dbConnect()
	redis := redisConnect()

	dbStorage := storage.NewDBStorage(db)
	redisStorage := storage.NewRedisStorage(redis)

	hdlr := handlers.AccountHandler{DBService: dbStorage, RedisService: redisStorage}

	r := gin.Default()

	// routing logic here
	r.GET("/accounts", hdlr.HandleGetAccounts)
	r.POST("/accounts", hdlr.HandleCreateAccounts)

	err := r.Run(":6789")
	if err != nil {
		fmt.Println("error starting server")
	}
}

func dbConnect() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic("db connection failed")
	} else {
		return db
	}
}

func redisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		panic("redis connection failed")
	} else {
		fmt.Println("redis connection live: " + pong)
		return client
	}
}
