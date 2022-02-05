package cache

import "log"

func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache()
	}
	if typ == "rocksdb" {
		c = newRocksdbCache()
	}
	if c == nil {
		panic("unknown cache type")
	}
	log.Println(typ, "ready to serve")
	return c
}
