package Signin

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"../models"
	"github.com/dgrijalva/jwt-go"
)

var MyJWT = []byte("my_secret_key")

var user = map[string]string{
	"username1": "password1",
	"username2": "password2",
}

type Claim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var CredsModel models.Sign

	err := json.NewDecoder(r.Body).Decode(&CredsModel)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// check if password and username is valid
	expectPW, ok := user[CredsModel.Username]
	if !ok || expectPW != CredsModel.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Wrong Username or Password"))
		// return
	} else {
		w.WriteHeader(200)
		w.Write([]byte("Logged In"))
	}

	expTime := time.Now().Add(5 * time.Minute)

	claims := &Claim{
		Username: CredsModel.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tokens.SignedString(MyJWT)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "access-token",
		Value:   tokenString,
		Expires: expTime,
	})
	log.Printf(tokenString)

}
