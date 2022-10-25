package resJSON

import (
	"encoding/json"
	"net/http"
)

var Version string = "0.2.0"

func Body(w http.ResponseWriter, statusMessage string, dataKey string, data interface{}, statusCode int) {
	res := make(map[string]any)
	res["status"] = statusMessage
	res[dataKey] = data
	b, _ := json.MarshalIndent(res, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(b)
}
