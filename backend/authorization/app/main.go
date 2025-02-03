package main

import (
	//"database/sql"
	"fmt"
	"net/http"
	//"time"
	//_ "github.com/lib/pq"
)

const (
	USER     = "postgres"
	PASSWORD = "user"
	DBNAME   = "chat"
)

func health(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.Header.Get("host"))
	fmt.Fprintf(w,"{'response':'%s'}",req)
}

/*
func registration(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		
		login := fmt.Sprint(req.FormValue("login"))
	
		if login == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found login'}")
		
		}else{
			
			connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",USER,PASSWORD,DBNAME)
			fmt.Println(connStr)
			db, err := sql.Open("postgres",connStr)
			if err != nil {
				panic(err)
			}
			defer db.Close()

			toDate    := time.Now()
			dateStr := fmt.Sprintf("%d-%d-%d",toDate.Day(),toDate.Month(),toDate.Year())
			reqData := fmt.Sprintf("insert into profiles (login,pass,firstname,lastname,date,email,jwt) values ('%s','%s','%s','%s','%s','%s','default')",login,password,firstName,lastName,dateStr,email)
			fmt.Println(reqData)
 			result, err := db.Exec(reqData)
			if err != nil {
				panic(err)
			}
			fmt.Println(result.RowsAffected())
			//запрос в бд список чатов
			
			fmt.Fprintf(w,"{'response':'ok'}")
		}
	}
}
*/
func main(){
	//http.HandleFunc("/",registration)
	http.HandleFunc("/health",health)

	http.ListenAndServe(":8090",nil)
}