package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/maadiab/majalisulelm/core"
	Database "github.com/maadiab/majalisulelm/database"
)

var Authorized bool

var UserVerified string

func CheckUser(db *sqlx.DB, user Credentials) {

	var userCred core.User
	query := "SELECT name, password FROM users where name =$1 and password =$2"

	err := db.Get(&userCred, query, user.Username, user.Password)

	// user.Username, user.Password

	// log.Println("checking db")
	// log.Printf("%T", user.Username)

	// log.Println(user.Password)
	// fmt.Println(user.Username, user.Password)

	if err != nil {
		log.Println("Please Check Username and Password !!!", err)
		return
	}

	log.Println("User Verified Successfully ...")
	// Authorized = true

	// Login()
	// return Credentials.Username
}

var JwtKey = []byte("sectret_key")

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
		log.Println("Error Decoding user Data for login")
		return
	}

	CheckUser(Database.DB, cred)

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: UserVerified,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

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
