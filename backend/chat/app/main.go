package main

import (
	"fmt"
	"net/http"
//	"encoding/json"
)

//healthcheck
//chats
//chats/{chats_id}
//chats/{chats_id}/delete
//chats/{chats_id}/append
//chats/{chats_id}/edit_name
//chats/{chats_id}/edit_photo
//chats/{chats_id}/clear
//chats/{chats_id}/notifications

// GET RESPONSE
/*
params := req.URL.Query()
name := params.Get("name")
fmt.Fprintf(w,"{'name':'%s'}",name)
*/

func chats(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		fmt.Println(req.FormValue("jwt"))
		//запрос в бд список чатов
		fmt.Fprintf(w,"{'response':'ok'}")
	}
}

func main(){
	http.HandleFunc("/chats",chats)
	

	http.ListenAndServe(":8090",nil)
}