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
}

func main(){
	addr := flag.String("addr",":8081","pass the network address")
	flag.Parse()

	// logger
	infoLog := log.New(os.Stdout,"INFO\t",log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errLog,
		infoLog: infoLog,
	}

	pswd := os.Getenv("MYSQL_PASSWORD")
	r := chi.NewRouter()
	db, err:= sql.Open("mysql","root:"+pswd+"@tcp(localhost:3306)/blogdb")
	if err!=nil{
		errLog.Fatal(err)
	}
	defer db.Close()

	// database check
	//err = db.Ping()
	//if err != nil {
	//	errLog.Fatal(err)
	//}


	// Routes
	r.Get("/", app.index)
	r.Get("/add", app.CreatePost)
	r.Get("/post/{id}", app.RetrievePost)
	r.Get("/post/update/{id}", app.UpdatePost)
	r.Get("/delete/{id}", app.DeletePost)

	server := &http.Server{
		Addr: *addr,
		ErrorLog: errLog,
		Handler: r,
	}

	infoLog.Println("Starting Server...")
	if err := server.ListenAndServe(); err!= nil{
		errLog.Fatal(err)
	}
}
