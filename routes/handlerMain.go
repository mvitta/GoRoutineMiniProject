package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/mvitta/GoRoutineMiniProject/fetch"
	"github.com/mvitta/GoRoutineMiniProject/types"
)

// https://app.quicktype.io/

var HostClient string = "http://localhost:5173"

type Books struct {
	Book1 types.Book `json:"book1"`
	Book2 types.Book `json:"book2"`
}
type Response struct {
	Res Books `json:"res"`
}

// var wg sync.WaitGroup

func HandlerMain(w http.ResponseWriter, r *http.Request) {

	var book1 types.Book
	var book2 types.Book

	fetch.ToDoRequest(os.Getenv("API_REQ_1"), &book1)
	fetch.ToDoRequest(os.Getenv("API_REQ_2"), &book2)

	res, _ := json.Marshal(
		Response{
			Res: Books{
				Book1: book1,
				Book2: book2,
			},
		},
	)

	// time.Sleep(time.Minute)
	w.Header().Add("Access-Control-Allow-Origin", HostClient)
	w.Header().Add("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))

}
