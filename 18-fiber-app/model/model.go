package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

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

// func Dbinit(ctx *fiber.Ctx) error {
// 	message := "Records Get All"
// 	return ctx.Status(fiber.StatusOK).SendString(message)
// }

func CreateSeedRecords() map[string]Course {

	fmt.Println("Adding Data")
	courses := make(map[string]Course)

	courses["123456787"] = Course{
		CourseId: "123456787", CourseName: "Python", CoursePrice: 299,
		Author: &Author{FullName: "Abhishek", Website: "google.com"},
	}

	courses["2"] = Course{
		CourseId: "2", CourseName: "ReactJS", CoursePrice: 299,
		Author: &Author{FullName: "Hitesh", Website: "lco.dev"},
	}
	courses["4"] = Course{
		CourseId: "4", CourseName: "MERN Stack", CoursePrice: 299,
		Author: &Author{FullName: "Abhi", Website: "lco.dev"},
	}
	return courses
}

func GetNewID() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(100))
}
