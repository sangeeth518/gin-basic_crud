package main

import (
	"crud/database"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Phno string `json:"phno"`
}

var todos = []todo{
	{Id: "01", Name: "sangeeth", Phno: "7994735007"},
}

func get(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
func createPost(context *gin.Context) {
	var newPost todo
	err := context.BindJSON(&newPost)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "enter valid input"})
		return
	}
	todos = append(todos, newPost)
	context.IndentedJSON(http.StatusCreated, newPost)

}
func postById(id string) (*todo, error) {
	for index, val := range todos {
		if val.Id == id {
			return &todos[index], nil
		}
	}
	return nil, errors.New("post not found")

}

func updatePost(context *gin.Context) {
	id := context.Param("id")
	newtodo, err := postById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}
	var newPost todo
	error := context.BindJSON(&newPost)
	if error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "enter valid input"})
		return
	}
	newtodo.Name = newPost.Name
	newtodo.Phno = newPost.Phno

	context.IndentedJSON(http.StatusOK, newtodo)

}

func deletepost(context *gin.Context) {
	id := context.Param("id")
	for index, val := range todos {
		if val.Id == id {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "post detetd"})
}
func getPost(context *gin.Context) {
	id := context.Param("id")
	newtodo, err := postById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, newtodo)
}

func init() {
	db := database.Connect()
	fmt.Println(db)
	// db.AutoMigrate(&todo{})
	// resp := db.Create(&todos)
	// fmt.Println(resp.Error, resp.RowsAffected)

}

func main() {
	router := gin.Default()

	router.GET("/todo", get)
	router.POST("/post", createPost)
	router.PATCH("/post/:id", updatePost)
	router.DELETE("/post/:id", deletepost)
	router.GET("/post/:id", getPost)
	router.Run("localhost:8080")
}
