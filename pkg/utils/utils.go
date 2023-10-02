package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, body interface{}) {
	if data, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(data), body); err != nil {
			return
		}

	}

}
