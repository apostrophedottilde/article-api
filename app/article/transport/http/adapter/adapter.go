package adapter

import (
	"fmt"
	"net/http"

	"github.com/apostrophedottilde/blog-article-api/app/article/controller"
	"github.com/gorilla/mux"
)

// HTTPAdapter implementation
type HTTPAdapter struct {
	router *mux.Router
}

// Start http adapter and listen for requests
func (adapter *HTTPAdapter) Start() error {
	fmt.Println("Starting HTTP connection...")
	err := http.ListenAndServe(":8000", adapter.router)

	if err != nil {
		return fmt.Errorf("error starting server")
	}
	return nil
}

// Stop http adapter
func (adapter *HTTPAdapter) Stop() {
	adapter.router = nil
}

// New creates a new instance of HTTPAdapter and returns a pointer to it.
func New(c *controller.ArticleController) *HTTPAdapter {
	r := mux.NewRouter()

	r.HandleFunc("/articles", c.Create).Methods("POST")
	r.HandleFunc("/articles", c.FindAll).Methods("GET")
	r.HandleFunc("/articles/{id}", c.FindOne).Methods("GET")
	r.HandleFunc("/articles/{id}", c.Update).Methods("PUT")

	http.Handle("/", r)

	return &HTTPAdapter{
		router: r,
	}
}
