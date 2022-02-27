package main

import (
	"net/http"
	"encoding/json"
	"sort"
	"github.com/gorilla/mux"
	"strconv"
)

func main() {
	http.ListenAndServe(":12345", MakeWebHandler())
}

type Student struct {
	Id int
	Name string
	Age int
	Score int
}

type Students []Student
func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

var (
	students map[int]Student
	lastId int
)

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	mux.HandleFunc("/students/{id:[0-9]+}", GetSpecificStudentHandler).Methods("GET")

	students = make(map[int]Student)
	students[1] = Student{1, "bconfiden2", 27, 100}
	students[2] = Student{2, "golang", 10, 50}
	students[3] = Student{3, "test", 1000, 10}
	
	lastId = 3
	return mux
}

func GetSpecificStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}