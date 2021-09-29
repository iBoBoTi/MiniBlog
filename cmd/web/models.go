package main

import (
	"database/sql"
	//"html/template"
)

type Posts struct{
	Id int
	Title string
	Body string
}


type PostModel struct{
	DB *sql.DB
}

//var DB *sql.DB

//var tpl *template.Template

