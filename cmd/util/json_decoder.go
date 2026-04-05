package util

import (
	"encoding/json"
	"net/http"
)

func DecodeJson(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}