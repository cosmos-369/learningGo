package httpserver

import (
	"fmt"
	"go_specs_greet/domain/interactions"
	"net/http"
)

func GetServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", GreetHandler)
	mux.HandleFunc("/curse", CurseHandler)

	return mux
}

func GreetHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	fmt.Fprint(writer, interactions.Greet(name))
}

func CurseHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	fmt.Fprint(writer, interactions.Curse(name))
}
