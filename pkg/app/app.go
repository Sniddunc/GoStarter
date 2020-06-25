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
	router *mux.Router
	DB     *sql.DB
	Host   string
}

// Config represents the config details required for the app
type Config struct {
	DBType       string
	DBConnString string
	Host         string
}

type handlerFunction func(c *Context)

// CreateApp takes in an Config and returns a configured App
func CreateApp(config Config) (App, error) {
	newApp := App{}

	db, err := sql.Open(config.DBType, config.DBConnString)
	if err != nil {
		return App{}, err
	}

	newApp.DB = db
	newApp.router = mux.NewRouter()
	newApp.Host = config.Host

	return newApp, nil
}

// Run launches the server and listens for connections
func (a *App) Run() {
	fmt.Printf("Server listening on %s\n", a.Host)
	http.ListenAndServe(a.Host, a.router)
}

// GET configures a GET route for the path provided
func (a *App) GET(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, handlerFunc).Methods("GET")
}

// POST configures a POST route for the path provided
func (a *App) POST(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, handlerFunc).Methods("POST")
}

// PUT configures a PUT route for the path provided
func (a *App) PUT(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, handlerFunc).Methods("PUT")
}

// DELETE configures a DELETE route for the path provided
func (a *App) DELETE(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, handlerFunc).Methods("DELETE")
}

// PATCH configures a PATCH route for the path provided
func (a *App) PATCH(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, handlerFunc).Methods("PATCH")
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
