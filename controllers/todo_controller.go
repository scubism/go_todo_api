package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net/http"
	"todo_center/go_todo_api/app"
	"todo_center/go_todo_api/models"
	"todo_center/go_todo_api/utils"
)

func IndexTodos(c *gin.Context) {
	todos, err := models.FindTodos(app.GetDB(c))
	if err != nil {
		utils.AbortWithPublicError(c, http.StatusInternalServerError, err, "Couldn't find todos")
		return
	}

	c.JSON(http.StatusOK, todos)
}

func ViewTodo(c *gin.Context) {
	todo, err := models.FindTodoById(app.GetDB(c), c.Param("id"))
	if err != nil {
		if err == mgo.ErrNotFound {
			utils.AbortWithPublicError(c, http.StatusUnauthorized, err, "You can't access to the todo")
		} else {
			utils.AbortWithPublicError(c, http.StatusInternalServerError, err, "Couldn't find the todo")
		}
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {
		return
	}

	if err := todo.Create(app.GetDB(c)); err != nil {
		utils.AbortWithPublicError(c, http.StatusInternalServerError, err, "Couldn't create the todo")
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	todo, err := models.FindTodoById(app.GetDB(c), c.Param("id"))
	if err != nil {
		if err == mgo.ErrNotFound {
			utils.AbortWithPublicError(c, http.StatusUnauthorized, err, "You can't access to the todo")
		} else {
			utils.AbortWithPublicError(c, http.StatusInternalServerError, err, "Couldn't find the todo")
		}
		return
	}

	_Id := todo.Id
	if err = c.BindJSON(&todo); err != nil {
		return
	}
	todo.Id = _Id

	if err = todo.Update(app.GetDB(c)); err != nil {
		utils.AbortWithPublicError(c, http.StatusInternalServerError, err, "Couldn't update the todo")
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	todo, err := models.FindTodoById(app.GetDB(c), c.Param("id"))
	if err != nil {
		if err == mgo.ErrNotFound {
			utils.AbortWithPublicError(c, http.StatusUnauthorized, err, "You can't access to the todo")
		} else {
			utils.AbortWithPublicError(c, http.StatusInternalServerError, err, "Couldn't find the todo")
		}
		return
	}

	if err := todo.Delete(app.GetDB(c)); err != nil {
		utils.AbortWithPublicError(c, http.StatusInternalServerError, err, "Couldn't delete the todo")
		return
	}

	c.JSON(http.StatusOK, todo)
}

func MoveTodo(c *gin.Context) {
	todo, err := models.FindTodoById(app.GetDB(c), c.Param("id"))
	if err != nil {
		if err == mgo.ErrNotFound {
			utils.AbortWithPublicError(c, http.StatusUnauthorized, err, "You can't access to the todo")
		} else {
			utils.AbortWithPublicError(c, http.StatusInternalServerError, err, "Couldn't find the todo")
		}
		return
	}

	var todoMove models.TodoMove
	if err = c.BindJSON(&todoMove); err != nil {
		return
	}

	if err = todo.Move(app.GetDB(c), &todoMove); err != nil {
		utils.AbortWithPublicError(c, http.StatusInternalServerError, err, "Couldn't move the todo")
		return
	}

	c.JSON(http.StatusOK, todo)
}
