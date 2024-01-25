package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func handleWSOrderBook(s *Server) func(ws *websocket.Conn) {
	return func(ws *websocket.Conn) {
		fmt.Println("New incoming connection from client to orderbook feed:", ws.RemoteAddr())
		//utils.UUIDv4()
		for {
			payload := fmt.Sprintf("orderbook data -> %d\n", time.Now().UnixNano())
			ws.WriteJSON(payload)
			time.Sleep(time.Second * 1)
		}
	}
}
func main() {

	app := fiber.New()
	app.Use("ws/bid", websocket.New(handleWSOrderBook(&Server{})))
	go func() {
		port := "3011"
		if port == "" {
			port = "3011"
		}

		if err := app.Listen(":" + port); err != nil {
			log.Panic(err)
		}
	}()
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	log.Println("Fiber was successful shutdown.")
}
