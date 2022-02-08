package main

import (
	"cache-server/server/cache"
	"cache-server/server/cluster"
	"cache-server/server/http"
	"cache-server/server/tcp"
	"flag"
	"log"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	node := flag.String("node", "127.0.0.1", "node address")
	clus := flag.String("cluster", "", "cluster address")
	ttl := flag.Int("ttl", 30, "cache time to live")
	flag.Parse()
	log.Println("type is: ", *typ)
	log.Println("node address is: ", *node)
	log.Println("cluster is: ", *clus)
	log.Println("ttl is: ", *ttl)
	c := cache.New(*typ, *ttl)
	n, e := cluster.New(*node, *clus)
	if e != nil {
		panic(e)
	}
	go tcp.New(c, n).Listen() // add tcp service
	http.New(c, n).Listen()
}
