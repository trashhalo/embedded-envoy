package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"

	"github.com/gofiber/fiber"
)

const socket = "/tmp/app.socket"

func main() {
	_ = os.Remove(socket)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		fmt.Println("get request")
		c.Fasthttp.Request.Header.VisitAll(func(key, val []byte) {
			fmt.Println(string(key), "=", string(val))
		})
		if rand.Intn(2) == 0 {
			c.SendStatus(http.StatusInternalServerError)
		} else {
			c.Send("Hello, World 👋!")
		}
	})

	app.Post("/", func(c *fiber.Ctx) {
		fmt.Println("post request")
		c.Fasthttp.Request.Header.VisitAll(func(key, val []byte) {
			fmt.Println(string(key), "=", string(val))
		})
		if rand.Intn(2) == 0 {
			c.SendStatus(http.StatusInternalServerError)
		} else {
			c.Send("Hello, World 👋!")
		}
	})

	sock, err := net.Listen("unix", socket)
	if err != nil {
		panic(err)
	}
	defer sock.Close()

	err = app.Listener(sock)
	if err != nil {
		panic(err)
	}
}
