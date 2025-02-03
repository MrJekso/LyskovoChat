package main

import (
	"fmt"
	"net/http"
)

func authentication(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		
		jwt := fmt.Sprint(req.FormValue("jwt"))
		
		if login == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found login'}")
		}else if password == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found password'}")
		}else{
			//запрос в бд список чатов
			fmt.Fprintf(w,"{'response':'jwt'}")
		}
	}
}

func main(){
	http.HandleFunc("/authentication",authentication)	

	http.ListenAndServe(":8090",nil)
}