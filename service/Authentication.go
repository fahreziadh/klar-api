package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	. "github.com/klar/config"
	. "github.com/klar/dao"
	. "github.com/klar/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var config = Config{}
var dao = DAO{}

// Get Access Token
func GetAccessToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		defer r.Body.Close()
		var authRequest AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
			var err ErrorResponse
			err.Code = "S004"
			err.HttpStatusCode = http.StatusInternalServerError
			err.Message = S004
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		//find username
		user, err := dao.SignIn(authRequest.UserName)
		if err != nil {
			var err ErrorResponse
			err.Code = "S005"
			err.HttpStatusCode = http.StatusUnauthorized
			err.Message = S005
			respondWithError(w, http.StatusUnauthorized, err)
			return
		}

		isCorrect := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authRequest.Password))

		if isCorrect != nil {
			var err ErrorResponse
			err.Code = "S006"
			err.HttpStatusCode = http.StatusUnauthorized
			err.Message = S006
			respondWithError(w, http.StatusUnauthorized, err)
			return
		}

		var authResponse AuthResponse
		tokenString, _ := DecodeJWT(user.UserName)
		authResponse.AccessToken = tokenString
		user.Password = ""
		authResponse.User = user
		respondWithJson(w, http.StatusAccepted, authResponse)
	}
}

//Register
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			var err ErrorResponse
			err.Code = "S004"
			err.HttpStatusCode = http.StatusInternalServerError
			err.Message = S004
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		user.UserID = bson.NewObjectId()

		//email existing
		emailExists := dao.FindByEmail(user)
		if emailExists == nil {
			var err ErrorResponse
			err.Code = "S001"
			err.HttpStatusCode = http.StatusInternalServerError
			err.Message = S001
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		//username existing
		userNameExists := dao.FindByUserName(user)
		if userNameExists == nil {
			var err ErrorResponse
			err.Code = "S002"
			err.HttpStatusCode = http.StatusInternalServerError
			err.Message = S002
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		//set password before register
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)

		//register
		registerErr := dao.InsertUser(user)
		if registerErr != nil {
			var err ErrorResponse
			err.Code = "S003"
			err.HttpStatusCode = http.StatusInternalServerError
			err.Message = S003
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		//access token
		var authResponse AuthResponse
		tokenString, _ := DecodeJWT(user.UserName)
		authResponse.AccessToken = tokenString
		user.Password = ""
		authResponse.User = user
		respondWithJson(w, http.StatusAccepted, authResponse)
	}
}

func DecodeJWT(username string) (string, error) {
	expirationTime := time.Now().Add(1000 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Key))
	return tokenString, err
}

func respondWithError(w http.ResponseWriter, code int, msg ErrorResponse) {
	respondWithJson(w, code, msg)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ParseJWT(tokenString string) string {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		hmacSampleSecret := []byte(config.Key)
		return hmacSampleSecret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println(claims["username"])
	} else {
		fmt.Println(ok)
	}

	return claims["username"].(string)
}
