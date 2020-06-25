# GoStarter

Go Starter is a boilerplate project designed to make the development of ultra light-weight RESTful APIs a breeze.

It uses a custom built micro framework to provide useful tools and logging for API endpoint creation.

# Getting Started

The first thing you'll want to do is clone this repository.

Next, modify the module name in `go.mod` to whatever your project is called.
You'll also want to rename the `cmd/gostarter` folder to reflect the name of your project.

You can then configure the application in `cmd/gostarter/main.go`

### App Config
The default AppConfig looks like this:
```go
appConfig := app.Config{
  DBType:       "sqlite3",
  DBConnString: "./database.sqlite",
  Host:         ":5000",
  DevMode:      os.Getenv("MODE") == "development",
}
```

### Adding API Endpoints
To add a new endpoint, you should first create a handler function. An example handler function can be seen below:
```go
func ExampleHandler(c *app.Context) {
  c.JSON(http.StatusOK, common.StandardResponse{
    Success: true
  }
}
```

This will simply send back some JSON with a property called `success` which is set to true.

Next, you'll want to configure the route which will use your newly created handler function.
This should be done in main after the call to `app.CreateApp()`

GoStarter has support for the following HTTP methods:
* GET
* POST
* PUT
* DELETE
* PATCH

To register a GET route, you would do something like:
```go
app.GET("/api/v1/example", ExampleHandler)
```

Similarly, to register a POST route, you would do something like this:
```go
app.POST("/api/v1/example/post", ExamplePostHandler)
```

The same is true for the rest of the supported HTTP methods.

### Response Tools
##### StandardResponse
In the `common` package within `/pkg`, there is a type definition for `StandardResponse`. `StandardResponse` is declared as:
```go
type StandardResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
```

Using `StandardResponse` or something similar to it is recommended to keep your responses predictable.

##### SendUnauthorized
`SendUnauthorized` is a function attached to the `context` which is provided to each route handler function.

All it does is set the response status code to `401` and send the word "Unauthorized" back to the user.

##### SendBadRequest
`SendBadRequest` is a function attached to the `context` which is provided to each route handler function.

It sends back some generic `StandardResponse` JSON data with `success` being false, `message` being any string which you provide it and `errors` being any interface you provide it.

##### SendInternalError
`SendInternalError` is a function attached to the `context` which is provided to each route handler function.

It sends back some generic `StandardResponse` JSON data with `success` being false and `message` informing the user that something went wrong.

### Changing 404 NotFound and 405 MethodNotAllowed Handlers
GoStarter's framework uses gorilla mux, so all you need to do is set `app.Router.NotFoundHandler` or `app.Router.MethodNotAllowedHandler` to your own `http.HandlerFunc` handler function.
