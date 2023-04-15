package main

import (
	"context"
	"log"
	"net/http"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/redis/go-redis/v9"
)

func InitServer() *http.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from server\n"))
	})
	server := &http.Server{
		Addr:    ":8000",
		Handler: handler,
	}
	return server
}

func InitSSEServer() (*http.Server, *sse.Server) {
	s := sse.NewServer(&sse.Options{})

	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from server\n"))
	})
	handler.Handle("/events/", s)

	server := &http.Server{
		Addr:    ":8000",
		Handler: handler,
	}
	return server, s
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
	log.Println("Server Initalizing...")
	server, sseServer := InitSSEServer()
	defer sseServer.Shutdown()
	ctx := context.Background()

	log.Println("Redis Initalizing...")
	rds, err := InitRedis(&ctx)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Println("Admin initialization")
		for {
			response := rds.XRead(ctx, &redis.XReadArgs{
				Streams: []string{"test-stm-2", "$"},
				Block:   0,
			})
			result, err := response.Result()

			if err != nil {
				log.Println(err)
			}
			if len(result) > 0 {
				data := result[0]
				messages := data.Messages
				for _, message := range messages {
					v := message.Values["name"]
					sseServer.SendMessage("/events/channel", sse.SimpleMessage(v.(string)))
				}
			}
		}
	}()

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
