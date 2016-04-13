package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
	"todo_center/go_todo_api/utils"
)

type Todo struct {
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title   string        `json:"title" bson:"title" binding:"required"`
	DueDate int64         `json:"due_date" bson:"due_date,omitempty"`
}

type TodoGroup struct {
	Id    bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Title string          `json:"title" bson:"title" binding:"required"`
	Todos []bson.ObjectId `json:"todos" bson:"todos,omitempty"`
}

type TodoMove struct {
	PriorSiblingId string `json:"prior_sibling_id"`
}

func findTodoGroup(db *mgo.Database) (*TodoGroup, error) {
	// TODO CRUD for TodoGroup
	group := TodoGroup{}
	err := db.C("todo_groups").
		Find(bson.M{}).
		Select(bson.M{"todos": 1}).
		One(&group)

	if err == mgo.ErrNotFound {
		group.Id = bson.NewObjectId()
		group.Title = "root"
		err = db.C("todo_groups").Insert(&group)
		if err != nil {
			return nil, err
		}
	}

	return &group, err
}

func FindTodos(db *mgo.Database) ([]Todo, error) {

	group, err := findTodoGroup(db)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	err = db.C("todos").
		Find(bson.M{"_id": bson.M{"$in": group.Todos}}).
		All(&todos)
	if err != nil {
		return nil, err
	}
	if todos == nil {
		todos = []Todo{}
	}

	N := len(group.Todos)
	sortedTodos := make([]Todo, N)
	idToTodo := make(map[bson.ObjectId]Todo)
	for _, todo := range todos {
		idToTodo[todo.Id] = todo
	}
	ptr := 0
	for i := 0; i < N; i++ {
		if todo, ok := idToTodo[group.Todos[i]]; ok {
			sortedTodos[ptr] = todo
			ptr++
		}
	}

	return sortedTodos[:ptr], err
}

func FindTodoById(db *mgo.Database, id string) (*Todo, error) {
	if err := utils.ValidateObjectIdHex(id); err != nil {
		return nil, err
	}

	todo := Todo{}
	err := db.C("todos").
		Find(bson.M{"_id": bson.ObjectIdHex(id)}).
		One(&todo)
	return &todo, err
}

func (todo *Todo) Create(db *mgo.Database) error {
	group, err := findTodoGroup(db)
	if err != nil {
		return err
	}

	todo.Id = bson.NewObjectId()
	err = db.C("todos").Insert(&todo)
	if err != nil {
		return err
	}

	err = db.C("todo_groups").Update(bson.M{"_id": group.Id},
		bson.M{"$addToSet": bson.M{"todos": todo.Id}})
	if err != nil {
		return err
	}

	return nil
}

func (todo *Todo) Update(db *mgo.Database) error {
	err := db.C("todos").Update(bson.M{"_id": todo.Id}, &todo)
	return err
}

func (todo *Todo) Delete(db *mgo.Database) error {
	group, err := findTodoGroup(db)
	if err != nil {
		return err
	}

	err = db.C("todo_groups").Update(bson.M{"_id": group.Id},
		bson.M{"$pull": bson.M{"todos": todo.Id}})
	if err != nil {
		return err
	}

	err = db.C("todos").Remove(bson.M{"_id": todo.Id})
	if err != nil {
		log.Println("Error in removing related todos: " + err.Error())
		// Ignore err
	}

	return nil
}

func (todo *Todo) Move(db *mgo.Database, todoMove *TodoMove) error {
	group, err := findTodoGroup(db)
	if err != nil {
		return err
	}

	movedTodos, err := utils.MoveInChildren(
		group.Todos, todo.Id, todoMove.PriorSiblingId)
	if err != nil {
		return err
	}
	group.Todos = movedTodos

	err = db.C("todo_groups").Update(bson.M{"_id": group.Id}, &group)
	return err
}
