package main

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) index(rw http.ResponseWriter, req *http.Request){
	// handler for the homepage as well as listing all the available posts

	header:= rw.Header()
	header.Add("Content-Type","text/html")
	rw.WriteHeader(http.StatusOK)

	stmt := "SELECT * FROM blogposts"

	rows, err := app.posts.DB.Query(stmt)
	if err != nil{
		return
	}
	defer rows.Close()

	var posts []Posts
	for rows.Next(){
		var p Posts
		err:= rows.Scan(&p.Id, &p.Title, &p.Body)
		if err != nil{
			return
		}
		posts = append(posts,p)
	}
	if err !=nil{
		app.errorLog.Fatal(err)
	}
	tmpl, _ := template.ParseFiles("ui/templates/index.html")
	err = tmpl.Execute(rw, posts)

	if err != nil {
		app.errorLog.Fatal(err)
	}

}

func (app *application) CreatePost(rw http.ResponseWriter, req *http.Request){
	// handler for creating a post
	// handler to present a form to add post

	tmpl, err := template.ParseFiles("ui/templates/createBlogPost_form.html")
	if err != nil {
		app.errorLog.Fatal(err)
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		app.errorLog.Fatal(err)
	}
}

func (app *application) RetrievePost(rw http.ResponseWriter, req *http.Request) {
	// handler for viewing a single post
	header := rw.Header()
	header.Add("Content-Type", "text/html")
	rw.WriteHeader(http.StatusOK)

	post := Posts{}
	id,err:= strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		app.errorLog.Fatal(err)
	}


	err = app.posts.DB.QueryRow("SELECT * FROM `blogdb`.`blogposts` WHERE id=?;",id).Scan(&post.Id,&post.Title,&post.Body)

	if err != nil {
		app.errorLog.Fatal(err)
	}
	tmpl, _ := template.ParseFiles("ui/templates/post_detail.html")
	err = tmpl.Execute(rw, post)
	if err != nil {
		app.errorLog.Fatal(err)
	}
}

func (app *application) UpdatePost(rw http.ResponseWriter, req *http.Request){
	// handler for updating a single post
	req.ParseForm()
	id := req.FormValue("id")
	row := app.posts.DB.QueryRow("SELECT * FROM `blogdb`.`blogposts` WHERE id=?;",id)
	var p Posts
	err := row.Scan(&p.Id,&p.Title, &p.Body)
	if err != nil{
		app.errorLog.Fatal(err)
	}

	tmpl, _ := template.ParseFiles("ui/templates/editBlogPost_form.html")
	tmpl.Execute(rw,p)
}

func (app *application) DeletePost(rw http.ResponseWriter, req *http.Request){
	// handler for deleting a post
		id := chi.URLParam(req, "id")
		del, err:= app.posts.DB.Prepare("DELETE FROM `blogdb`.`blogposts` WHERE (`id`=?);")
		if err != nil{
			app.errorLog.Fatal(err)
		}
		defer del.Close()
		_, err = del.Exec(id)
		http.Redirect(rw, req, "/", http.StatusFound)
}


func (app *application) CreateBlogPostHandler(rw http.ResponseWriter, req *http.Request){
	var post = &Posts{}
	header := rw.Header()
	header.Add("Content-Type", "text/html")

	err := req.ParseForm()
	if err != nil {
		return 
	}

	title := req.PostFormValue("Title")
	body := req.PostFormValue("Body")

	post.Title= title
	post.Body = body

	stmt:= "INSERT INTO `blogdb`.`blogposts` (`title`, `body`) VALUES (?,?);"

	if title == "" || body == "" {
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		rw.WriteHeader(http.StatusCreated)
		prepare, err := app.posts.DB.Prepare(stmt)
		if err != nil {
			app.errorLog.Fatal(err)
		}
		defer prepare.Close()
		_,err = prepare.Exec(post.Title,post.Body)
	}

	tmpl, _ := template.ParseFiles("ui/templates/createdBlogPost_confirm.html")
	err = tmpl.Execute(rw, nil)
	if err != nil {
		app.errorLog.Fatal(err)
	}
}


func(app *application) UpdateBlogPostHandler(rw http.ResponseWriter, req *http.Request){
	id := chi.URLParam(req, "Id")

	req.ParseForm()
	//
	//if req.FormValue("Title") != "" && req.FormValue("Body") != "" {
	//	post := blog.Post{
	//		Id:    id,
	//		Title: req.FormValue("Title"),
	//		Body:  req.FormValue("Body"),
	//	}
	//
	//	for i, v := range blog.GolangBlog.Posts {
	//		if post.Id == v.Id {
	//			blog.GolangBlog.Posts[i] = post
	//		}
	//	}
	//}
	http.Redirect(rw, req, "/"+id, http.StatusFound)
}

