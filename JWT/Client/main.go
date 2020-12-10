package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// set MY_JWT_TOKEN=mysupersecretphrase
//  var mySigningKey = os.Get("MY_JWT_TOKEN")
var mySiningKey = []byte("mysupersecretphrase")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	//1. Creates New Request
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
	//2. Appends valid token to the url header
	req.Header.Set("Token", validToken)
	//3. enters URL
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, string(body))
	// fmt.Fprintf(w, validToken)
}

// GenerateJWT - Generates a JWT webtoken
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Greg Hughes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySiningKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, err
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("My Simple Client")
	handleRequests()
	// tokenString, err := GenerateJWT()
	// if err != nil {
	// 	fmt.Println("Error generating token string")
	// }
	// fmt.Println(tokenString)
	// jwt.
}
