package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type student struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Courses []string `json:"courses"` //skal muligvis senere relatere til course "objekter"? samme for de andre arrays
	//how do we define course workloads?
}

type course struct {
	ID       string   `json:"id"`
	Rating   float64  `json:"rating"`
	Name     string   `json:"name"`
	Students []string `json:"students"`
}

type teacher struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Courses []string `json:"courses"`
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCourses)

	router.Run("localhost:8080")
}

var courses = []course{
	{ID: "1", Rating: 9.9, Name: "Dømat", Students: []string{"Stefan", "Dagrun", "Marcus"}},
	{ID: "2", Rating: 5.2, Name: "BPAK", Students: []string{}},
	{ID: "3", Rating: 10.0, Name: "GrPro", Students: []string{"Christoffer", "Martin", "Henning"}},
}

var students = []student{
	{ID: "1", Name: "Jannick", Courses: []string{"Math"}},
	{ID: "2", Name: "Henning", Courses: []string{"RUC-fag", "How to get a job"}},
	{ID: "3", Name: "Ib", Courses: []string{"Science", "Math", "Computer Science"}},
}

func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func postCourses(c *gin.Context) {
	var newCourse course

	// Call BindJSON to bind the received JSON to newStudent
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	// Add the new student to the slice.
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}
