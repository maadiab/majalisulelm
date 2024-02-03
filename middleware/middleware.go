package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

var Authorized = false

var UserVerified string

func CheckUser(db *sqlx.DB, user Credentials) {

	err := db.Get("SELECT EXISTS (SELECT 1 FROM users WHERE name = $1 AND password = $2)", user.Username, user.Password)

	log.Println("checking db")
	// fmt.Println(user.Username, user.Password)

	if err != nil {
		log.Println("Please Check Username and Password !!!")
	} else {

		Authorized = true

	}
	// UserVerified = Credentials.Username
	// return Credentials.Username
}

var jwtKey = []byte("sectret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {

	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: UserVerified,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}
