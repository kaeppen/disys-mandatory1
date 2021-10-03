package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

//hvis man virkelig er fresh, så skal nogle af felterne i struct være pointers, men det er ikke i denne opgaves scope

//THESE STRUCTS ARE JUST FOR SHOW 
type teacher struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Popularity float64 `json: "popularity"`
}

type student struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Courses []string `json:"courses"` 
}
//-----------------------

type course struct {
	ID       string   `json:"id"`
	Rating   float64  `json:"rating"`
	Name     string   `json:"name"`
	Workload float64  `json:"workload"`
	Students []string `json:"students"`
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", addCourse)
	router.DELETE("courses/:id", deleteCourseByID)
	router.PUT("courses/:id", updateCourse)
	router.Run("localhost:8080")
}

var courses = []course{
	{ID: "1", Rating: 9.9, Name: "Dømat", Workload: 9001.0, Students: []string{"Stefan", "Dagrun", "Marcus"}},
	{ID: "2", Rating: 5.2, Name: "BPAK", Workload: 0.0, Students: []string{}},
	{ID: "3", Rating: 10.0, Name: "GrPro", Workload: 15.0, Students: []string{"Christoffer", "Martin", "Henning"}},
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func addCourse(c *gin.Context) {
	var newCourse course

	// Call BindJSON to bind the received JSON to newCourse
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	// Add the new course to the slice.
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, gin.H{"Course":newCourse, "message":"successfully added"}) 
	return
}

func getCourseByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Course not found"})
}

func deleteCourseByID(c *gin.Context) {
	id := c.Param("id")
	var deleted course
	var oldLength = len(courses) //get the length of the array before (to be used when determining status)
	var newArray []course
	for i := 0; i < len(courses); i++ {
		if courses[i].ID == id {
			deleted = courses[i]
		} else {
			//append to the new array
			newArray = append(newArray, courses[i])
		}
	}
	//copy the new array into courses (also possible with copy() method?)
	courses = newArray

	//compare the length of the new array to the length of the old array
	//same length = nothing deleted :(
	if len(courses) == oldLength {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Course not found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"Course": deleted, "message":"successfully deleted"})
		return
	}
}

//kombination af tilføj og delete
func updateCourse(c *gin.Context) {
	var newCourse course

	//bind the recieved json to the new object
	if err := c.BindJSON(&newCourse); err != nil {
		fmt.Println(err.Error())
		return 
	}

	//iterate over all courses and replace the one in question
	id := c.Param("id")
	for i := 0; i < len(courses); i++ {
		if courses[i].ID == id {
			courses[i] = newCourse
			//all good
			c.IndentedJSON(http.StatusOK, gin.H{"Course": newCourse, "message": "successfully updated"})
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "something went wrong"})
	return 
}
