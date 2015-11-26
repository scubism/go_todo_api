package models

import (
	"../app"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Todo struct {
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title string        `json:"title" bson:"title" binding:"required"`
}

func FindTodos() ([]Todo, error) {
	var todos []Todo
	err := app.DB.C("todos").
		Find(bson.M{}).
		Limit(20).
		All(&todos)
	if todos == nil {
		todos = []Todo{}
	}
	return todos, err
}

func FindTodoById(id string) (*Todo, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("Id is not a valid format")
	}
	todo := Todo{}
	err := app.DB.C("todos").
		Find(bson.M{"_id": bson.ObjectIdHex(id)}).
		One(&todo)
	return &todo, err
}

func (todo *Todo) Create() error {
	todo.Id = bson.NewObjectId()
	err := app.DB.C("todos").Insert(&todo)
	return err
}

func (todo *Todo) Update() error {
	err := app.DB.C("todos").Update(bson.M{"_id": todo.Id}, &todo)
	return err
}

func (todo *Todo) Delete() error {
	err := app.DB.C("todos").Remove(bson.M{"_id": todo.Id})
	return err
}
