package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// App is the main struct for our application
type App struct {
	config *Config
	Router *mux.Router
	DB     *sql.DB
}

// Config represents the config details required for the app
type Config struct {
	DBType       string
	DBConnString string
	Host         string
	DevMode      bool
}

type handlerFunction func(c *Context)

// CreateApp takes in an Config and returns a configured App
func CreateApp(config Config) (App, error) {
	db, err := sql.Open(config.DBType, config.DBConnString)
	if err != nil {
		return App{}, err
	}

	newApp := App{
		config: &config,
		Router: mux.NewRouter(),
		DB:     db,
	}

	// Set NotFound (404) route handler
	newApp.Router.NotFoundHandler = newApp.Handle(func(c *Context) {
		c.Writer.WriteHeader(http.StatusNotFound)
	})

	// Set MethodNotAllowed (405) route handler
	newApp.Router.MethodNotAllowedHandler = newApp.Handle(func(c *Context) {
		c.Writer.WriteHeader(http.StatusMethodNotAllowed)
	})

	return newApp, nil
}

// Run launches the server and listens for connections
func (a *App) Run() {
	fmt.Printf("Server listening on %s\n", a.config.Host)
	http.ListenAndServe(a.config.Host, a.Router)
}

// DevModeEnabled returns the value of DevMode in the app's config
func (a *App) DevModeEnabled() bool {
	return a.config.DevMode
}

// GET configures a GET route for the path provided
func (a *App) GET(path string, handler handlerFunction) {
	buildHandlerFunc(a, path, handler).Methods("GET")
}

// POST configures a POST route for the path provided
func (a *App) POST(path string, handler handlerFunction) {
	buildHandlerFunc(a, path, handler).Methods("POST")
}

// PUT configures a PUT route for the path provided
func (a *App) PUT(path string, handler handlerFunction) {
	buildHandlerFunc(a, path, handler).Methods("PUT")
}

// DELETE configures a DELETE route for the path provided
func (a *App) DELETE(path string, handler handlerFunction) {
	buildHandlerFunc(a, path, handler).Methods("DELETE")
}

// PATCH configures a PATCH route for the path provided
func (a *App) PATCH(path string, handler handlerFunction) {
	buildHandlerFunc(a, path, handler).Methods("PATCH")
}

func buildHandlerFunc(a *App, path string, handler handlerFunction) *mux.Route {
	return a.Router.HandleFunc(path, a.Handle(handler))
}

// Handle takes in a custom handlerFunction and turns it into a handlerFunc which can be understood by our router
func (a *App) Handle(handler handlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &Context{
			Request: r,
			Writer:  wrapResponseWriter(w),
			DB:      a.DB,
		}

		// Times are divided by 1000 to make them microseconds
		startTime := time.Now().UnixNano() / 1000
		handler(c)
		endTime := time.Now().UnixNano() / 1000

		logRequest(c, endTime-startTime)
	}
}
