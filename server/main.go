package main

import (
	"cache-server/server/cache"
	"cache-server/server/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
