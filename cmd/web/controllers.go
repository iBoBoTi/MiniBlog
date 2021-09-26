package main

import "net/http"

func index(rw http.ResponseWriter, req *http.Request){
	// handler for the homepage as well as listing all the available posts

}

func CreatePost(rw http.ResponseWriter, req *http.Request){
	// handler for creating a post
}

func RetrievePost(rw http.ResponseWriter, req *http.Request){
	// handler for viewing a single post
}

func UpdatePost(rw http.ResponseWriter, req *http.Request){
	// handler for updating a single post
}

func DeletePost(rw http.ResponseWriter, req *http.Request){
	// handler for deleting a post
}

