package main

import "net/http"

func (app *application) index(rw http.ResponseWriter, req *http.Request){
	// handler for the homepage as well as listing all the available posts

}

func (app *application) CreatePost(rw http.ResponseWriter, req *http.Request){
	// handler for creating a post
}

func (app *application) RetrievePost(rw http.ResponseWriter, req *http.Request){
	// handler for viewing a single post
}

func (app *application) UpdatePost(rw http.ResponseWriter, req *http.Request){
	// handler for updating a single post
}

func (app *application) DeletePost(rw http.ResponseWriter, req *http.Request){
	// handler for deleting a post
}

