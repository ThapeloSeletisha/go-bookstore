package utils 

import (
	"encoding/json"
	"io"
	"net/http"
)

// Unmarshals a request body from json to 
// whatever object is passed as variable x
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
} 