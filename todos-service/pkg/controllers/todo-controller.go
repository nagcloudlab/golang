package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nagcloudlab/todos-service/pkg/model"
	"net/http"
	"strconv"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := model.FindAll()
	res, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}
func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	todoId, err1 := strconv.ParseInt(id, 10, 0)
	if err1 != nil {
		fmt.Println("error while parsing")
	}
	todo := model.FindById(todoId)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err2 := w.Write(res)
	if err2 != nil {
		return
	}
}
