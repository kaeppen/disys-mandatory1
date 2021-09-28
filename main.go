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
	ID       string    `json:"id"`
	Rating   []float64 `json:"rating"`
	Name     string    `json:"name"`
	Students []string  `json:"students"`
}

type teacher struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Courses []string `json:"courses"`
}

func main() {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.POST("/students", postStudents)
	router.Run("localhost:8080")
}

var students = []student{
	{ID: "1", Name: "Jannick", Courses: []string{"Math"}},
	{ID: "2", Name: "Henning", Courses: []string{"RUC-fag", "How to get a job"}},
	{ID: "3", Name: "Ib", Courses: []string{"Science", "Math", "Computer Science"}},
}

func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

func postStudents(c *gin.Context) {
	var newStudent student

	// Call BindJSON to bind the received JSON to newStudent
	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	// Add the new student to the slice.
	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}
