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
    w.Header().Set("Content-Type", "application/json")
    
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

func GetMyDataProfile(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")

    userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
    if r.Method != "GET" {
        http.Error(w, "Unsupported http method", http.StatusBadRequest)
        return
    }

    response := data.Response{
        Status : 200,
        Message : "Success",
        Data : userInfo["User"],
    }
    json.NewEncoder(w).Encode(response)
}

func GetAllDataUser(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")

    userInfo := UserInfo(r.Context().Value("userInfo").(jwt.MapClaims))
    if r.Method != "GET" {
        http.Error(w, "Unsupported http method", http.StatusBadRequest)
        return
    }
    if(userInfo.Role != "superuser"){
        http.Error(w, "You Can't Access this, because you not superuser", http.StatusBadRequest)   
        return
    }

    allDataUser := data.GetAllDataUsers()

    response := data.Response{
        Status : 200,
        Message : "Success",
        Data : allDataUser,
    }

    log.Println(response)
    json.NewEncoder(w).Encode(response)
}

func UserInfo (val jwt.MapClaims) data.Users{
    dataUser := val["User"].(map[string]interface{})

    user := data.Users{
        Id : dataUser["id"].(float64),
        Username : dataUser["username"].(string),
        Role : dataUser["role"].(string),
    }

    return user
}