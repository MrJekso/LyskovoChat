package main

import (
	"fmt"
	"net/http"
)



// GET RESPONSE
/*
params := req.URL.Query()
name := params.Get("name")
fmt.Fprintf(w,"{'name':'%s'}",name)
*/

func profile(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		fmt.Println(req.FormValue("jwt"))
		//запрос в бд список чатов
		fmt.Fprintf(w,"{'response':'ok'}")
	}
}

func open(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		fmt.Println(req.FormValue("jwt"))
		//запрос в бд список чатов
		fmt.Fprintf(w,"{'response':'ok'}")
	}
}

func append(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		fmt.Println(req.FormValue("jwt"))
		//запрос в бд список чатов
		fmt.Fprintf(w,"{'response':'ok'}")
	}
}


func chat(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		fmt.Println(req.FormValue("jwt"))
		//запрос в бд список чатов
		fmt.Fprintf(w,"{'response':'ok'}")
	}
}

func delete(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		fmt.Println(req.FormValue("jwt"))
		//запрос в бд список чатов
		fmt.Fprintf(w,"{'response':'ok'}")
	}
}

func dislike(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		fmt.Println(req.FormValue("jwt"))
		//запрос в бд список чатов
		fmt.Fprintf(w,"{'response':'ok'}")
	}
}

func like(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		fmt.Println(req.FormValue("jwt"))
		//запрос в бд список чатов
		fmt.Fprintf(w,"{'response':'ok'}")
	}
}

func main(){
	http.HandleFunc("/profile",profile)
	http.HandleFunc("/open",open)
	http.HandleFunc("/append",append)
	http.HandleFunc("/chat",chat)
	http.HandleFunc("/like",like)
	http.HandleFunc("/dislike",dislike)
	
	http.ListenAndServe(":8090",nil)
}