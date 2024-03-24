package handlers

import (
	"github.com/LoomingLunar/LunarLoom-websocket-service/pkg/ws"
	"github.com/gofiber/contrib/websocket"
	"github.com/golang-jwt/jwt/v5"
)

// Handling websocket connection
func WebSocketHandler(c *websocket.Conn) {
	// Getting sessionId,userId from jwt token
	var user = c.Locals("user").(*jwt.Token)
	var claims = user.Claims.(jwt.MapClaims)
	var userId = claims["userId"].(string)

	// Creating new client and reading and writing messages
	var cli = ws.NewClient(userId, c)
	go cli.ListenMsg()
	cli.WriteMsg()
}
