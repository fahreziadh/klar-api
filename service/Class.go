package service

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	. "github.com/klar/models"
	gonanoid "github.com/matoous/go-nanoid"
	"gopkg.in/mgo.v2/bson"
)

func CreateClass(w http.ResponseWriter, r *http.Request) {
	var class Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		var err ErrorResponse
		err.Code = ""
		err.HttpStatusCode = 1
		err.Message = ""
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	class.ID = bson.NewObjectId()
	id, _ := gonanoid.Generate("1234567890", 6)
	class.ClassID = "K" + id
	if err := dao.InsertClass(class); err != nil {
		var e ErrorResponse
		e.Code = "S008"
		e.HttpStatusCode = http.StatusInternalServerError
		e.Message = S008
		respondWithError(w, http.StatusInternalServerError, e)
		return
	}

	respondWithJson(w, http.StatusOK, class)
}

func GetListClass(w http.ResponseWriter, r *http.Request) {
	var (
		class []Class
	)

	c := r.Header.Get("Authorization")
	tokenString := strings.Replace(c, "Bearer ", "", -1)
	username := ParseJWT(tokenString)

	//Find By Owner
	classOwned, classErr := dao.FindClassByOwner(username)
	if classErr != nil {
		var e ErrorResponse
		e.Code = "S009"
		e.HttpStatusCode = http.StatusNotFound
		e.Message = S009
		respondWithError(w, http.StatusNotFound, e)
		return
	}
	class = append(classOwned)

	//Find by student
	student, sterr := dao.FindClassByStudent(username)

	if sterr != nil {
		var err ErrorResponse
		err.Code = "S009"
		err.HttpStatusCode = http.StatusNotFound
		err.Message = S009
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	totalStudent := len(student)

	for i := 0; i < totalStudent; i++ {
		s, _ := dao.FindClassByClassId(student[i].ClassID)
		class = append(class, s)
	}

	if len(class) == 0 {
		var err ErrorResponse
		err.Code = "S009"
		err.HttpStatusCode = http.StatusNotFound
		err.Message = S009
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	respondWithJson(w, http.StatusOK, class)
}

func JoinClass(w http.ResponseWriter, r *http.Request) {
	var joinClassRequest JoinClassRequest
	if err := json.NewDecoder(r.Body).Decode(&joinClassRequest); err != nil {
		var err ErrorResponse
		err.Code = ""
		err.HttpStatusCode = 1
		err.Message = ""
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	c := r.Header.Get("Authorization")
	tokenString := strings.Replace(c, "Bearer ", "", -1)
	username := ParseJWT(tokenString)

	class, err := dao.FindClassByClassId(joinClassRequest.ClassID)
	user, userErr := dao.FindUserProfile(username)

	if err != nil || userErr != nil {
		var e ErrorResponse
		e.Code = "S009"
		e.HttpStatusCode = http.StatusNotFound
		e.Message = err.Error()
		respondWithError(w, http.StatusNotFound, e)
		return
	}

	var student Student
	student.ClassID = class.ClassID
	student.FullName = user.FullName
	student.ImageProfile = user.ImageProfile
	student.UserName = user.UserName
	errStudent := dao.InsertStudent(student)

	if errStudent != nil {
		var e ErrorResponse
		e.Code = "S012"
		e.HttpStatusCode = http.StatusForbidden
		e.Message = errStudent.Error()
		respondWithError(w, http.StatusForbidden, e)
		return
	}

	respondWithJson(w, http.StatusOK, student)
}

func GetDetailClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	class, err := dao.FindClassByClassId(id)

	if err != nil {
		var err ErrorResponse
		err.Code = "S009"
		err.HttpStatusCode = http.StatusNotFound
		err.Message = S009
		respondWithError(w, http.StatusNotFound, err)
		return
	}
	respondWithJson(w, http.StatusOK, class)
}
