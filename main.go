package main

import (
	"log/slog"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mvitta/GoRoutineMiniProject/routes"
)

func main() {

	mux := http.NewServeMux()
	err := godotenv.Load(".env")
	if err != nil {
		slog.Info("Failed to load env")
	}

	mux.HandleFunc("GET /", routes.HandlerMain)

	http.ListenAndServe(":8080", mux)
}
