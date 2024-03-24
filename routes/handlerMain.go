package routes

import (
	"net/http"
	"os"

	"github.com/mvitta/GoRoutineMiniProject/fetch"
)

var HostClient string = "http://localhost:3000"

func HandlerMain(w http.ResponseWriter, r *http.Request) {

	book := fetch.Fetch(os.Getenv("API_REQ_2"))

	w.Header().Add("Access-Control-Allow-Origin", HostClient)
	w.Header().Add("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(book)
}
