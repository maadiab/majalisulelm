package Middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"github.com/maadiab/majalisulelm/core"
	Database "github.com/maadiab/majalisulelm/database"
	"golang.org/x/crypto/bcrypt"
)

// var Authorized bool

var UserPerms []string

func ComparePassword(hashedPassword []byte, inputPassword string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(inputPassword))

}

func CheckUser(ctx context.Context, db *sqlx.DB, user Credentials) {

	var userCred core.User

	var inputPassword = user.Password

	// hashedPassword, err := helper.HashPassword(inputPassword)

	log.Println(user.Username)
	log.Println(inputPassword)
	// log.Println(string(hashedPassword))

	var hashedPassword string
	query0 := "SELECT password FROM users where name =$1"

	err := db.Get(&hashedPassword, query0, user.Username)

	if err != nil {
		log.Println("Please Check Username and Password !!!", err)
		log.Println(hashedPassword)
		return
	}

	err = ComparePassword([]byte(hashedPassword), user.Password)
	if err != nil {
		log.Println("Error Hashing password !!!", err)
	}

	// err = ComparePassword(hashedPassword, string(inputPassword))
	// if err != nil {
	// 	log.Println("Password not same !!!", err)
	// }
	query := "SELECT * FROM users where name =$1 and password =$2"

	err = db.Get(&userCred, query, user.Username, string(hashedPassword))

	if err != nil {
		log.Println("Please Check Username and Password !!!", err)
		log.Println(userCred)
		return
	}

	var userPermissions []string
	err = db.Select(&userPermissions, "SELECT permission_value FROM permissions WHERE permission_name = $1", userCred.Permissions)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("User Verified Successfully ...")

	// UserPerms = userPermissions
	log.Println(userPermissions)
	UserPerms = userPermissions

}

var JwtKey = []byte("secret_key")

// type Permissions struct {
// 	UserType  string `json:"permission_name"`
// 	Userperms string `json:"permission_value"`
// }

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error Decoding user Data for login !!!")
		return
	}

	CheckUser(r.Context(), Database.DB, cred)

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username:    cred.Username,
		Permissions: UserPerms,
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

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "refreshed_token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
		}

		// w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))

		// claims, ok := tkn.Claims.(*Claims)
		// if !ok {
		// 	log.Println(ok)
		// }

		log.Println("Hello, ", claims.Username)

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))

		// next.RequiredPermissions
		// log.Println(next.RequiredPermissions)

	}

}
