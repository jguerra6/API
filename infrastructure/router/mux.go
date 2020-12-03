package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

//NewMuxRouter returns a gorilla mux router
func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) DELETE(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}

func (*muxRouter) SERVE(port string) {
	log.Println("Listening to port ", port)
	http.ListenAndServe(port, muxDispatcher)
}

func (*muxRouter) GetIDFromRequest(request *http.Request) string {
	vars := mux.Vars(request)
	id := vars["id"]
	return id
}
