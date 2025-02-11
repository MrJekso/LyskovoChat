package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	_ "github.com/lib/pq"
)

const (
	USER     = "postgres"
	PASSWORD = "user"
	DBNAME   = "chat"
)

func health(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"{'response':'ok'}")
}

func registration(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		
		email := fmt.Sprint(req.FormValue("email"))
		login := fmt.Sprint(req.FormValue("login"))
		pass := fmt.Sprint(req.FormValue("password"))
		firstName := fmt.Sprint(req.FormValue("firstName"))
		lastName := fmt.Sprint(req.FormValue("lastName"))
		
		if email == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found email'}")
		}else if login == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found login'}")
		}else if pass == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found password'}")
		}else if firstName == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found first name'}")
		}else if lastName  == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found last name'}")
		}else{
			
			connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",USER,PASSWORD,DBNAME)
			
			db, err := sql.Open("postgres",connStr)
			if err != nil {
				panic(err)
			}
			defer db.Close()

			toDate    := time.Now()
			dateStr := fmt.Sprintf("%d-%d-%d",toDate.Day(),toDate.Month(),toDate.Year())
			reqData := fmt.Sprintf("insert into profiles (login,pass,firstname,lastname,date,email,jwt) values ('%s','%s','%s','%s','%s','%s','default')",login,pass,firstName,lastName,dateStr,email)
			_, err = db.Exec(reqData)
			if err != nil {
				panic(err)
			}
			
			reqSelect := fmt.Sprintf("select id from profiles where login='%s' and pass='%s'", login,pass)
			rows, err := db.Query(reqSelect)
			if err != nil {
				panic(err)
			}
			defer rows.Close()

			var idUser int
			
			for rows.Next() {
				err := rows.Scan(&idUser)
				if err != nil {
					panic(err)
				}
			}

			reqData = fmt.Sprintf("insert into profile (key) values ('%d')",idUser)
			_, err = db.Exec(reqData)
			if err != nil {
				panic(err)
			}
	
			fmt.Fprintf(w,"{'response':'ok'}")
		}
	}
}

func main(){
	http.HandleFunc("/registration",registration)
	http.HandleFunc("/health",health)

	http.ListenAndServe(":8090",nil)
}