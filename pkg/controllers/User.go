package controllers

import(
	"net/http"
	"encoding/json"

	"github.com/irvandandung/goAPI/pkg/data"
    "log"

	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
    jwt.StandardClaims
    User data.Users
}

var APPLICATION_NAME = "My Simple API with JWT"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("signature of my API")

func Login(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
        http.Error(w, "Unsupported http method", http.StatusBadRequest)
        return
    }

    username, password, ok := r.BasicAuth()
    if !ok {
        http.Error(w, "Invalid username or password", http.StatusBadRequest)
        return
    }

    ok, userInfo := data.GetDataUser(username, password)
    if !ok {
        http.Error(w, "Invalid username or password", http.StatusBadRequest)
        return
    }

    claims := MyClaims{
        StandardClaims: jwt.StandardClaims{
            Issuer:    APPLICATION_NAME,
            ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
        },
        User : userInfo,
    }
    log.Println(claims)

    token := jwt.NewWithClaims(
        JWT_SIGNING_METHOD,
        claims,
    )

    signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": signedToken})
}