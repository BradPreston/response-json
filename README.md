# response-json
The response-json module allows you to return a well formatted body in an HTTP handler.

## How to use
To install response-json in your project:
```
go get github.com/BradPreston/response-json
```

To use in your file:
```
import response-json "github.com/BradPreston/response-json"

response-json.ResponseJSON(w, "status", "success", "data", data, 200)
```
