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

resJSON.Body(w, "success", "data", data, 200)
```

The Body function takes 5 parameters: 
1. a response writer: http.ResponseWriter
2. statusMessage: string
3. dataKey: string
4. data: any or interface{}
5. statusCode: int

The status message should be displayed as "success" or "fail" for an easy to understand indication of the status of the response. The data key is the key for the data your displaying. For example "user" if the data you're displaying is user information. The data is, as stated before, the data you want to display. The status code is the status of the response, for example 200, 201, 400, 404, etc.
