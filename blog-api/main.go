package main

import (
	"net/http"

	"example.com/blog-api/routes"
)

func main() {

	router := routes.New()
	router.ArticlesRoutes()
	r := router.GetHttpRouter()
	http.ListenAndServe("localhost:9000", r)
}
