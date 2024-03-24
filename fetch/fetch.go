package fetch

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func RequestWithHeader(URL string) *http.Request {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		slog.Info(err.Error())
		return nil
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_KEY"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("API_HOST"))
	return req
}

func Fetch(URL string) []byte {
	var book []byte

	req := RequestWithHeader(URL)
	if !(req != nil) {
		slog.Info(errors.New("the request was not create").Error())
		return book
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Info(err.Error())
		return book
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		book, _ = io.ReadAll(response.Body)
	}
	return book
}
