package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// initialize DB
var DB []Course

// verify if course name and id
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hey! Welcome to the First API Home route</h1>"))
}

func GetAllDataFromDB(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All data have been fetched!")
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DB)
}

func GetCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Single course!")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, data := range DB {
		if data.CourseId == params["id"] {
			json.NewEncoder(w).Encode(data)
		}
	}
	json.NewEncoder(w).Encode("No Data found with this id!")
	return
}
