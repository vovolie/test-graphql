package main

import (
	"github.com/graphql-go/handler"
	"qiushibaike.com/test-vue/data"
	"log"
	"net/http"
)

func main() {

	h := handler.New(&handler.Config{
		Schema: &data.Schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	port := ":9999"
	log.Printf(`GraphQL server starting up on http://localhost%v`, port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", err)
	}
}