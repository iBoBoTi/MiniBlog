package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main(){
	r := chi.NewRouter()

	// Routes
	r.Get("/", index)
	r.Get("/add", CreatePost)
	r.Get("/post/{id}", RetrievePost)
	r.Get("/post/update/{id}", UpdatePost)
	r.Get("/delete/{id}", DeletePost)

	fmt.Println("Starting Server...")
	if err := http.ListenAndServe("", r); err!= nil{
		log.Fatal(err)
	}
}
