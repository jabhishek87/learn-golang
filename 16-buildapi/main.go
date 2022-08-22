package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB
// course of with slice of type Course struct
var courses []Course

//midleware, helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API Main")

	// Create initial rec on course
	courses = append(courses, Course{
		CourseId:    "123456787",
		CourseName:  "Python",
		CoursePrice: 299,
		Author:      &Author{FullName: "Abhishek", Website: "google.com"},
	})

	courses = append(courses, Course{
		CourseId:"2",CourseName:  "ReactJS",CoursePrice: 299,
		Author:&Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"},
	})
	courses = append(courses, Course{
		CourseId: "4", CourseName: "MERN Stack",CoursePrice: 199,
		Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}
	})

	r := mux.NewRouter()

	r.HandleFunc("/", serverHome)
	r.HandleFunc("/courses", getAllCources).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// Server Home  Route
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang API</h1>"))
}

func getAllCources(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
	return
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// get the ID form request
	params := mux.Vars(r)

	// loop through the courseID find the macthing idea and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course FOund with given ID")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body Is Empty, Set Some Data")
		return
	}

	// what if data is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Body Is Empty, Set Some Data")
		return
	}

	// generate a unique ID, and convert ID to String
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")

	// get the ID form request
	params := mux.Vars(r)

	// action
	// First find the index of slice
	// remove data from slice
	// add data to slice with the id
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO: Sned response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	// get the ID form request
	params := mux.Vars(r)

	// action
	// First find the index of slice
	// remove data from slice
	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			// var course Course
			// _ = json.NewDecoder(r.Body).Decode(&course)
			// course.CourseId = params["id"]
			// courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// TODO: Sned response when id is not found
}
