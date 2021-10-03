package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
//hvis man virkelig er fresh, så skal nogle af felterne i struct være pointers, men det er ikke i denne opgaves scope

type course struct {
	ID       string   `json:"id"`
	Rating   float64  `json:"rating"`
	Name     string   `json:"name"`
	Students []string `json:"students"`
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


func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func postCourses(c *gin.Context) {
	var newCourse course

	// Call BindJSON to bind the received JSON to newCourse
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	// Add the new course to the slice.
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func getCourseByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}
