package main

import (
	"fmt"
	"net/http"
)

func registartion(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		
		email := fmt.Sprint(req.FormValue("email"))
		login := fmt.Sprint(req.FormValue("login"))
		password := fmt.Sprint(req.FormValue("password"))
		firstName := fmt.Sprint(req.FormValue("firstName"))
		lastName := fmt.Sprint(req.FormValue("lastName"))
		
		if email == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found email'}")
		}else if login == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found login'}")
		}else if password == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found password'}")
		}else if firstName == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found first name'}")
		}else if lastName  == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found last name'}")
		}else{
			//запрос в бд список чатов
			fmt.Fprintf(w,"{'response':'ok'}")
		}
	}
}

func main(){
	http.HandleFunc("/registartion",registartion)	

	http.ListenAndServe(":8090",nil)
}