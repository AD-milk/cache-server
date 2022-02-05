package main

import (
	"cache-server/server/cache"
	"cache-server/server/http"
	"cache-server/server/tcp"
	"flag"
	"log"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is: ", *typ)
	c := cache.New(*typ)
	go tcp.New(c).Listen() // add tcp service
	http.New(c).Listen()
}
