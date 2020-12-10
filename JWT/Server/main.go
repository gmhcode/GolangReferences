package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			//1. grabs the token
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				//2. Checks to see if the key is correct
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					//3. if its not correct, return
					return nil, fmt.Errorf("There was an error")
				}
				//3. returns the token... so mySigningKey becomes "token"
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			//4. if the token is valid, then we can pass the originally requested endpoint
			if token.Valid {
				endpoint(w, r)
			}
			//if there is no token then print this
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("My Simple Server")
	handleRequests()
}
