package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	. "github.com/klar/models"
	"gopkg.in/mgo.v2/bson"
)

//Create new Schedule

func CreateSchedule(w http.ResponseWriter, r *http.Request) {
	var scheduleRequest ScheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&scheduleRequest); err != nil {
		var e ErrorResponse
		e.Code = "S004"
		e.HttpStatusCode = http.StatusBadRequest
		e.Message = S004
		respondWithError(w, http.StatusBadRequest, e)
		ShowError(err.Error())
		return
	}
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, scheduleRequest.StartDate)
	endDate := startDate.AddDate(0, scheduleRequest.TotalMonth, 0)
	diff := endDate.Sub(startDate).Hours() / 24
	weekDayRequest := scheduleRequest.DetailTime

	if err != nil {
		var e ErrorResponse
		e.Code = "S010"
		e.HttpStatusCode = http.StatusBadRequest
		e.Message = S010
		respondWithError(w, http.StatusBadRequest, e)
		ShowError(err.Error())
		return
	}

	var wg sync.WaitGroup
	wg.Add(int(diff))
	fmt.Println("Running for loop…")
	//Create Schedule
	for i := 0; i < int(diff); i++ {
		go func(i int) {
			defer wg.Done()
			form := "3 4 PM"
			date := startDate.AddDate(0, 0, i-1)

			//data you needed
			weekday := date.Weekday().String()
			var finalTime = ""
			finalDate := date.String()

			//check days schedule
			if weekday == "Sunday" {
				detailTime, err := time.Parse(form, weekDayRequest.Sunday)
				if err == nil {
					finalTime = timeFormat(detailTime.Hour(), detailTime.Minute())
				}
			} else if weekday == "Monday" {
				detailTime, err := time.Parse(form, weekDayRequest.Monday)
				if err == nil {
					finalTime = timeFormat(detailTime.Hour(), detailTime.Minute())
				}
			} else if weekday == "Tuesday" {
				detailTime, err := time.Parse(form, weekDayRequest.Tuesday)
				if err == nil {
					finalTime = timeFormat(detailTime.Hour(), detailTime.Minute())
				}
			} else if weekday == "Wednesday" {
				detailTime, err := time.Parse(form, weekDayRequest.Wednesday)
				if err == nil {
					finalTime = timeFormat(detailTime.Hour(), detailTime.Minute())
				}
			} else if weekday == "Thursday" {
				detailTime, err := time.Parse(form, weekDayRequest.Thursday)
				if err == nil {
					finalTime = timeFormat(detailTime.Hour(), detailTime.Minute())
				}
			} else if weekday == "Friday" {
				detailTime, err := time.Parse(form, weekDayRequest.Friday)
				if err == nil {
					finalTime = timeFormat(detailTime.Hour(), detailTime.Minute())
				}
			} else if weekday == "Saturday" {
				detailTime, err := time.Parse(form, weekDayRequest.Saturday)
				if err == nil {
					finalTime = timeFormat(detailTime.Hour(), detailTime.Minute())
				}
			}
			// schedule := Schedule(, , finalDate, ,)
			if finalTime != "" {
				var schedule Schedule
				schedule.ScheduleID = bson.NewObjectId()
				schedule.Class = scheduleRequest.Class
				schedule.StartDate = finalDate
				schedule.StartTime = finalTime
				if err := dao.InsertSchedule(schedule); err != nil {
					var e ErrorResponse
					e.Code = "S011"
					e.HttpStatusCode = http.StatusInternalServerError
					e.Message = S011
					ShowError(err.Error())
					respondWithError(w, http.StatusInternalServerError, e)
					return
				}
			}
		}(i)
	}
	fmt.Println("Waiting for loop…")
	wg.Wait()
	fmt.Println("Finish for loop…")

	respondWithJson(w, http.StatusOK, "")
}

func ShowError(s string) {
	fmt.Println(s)
}

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	var isToday bool
	if r.URL.Query()["today"][0] == "true" {
		isToday = true
	} else {
		isToday = false
	}
	c := r.Header.Get("Authorization")
	tokenString := strings.Replace(c, "Bearer ", "", -1)
	username := ParseJWT(tokenString)
	var (
		schedule []Schedule
		err      error
		class    []Class
	)

	//Find By Class Owner------------------------------------------------------
	class, err = dao.FindClassByOwner(username)
	classTotal := len(class) - 1

	if err != nil {
		var err ErrorResponse
		err.Code = "S009"
		err.HttpStatusCode = http.StatusNotFound
		err.Message = S009 + " class"
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	if class != nil {
		var wg sync.WaitGroup
		wg.Add(classTotal)
		for i := 0; i < classTotal; i++ {
			go func(i int) {
				defer wg.Done()
				s, err := dao.FindAllSchedule(isToday, username, class[i].ClassID)

				if err != nil {
					var e ErrorResponse
					e.Code = "S009"
					e.HttpStatusCode = http.StatusNotFound
					e.Message = err.Error()
					respondWithError(w, http.StatusNotFound, e)
					return
				}

				if s != nil {
					schedule = append(s)
				}

			}(i)
		}
		wg.Wait()
	}

	// Find as a Student
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
		s, _ := dao.FindAllSchedule(isToday, username, student[i].ClassID)
		schedule = append(s)
	}

	if schedule == nil {
		var err ErrorResponse
		err.Code = "S009"
		err.HttpStatusCode = http.StatusNotFound
		err.Message = S009
		respondWithError(w, http.StatusNotFound, err)
		return
	}

	respondWithJson(w, http.StatusOK, schedule)

}

func timeFormat(hour int, minute int) string {
	var m string
	if minute < 10 {
		m = "0" + strconv.Itoa(minute)
	} else {
		m = strconv.Itoa(minute)
	}
	return strconv.Itoa(hour) + "-" + m + "-00"
}
