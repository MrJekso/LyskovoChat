package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"net/http"
	"strings"
	"time"
)

const (
	USER     = "postgres"
	PASSWORD = "user"
	DBNAME   = "chat"
)

var secretKey = []byte("my-secret-key")

func createToken(username string) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"username" : username,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString,err
}


func health(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "{'response':'ok'}")
}

func authentication(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":

		login := fmt.Sprint(req.FormValue("login"))
		pass := fmt.Sprint(req.FormValue("password"))

		if login == "" {
			fmt.Fprintf(w, "{'response':'error','message':'not found login'}")
		} else if pass == "" {
			fmt.Fprintf(w, "{'response':'error','message':'not found password'}")
		} else {
			connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASSWORD, DBNAME)
			fmt.Println(connStr)
			db, err := sql.Open("postgres", connStr)
			if err != nil {
				panic(err)
			}
			defer db.Close()

			reqSelect := fmt.Sprintf("select jwt from profiles where login='%s' and pass='%s'", login, pass)
			fmt.Println(reqSelect)
			rows, err := db.Query(reqSelect)
			if err != nil {
				panic(err)
			}
			defer rows.Close()

			jwt := ""
			for rows.Next() {
				err := rows.Scan(&jwt)
				if err != nil {
					panic(err)
				}
				fmt.Println(jwt)
			}
			jwt = strings.Replace(jwt, " ", "", -1)
			//запрос в бд список чатов
			if jwt == "default" {
				//ЕСЛИ ЗНАЧЕНИЕ ПО УМОЛЧАНИЮ
				//СГЕНЕРИРОВАТЬ НОВОЕ
				key, err := createToken(login)
				if err != nil {
					panic(err)
				}
				fmt.Fprintf(w, "{'response':'%s'}", key)
			} else if jwt == "" {
				fmt.Fprintf(w, "{'response':'Not found'}")
			} else {
				fmt.Fprintf(w, "{'response':'%s'}", jwt)
			}

		}
	}
}

func main() {
	http.HandleFunc("/authentication", authentication)
	http.HandleFunc("/health", health)

	http.ListenAndServe(":8090", nil)
}
