package models

// TODO clean up for consistency

import (
	"../app"
	"../utils"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Todo struct {
	Id       bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Title    string          `json:"title" bson:"title" binding:"required"`
	Parent   bson.ObjectId   `json:"parent" bson:"parent,omitempty"`
	Children []bson.ObjectId `json:"children" bson:"children,omitempty"`
}

type TodoMove struct {
	// Parent       bson.ObjectId `json:"parent"`
	PriorSiblingId string `json:"prior_sibling_id"`
}

func findRootTodo() (*Todo, error) {
	root := Todo{}
	err := app.DB.C("todos").
		Find(bson.M{"parent": nil}).
		Select(bson.M{"children": 1}).
		One(&root)

	if err == mgo.ErrNotFound {
		root.Id = bson.NewObjectId()
		root.Title = "root"
		err = app.DB.C("todos").Insert(&root)
		if err != nil {
			return nil, err
		}
	}

	return &root, err
}

func FindTodos() ([]Todo, error) {

	root, err := findRootTodo()
	if err != nil {
		return nil, err
	}

	var todos []Todo
	err = app.DB.C("todos").
		Find(bson.M{"parent": root.Id}).
		All(&todos)
	if todos == nil {
		todos = []Todo{}
	}

	idToTodo := make(map[bson.ObjectId]Todo)
	for _, todo := range todos {
		idToTodo[todo.Id] = todo
	}

	N := len(root.Children)
	converted := make([]Todo, N)
	ptr := 0
	for i := 0; i < N; i++ {
		if todo, ok := idToTodo[root.Children[i]]; ok {
			converted[ptr] = todo
			ptr++
			delete(idToTodo, todo.Id)
		} else {
			log.Println("A missing sorted reference: " + root.Children[i].Hex())

			err = app.DB.C("todos").Update(bson.M{"_id": root.Id},
				bson.M{"$pull": bson.M{"children": root.Children[i]}})
			if err != nil {
				// Ignore the error for availability
				log.Println("Error in todos sort filed pull: " + err.Error())
			}
		}
	}
	converted = converted[:ptr]

	for _, todo := range idToTodo {
		log.Println("A missing master todo: " + todo.Id.Hex())
		converted = append(converted, todo)
	}

	return converted, err
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
	root, err := findRootTodo()
	if err != nil {
		return err
	}

	todo.Id = bson.NewObjectId()
	todo.Parent = root.Id
	err = app.DB.C("todos").Insert(&todo)
	if err != nil {
		return err
	}

	err = app.DB.C("todos").Update(bson.M{"_id": root.Id},
		bson.M{"$addToSet": bson.M{"children": todo.Id}})
	if err != nil {
		// Ignore the error for availability
		log.Println("Error in addToSet for sorted references: " + err.Error())
	}

	return nil
}

func (todo *Todo) Update() error {
	err := app.DB.C("todos").Update(bson.M{"_id": todo.Id}, &todo)
	return err
}

func (todo *Todo) Delete() error {
	root, err := findRootTodo()
	if err != nil {
		return err
	}

	err = app.DB.C("todos").Remove(bson.M{"_id": todo.Id})
	if err != nil {
		return err
	}

	err = app.DB.C("todos").Update(bson.M{"_id": root.Id},
		bson.M{"$pull": bson.M{"children": todo.Id}})
	if err != nil {
		// Ignore the error for availability
		log.Println("Error in pull for sorted references: " + err.Error())
	}

	return nil
}

func (todo *Todo) Move(todoMove *TodoMove) error {
	root, err := findRootTodo()
	if err != nil {
		return err
	}

	children, err := utils.MoveInChildren(
		root.Children, todo.Id, todoMove.PriorSiblingId)
	if err != nil {
		return err
	}
	root.Children = children

	err = app.DB.C("todos").Update(bson.M{"_id": root.Id}, &root)
	return err
}
