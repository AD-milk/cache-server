package cache

// #include <stdlib.h>
// #include "rocksdb/c.h"
// #cgo CFLAGS: -I${SRCDIR}/../../rocksdb/include
// #cgo LDFLAGS: -L${SRCDIR}/../../rocksdb -ldl -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -O3
import "C"
import (
	"time"
	"unsafe"
)

const BATCH_SIZE = 100

func writeBatch(db *C.rocksdb_t, c chan *pair, writeOptions *C.rocksdb_writeoptions_t) {
	count := 0
	timer := time.NewTimer(time.Second)
	wb := C.rocksdb_writebatch_create()
	for {
		select {
		case p := <-c:
			count++
			key := C.CString(p.k)
			val := C.CBytes(p.v)
			C.rocksdb_writebatch_put(wb, key, C.size_t(len(p.k)), (*C.char)(val), C.size_t(len(p.v)))
			C.free(unsafe.Pointer(key))
			C.free(val)
			if count == BATCH_SIZE {
				flushBatch(db, wb, writeOptions)
				count = 0
			}
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(time.Second)
		case <-timer.C:
			if count != 0 {
				flushBatch(db, wb, writeOptions)
				count = 0
			}
			timer.Reset(time.Second)
		}
	}
}

func flushBatch(db *C.rocksdb_t, writeBatch *C.rocksdb_writebatch_t, writeOptions *C.rocksdb_writeoptions_t) {
	var err *C.char
	C.rocksdb_write(db, writeOptions, writeBatch, &err)
	if err != nil {
		panic(C.GoString(err))
	}
	C.rocksdb_writebatch_clear(writeBatch)
}

func (c *rocksdbCache) Set(key string, value []byte) error {
	c.ch <- &pair{key, value}
	return nil
}
