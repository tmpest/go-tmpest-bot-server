package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := newTmpestBotServer()
	server.start()
}

type tmpestBotServer struct {
	server *http.Server
}

func newTmpestBotServer() *tmpestBotServer {
	return &tmpestBotServer{
		server: &http.Server{
			Addr:    ":44910",
			Handler: oauth2RedirectHandler{},
		},
	}
}

func (s *tmpestBotServer) start() {
	s.server.ListenAndServe()
}

type oauth2RedirectHandler struct{}

func (handler oauth2RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved a request!")
}
