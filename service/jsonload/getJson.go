package jsonload

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetJson(url string, model interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error in getting response from %v: %w", url, err)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("error in reading response body: %w", err)
	}

	if err = json.Unmarshal(data, model); err != nil {
		return fmt.Errorf("error in json parsing: %w", err)
	}

	return nil
}
