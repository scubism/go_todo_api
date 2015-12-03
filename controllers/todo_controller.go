package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexTodos(c *gin.Context) {
	todos, err := models.FindTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func ViewTodo(c *gin.Context) {
	todo, err := models.FindTodoById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := todo.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	todo, err := models.FindTodoById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = todo.Update(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	todo, err := models.FindTodoById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := todo.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func MoveTodo(c *gin.Context) {
	todo, err := models.FindTodoById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todoMove models.TodoMove
	if err = c.BindJSON(&todoMove); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = todo.Move(&todoMove); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
}
