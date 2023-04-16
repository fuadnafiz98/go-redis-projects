package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func handleFixedWindow(rds *redis.Client, w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	log.Println(ip)
	if err != nil {
		log.Println(err)
	}
	// try to get the value with the given IP
	result, err := rds.Get(ctx, ip).Result()
	fmt.Println(result)
	if err != nil {
		// no key value pair.
		_, err = rds.Set(ctx, ip, 0, time.Second*20).Result()
		if err != nil {
			log.Println(err)
			return
		}
		w.Write([]byte("Current Count: 0\n"))
		return
	}

	// convert it to number
	currNumber, err := strconv.Atoi(result)
	if err != nil {
		log.Println(err)
	}
	if currNumber > 5 {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}

	_, err = rds.Incr(ctx, ip).Result()
	if err != nil {
		log.Println(err)
	}

	result, err = rds.Get(ctx, ip).Result()
	if err != nil {
		log.Println(err)
	}
	w.Write([]byte("Current Count: " + result + "\n"))
}

func withRedis(
	rds *redis.Client,
	handler func(rdx *redis.Client, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(rds, w, r)
	})
}

func InitServer(rds *redis.Client) *http.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/", withRedis(rds, handleFixedWindow))

	server := &http.Server{
		Addr:    ":8000",
		Handler: handler,
	}
	return server
}

func InitRedis(ctx *context.Context) (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr: ":6379",
		DB:   0,
	})

	_, err := rds.Ping(*ctx).Result()
	if err != nil {
		return nil, err
	}
	return rds, nil
}

func main() {
	ctx := context.Background()
	log.Println("Redis Initalizing...")
	rds, err := InitRedis(&ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server Initalizing...")
	server := InitServer(rds)
	defer server.Shutdown(ctx)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
