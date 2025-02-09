package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"strings"
	_ "github.com/lib/pq"
	"github.com/golang-jwt/jwt/v5"
)

const (
	USER     = "postgres"
	PASSWORD = "user"
	DBNAME   = "chat"
	HOST     = "localhost"
)

var secretKey = []byte("my-secret-key")

func health(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"{'response':'OK'}")
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 48).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error token: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}
	return token, nil
}
				
func authorization(w http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		
		jwtClient := fmt.Sprint(req.FormValue("jwt"))
	
		if jwtClient == "" {
			fmt.Fprintf(w,"{'response':'error','message':'not found token'}")
		} else {
			
			connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",USER,PASSWORD,DBNAME,HOST)
			
			db, err := sql.Open("postgres",connStr)
			if err != nil {
				panic(err)
			}
			defer db.Close()

			reqSelect := fmt.Sprintf("select login from profiles where jwt='%s'", jwtClient)
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
				
				token, err := verifyToken(jwtClient)
				if err != nil {
					fmt.Printf("Ошибка при верификации токена: %v\n", err)
					fmt.Fprintf(w,"{response: 'error'}")
				}

				var tokenTime int64
				// Извлекаем claims (утверждения) из токена
				if claims, ok := token.Claims.(jwt.MapClaims); ok {
					tokenTime = int64(claims["exp"].(float64))
				} else {
					fmt.Fprintf(w,"{response: 'error'}")
				}

				resTime := tokenTime - time.Now().Add(time.Hour).Unix()
				if resTime <= 82800 && resTime >= 1 {

					key, err := createToken(isLogin)
					
					if err != nil {
						panic(err)
					}

					sqlReq := fmt.Sprintf("update profiles set jwt = '%s' where login='%s' and jwt='%s'", key, isLogin, jwtClient)

					_, err = db.Exec(sqlReq)	
					if err != nil {
						panic(err)
					}

				}
				if resTime <= 0 {
					fmt.Fprintf(w,"{response: 'error', msg: 'token expired'}")
				} else {
					fmt.Fprintf(w,"{response: 'OK'}")
				}
			} else {
				fmt.Fprintf(w, "{error: 'is not valid login or password'}")
			}
		}
	}
}

func main(){
	http.HandleFunc("/authorization",authorization)
	http.HandleFunc("/health",health)

	http.ListenAndServe(":8090",nil)
}