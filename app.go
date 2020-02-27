package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	. "github.com/klar/config"
	. "github.com/klar/service"

	"github.com/urfave/negroni"
)

var config = Config{}

func main() {
	NewRouter()
}

func NewRouter() {
	config.Read()

	r := mux.NewRouter()
	ar := mux.NewRouter()

	mw := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Key), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	//Extended Router
	AuthHandler(r)
	UserHandler(ar)
	ClassHandler(ar)
	ScheduleHandler(ar)
	AttendanceHandler(ar)
	StudentHandler(ar)

	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(ar))
	r.PathPrefix("/api").Handler(an)

	n := negroni.Classic()
	n.UseHandler(r)

	s := &http.Server{
		Addr:           ":3000",
		Handler:        n,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

func AuthHandler(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	r.HandleFunc("/api/signin", GetAccessToken)
	r.HandleFunc("/api/register", Register)
}

func UserHandler(r *mux.Router) {
	r.HandleFunc("/api/user", ChangeProfile)
}

func ClassHandler(r *mux.Router) {
	r.HandleFunc("/api/class", CreateClass).Methods("POST")
	r.HandleFunc("/api/class/join", JoinClass).Methods("POST")
	r.HandleFunc("/api/class", GetListClass).Methods("GET")
	r.HandleFunc("/api/class/{id}", GetDetailClass).Methods("GET")
}

func ScheduleHandler(r *mux.Router) {
	r.HandleFunc("/api/schedule", CreateSchedule).Methods("POST")
	r.HandleFunc("/api/schedule", GetSchedule).Methods("GET")
}

func AttendanceHandler(r *mux.Router) {
	r.HandleFunc("/api/attendance", GetListAttendance).Methods("GET")
	r.HandleFunc("/api/attendance", InputAttendance).Methods("POST")
}

func StudentHandler(r *mux.Router) {
	r.HandleFunc("/api/student", GetListStudent).Methods("GET")

}
