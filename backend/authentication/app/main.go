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

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
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
			db, err := sql.Open("postgres", connStr)
			if err != nil {
				panic(err)
			}
			defer db.Close()

			reqSelect := fmt.Sprintf("select login from profiles where login='%s' and pass='%s'", login, pass)
			rows, err := db.Query(reqSelect)
			if err != nil {
				panic(err)
			}
			defer rows.Close()

			var isLogin string

			for rows.Next() {
				err := rows.Scan(&isLogin)
				if err != nil {
					panic(err)
				}
			}

			isLogin = strings.Replace(isLogin, " ", "", -1)

			if isLogin != "" {

				key, err := createToken(login)
				if err != nil {
					panic(err)
				}
				sqlReq := fmt.Sprintf("update profiles set jwt = '%s' where login='%s' and pass='%s'", key, login, pass)
				_, err = db.Exec(sqlReq)
				if err != nil {
					panic(err)
				}
				fmt.Fprintf(w, "{'response':'%s'}", key)
			} else {
				fmt.Fprintf(w, "{error: 'is not valid login or password'}")
			}

		}
	}
}

func main() {
	http.HandleFunc("/authentication", authentication)
	http.HandleFunc("/health", health)

	http.ListenAndServe(":8090", nil)
}
