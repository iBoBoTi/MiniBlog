package main

import (
	"database/sql"
	"flag"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"os"
)



type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
	posts *PostModel
}

func main(){
	addr := flag.String("addr",":8080","pass the network address")
	flag.Parse()

	// logger
	infoLog := log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)



	pswd := os.Getenv("MYSQL_PASSWORD")
	r := chi.NewRouter()
	db, err:= sql.Open("mysql","root:"+pswd+"@tcp(localhost:3306)/blogdb")
	if err!=nil{
		errLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: errLog,
		infoLog: infoLog,
		posts: &PostModel{DB:db},
	}

	// database check
	err = db.Ping()
	if err != nil {
		errLog.Fatal(err)
	}


	// Routes
	r.Get("/", app.index)
	r.Get("/create", app.CreatePost)
	r.Get("/{id}", app.RetrievePost)
	r.Get("/update/{Id}", app.UpdatePost)
	r.Get("/delete/{Id}", app.DeletePost)
	r.Post("/add", app.CreateBlogPostHandler)
	r.Post("/postupdate/{Id}", app.UpdateBlogPostHandler)

	server := &http.Server{
		Addr: *addr,
		ErrorLog: errLog,
		Handler: r,
	}

	err = db.Ping()
	if err != nil {
		errLog.Fatal(err)
	}

	infoLog.Printf("Starting Server on %s...",*addr)
	if err := server.ListenAndServe(); err!= nil{
		errLog.Fatal(err)
	}
}
