package responseJson

import (
	"encoding/json"
	"net/http"
)

var Version string = "1.0"

func ResponseJSON(w http.ResponseWriter, statusKey string, statusMessage string, dataKey string, data any, statusCode int) {
	res := make(map[string]any)
	res[statusKey] = statusMessage
	res[dataKey] = data
	resJSON, _ := json.MarshalIndent(res, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resJSON)
}
