package model

import (
	"github.com/jinzhu/gorm"
	"github.com/nagcloudlab/todos-service/pkg/config"
)

type Todo struct {
	Id        int32
	Title     string
	Completed bool
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}

func Save(todo *Todo) *Todo {
	db.NewRecord(todo)
	db.Create(&todo)
	return todo
}

func FindAll() []Todo {
	var todos []Todo
	db.Find(&todos)
	return todos
}

func FindById(id int64) *Todo {
	var todo Todo
	db.Where("id=?", id).Find(&todo)
	return &todo
}

func DeleteTodo(id int64) Todo {
	var todo Todo
	db.Where("id=?", id).Delete(&todo)
	return todo
}
