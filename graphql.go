package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/rs/cors"
	"qiushibaike.com/test-vue/data"
)

func main() {

	h := handler.New(&handler.Config{
		Schema: &data.Schema,
		Pretty: true,
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	http.Handle("/graphql", c.Handler(h))
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	port := ":9999"
	log.Printf(`GraphQL server starting up on http://localhost%v`, port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", err)
	}
}
