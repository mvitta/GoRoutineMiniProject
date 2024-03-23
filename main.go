package main

import (
	"net/http"

	"github.com/mvitta/GoRoutineMiniProject/routes"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", routes.HandlerMain)

	http.ListenAndServe(":8080", mux)
}
