package fetch

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

func ToDoRequest(RawURL string, b any) {
	URL, _ := url.Parse(RawURL)
	response, err := http.DefaultClient.Do(&http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme:   URL.Scheme,
			Host:     URL.Host,
			Path:     URL.Path,
			RawQuery: URL.RawQuery,
		},
	})

	if err != nil {
		slog.Info(err.Error())
	}

	if response.StatusCode == http.StatusOK {
		var data, _ = io.ReadAll(response.Body)
		defer response.Body.Close()
		json.Unmarshal(data, b)
	}

}
