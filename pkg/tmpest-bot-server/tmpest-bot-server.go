package tmpestbotserver

import "net/http"

import "fmt"

type TmpestBotServer struct {
	server *http.Server
}

func NewTmpestBotServer() *TmpestBotServer {
	return &TmpestBotServer{
		server: &http.Server{
			Addr:    ":44910",
			Handler: oauth2RedirectHandler{},
		},
	}
}

func (s *TmpestBotServer) Start() {
	s.server.ListenAndServe()
}

type oauth2RedirectHandler struct{}

func (handler oauth2RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved a request!")
}
