package main

import (
	"cache-server/server/cache"
	"cache-server/server/http"
	"cache-server/server/tcp"
)

func main() {
	c := cache.New("inmemory")
	go tcp.New(c).Listen() // add tcp service
	http.New(c).Listen()
}
