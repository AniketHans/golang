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

//Model for course
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fake db
var Courses []Course

//Helper functions
func (c *Course)IsEmpty() bool{
	// return c.CourseId == "" && c.CourseName==""
	return c.CourseName==""
}

func main(){
	fmt.Println("Apis")
	r := mux.NewRouter()

	Courses = append(Courses, Course{CourseId: "1", CourseName: "Python", CoursePrice: 599, Author: &Author{Fullname: "Aniket Hans",Website: "anikethans.com"}})

	r.HandleFunc("/",serveHome).Methods("Get")
	r.HandleFunc("/courses",getAllCourses).Methods("Get")
	r.HandleFunc("/course/{id}",getCourseById).Methods("Get")
	r.HandleFunc("/course",addCourse).Methods("Post")
	r.HandleFunc("/course/{id}",updateCourseById).Methods("Put")
	r.HandleFunc("/course/{id}",deleteCourseById).Methods("Delete")

	log.Fatal(http.ListenAndServe(":4001",r))

}

//controllers

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to the APIs</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get All Courses")

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func getCourseById(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get One Course")
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, course := range Courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
}


func addCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Add a Course")
	w.Header().Set("Content-type", "application/json")

	if r.Body==nil{
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data found")
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode("Course added")
}


func updateCourseById(w http.ResponseWriter,r *http.Request){
	fmt.Println("Updating the course by course id")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	for index, course := range Courses{
		if course.CourseId == params["id"]{
			Courses = append(Courses[:index],Courses[index+1:]...)
			var newCourse Course
			json.NewDecoder(r.Body).Decode(&newCourse)
			newCourse.CourseId= params["id"]
			Courses = append(Courses, newCourse)
			json.NewEncoder(w).Encode("Course updated")
			return 
		}
	}

	json.NewEncoder(w).Encode("No course found with the id")
}

func deleteCourseById(w http.ResponseWriter,r *http.Request){
	fmt.Println("Deleting the course by course id")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	var deletedCourse *Course
	for index, course := range Courses{
		if course.CourseId == params["id"]{
			Courses = append(Courses[:index], Courses[index+1:]...)
			deletedCourse = &course
			break
		}
	}
	if deletedCourse != nil{
		json.NewEncoder(w).Encode(*deletedCourse)
	}else{
	json.NewEncoder(w).Encode("No course found with the id")
	}
}