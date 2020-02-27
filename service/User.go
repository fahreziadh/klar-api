package service

import (
	"encoding/json"
	"net/http"
	"strings"

	. "github.com/klar/models"
)

func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			var err ErrorResponse
			err.Code = "S004"
			err.HttpStatusCode = http.StatusInternalServerError
			err.Message = S004
			respondWithError(w, http.StatusBadRequest, err)
			return
		}
		fullname := user.FullName
		imageProfile := user.ImageProfile

		c := r.Header.Get("Authorization")
		tokenString := strings.Replace(c, "Bearer ", "", -1)
		parseJwt := ParseJWT(tokenString)
		if err := dao.ChangeProfile(parseJwt, fullname, imageProfile); err != nil {
			var err ErrorResponse
			err.Code = "S007"
			err.HttpStatusCode = http.StatusUpgradeRequired
			err.Message = S007
			respondWithError(w, http.StatusUpgradeRequired, err)
			return
		}
	}
}

func GetListStudent(w http.ResponseWriter, r *http.Request) {
	classId := r.URL.Query()["classid"][0]
	student, err := dao.FindStudentByClass(classId)

	if err != nil {
		var err ErrorResponse
		err.Code = "S008"
		err.HttpStatusCode = http.StatusNotFound
		err.Code = S008
		return
	}

	if len(student) == 0 {
		var err ErrorResponse
		err.Code = "S008"
		err.HttpStatusCode = http.StatusNotFound
		err.Code = S008
		return
	}

	respondWithJson(w, http.StatusOK, student)
}
