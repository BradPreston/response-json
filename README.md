# response-json
The response-json module allows you to return a well formatted body in an HTTP handler.

## How to use
To install response-json in your project:
```
go get github.com/BradPreston/response-json
```

To use in your file:
```
import resJSON "github.com/BradPreston/response-json"

resJSON.Body(w, "status", "success", "data", data, 200)
```

The Body function takes 5 parameters: a response writer, status message, data key, data, and status code.

The status message should be displayed as "success" or "fail" for an easy to understand indication of the status of the response. The data key is the key for the data your displaying. For example "user" if the data you're displaying is user information. The data is, as stated before, the data you want to display. The status code is the status of the response, for example 200, 201, 400, 404, etc.