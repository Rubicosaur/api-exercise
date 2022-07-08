package main

// Simple REST API
//To do list

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type toDo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var toDos = []toDo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read book", Completed: false},
	{ID: "3", Item: "Record video", Completed: false},
}

func getToDos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, toDos)
}
func addToDo(context *gin.Context) {
	var newToDo toDo
	if err := context.BindJSON(&newToDo); err != nil {
		return
	}

	toDos = append(toDos, newToDo)

	context.IndentedJSON(http.StatusCreated, newToDo)
}

func getToDo(context *gin.Context) {
	id := context.Param("id")
	toDo, err := getToDoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "to do not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, toDo)

}

func getToDoById(id string) (*toDo, error) {
	for i, t := range toDos {
		if t.ID == id {
			return &toDos[i], nil
		}
	}

	return nil, errors.New("todo not found")

}
func toggleToDoStatus(context *gin.Context) {
	id := context.Param("id")
	toDo, err := getToDoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"messag": "to do not found"})
		return
	}

	toDo.Completed = !toDo.Completed
	context.IndentedJSON(http.StatusOK, toDo)
}

func main() {

	router := gin.Default()
	router.GET("/todos", getToDos)
	router.GET("/todos/:id", getToDo)
	router.PATCH("/todos/:id", toggleToDoStatus)
	router.POST("/todos", addToDo)
	router.Run("localhost:8081")

}
