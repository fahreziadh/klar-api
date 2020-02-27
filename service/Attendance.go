package service

import (
	"encoding/json"
	"net/http"

	. "github.com/klar/models"
	"gopkg.in/mgo.v2/bson"
)

func InputAttendance(w http.ResponseWriter, r *http.Request) {
	var attendance Attendance

	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		var err ErrorResponse
		err.Code = "S004"
		err.HttpStatusCode = http.StatusBadRequest
		err.Message = S004
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	attendance.AttendanceID = bson.NewObjectId()

	if err := dao.InsertAttendance(attendance); err != nil {
		var err ErrorResponse
		err.Code = "S008"
		err.HttpStatusCode = http.StatusForbidden
		err.Message = S008
		respondWithError(w, http.StatusForbidden, err)
		return
	}

	respondWithJson(w, http.StatusOK, "")
}

func GetListAttendance(w http.ResponseWriter, r *http.Request) {
	scheduleid := r.URL.Query()["scheduleid"][0]
	attendance, err := dao.FindAttendanceBySchedule(scheduleid)
	if err != nil {
		var err ErrorResponse
		err.Code = "S008"
		err.HttpStatusCode = http.StatusForbidden
		err.Message = S008
		respondWithError(w, http.StatusForbidden, err)
		return
	}
	if attendance == nil {
		var err ErrorResponse
		err.Code = "S009"
		err.HttpStatusCode = http.StatusNotFound
		err.Message = S009
		respondWithError(w, http.StatusNotFound, err)
		return
	}
	respondWithJson(w, http.StatusOK, attendance)
}
