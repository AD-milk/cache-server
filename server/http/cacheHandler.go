package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type cacheHandler struct {
	*Server
}

func (cacheHandler *cacheHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := strings.Split(request.URL.EscapedPath(), "/")[2]
	if len(key) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	method := request.Method
	if method == http.MethodPut {
		body, _ := ioutil.ReadAll(request.Body)
		if len(body) != 0 {
			e := cacheHandler.Set(key, body)
			if e != nil {
				log.Println(e)
				writer.WriteHeader(http.StatusInternalServerError)
			}
		}
		return
	}
	if method == http.MethodGet {
		body, e := cacheHandler.Get(key)
		if e != nil {
			log.Println(e)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		if len(body) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.Write(body)
		return
	}

	if method == http.MethodDelete {
		e := cacheHandler.Del(key)
		if e != nil {
			log.Println(e)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	writer.WriteHeader(http.StatusMethodNotAllowed)

}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}
