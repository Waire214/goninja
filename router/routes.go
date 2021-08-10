package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Waire214/goninja/controllers"
	"github.com/go-chi/chi/v5"
)

func ProductEndPoint() http.Handler {
	r := chi.NewRouter()

	r.Get("/product/{productbyid}", controllers.GetProductById)

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("router connected")
	return r
}
