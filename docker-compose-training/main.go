package main 

import (
	"context"
	"fmt"
	"time"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Hello everyone !")
	client := redis.NewClient(&redis.Options{
		Addr:  "redis:6379",
		Password: "mypassword",
		DB: 0,
	})

	time.Sleep(5 * time.Second)
	err := client.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("key is set")
	time.Sleep(10 * time.Minute)
}