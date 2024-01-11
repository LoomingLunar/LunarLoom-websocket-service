package main

import (
	"os"

	"github.com/LoomingLunar/LunarLoom-websocket-service/database"
	"github.com/LoomingLunar/LunarLoom-websocket-service/pkg/redis"
	"github.com/LoomingLunar/LunarLoom-websocket-service/server"
	"github.com/joho/godotenv"
)

func main() {
	// Loading enviornment variables
	godotenv.Load(".env")

	// Server port
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "9000"
	}

	// Connecting and creating redis client
	database.RedisSetUp()
	redis.ListenChannels()

	server.Run(port)
}
