package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()
var rdb *redis.Client

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("✅ Redis connected successfully!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New visitor!")

	// Atomic increment in Redis (no need to fetch and convert manually)
	visits, err := rdb.Incr(ctx, "visits").Result()
	if err != nil {
		http.Error(w, "Error updating visit count", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello, World! You are visitor number %d.", visits)
}

func main() {
	initRedis()

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	port := os.Getenv("APP_PORT")
	fmt.Println("🚀 Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
