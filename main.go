package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/DaniDMoura/todo-api-stdlibrary-golang/domain"
)

var todos []domain.Todo = []domain.Todo{}

func loadTodos() {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &todos)
}

func saveTodos() {
	data, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}

	os.WriteFile("tasks.json", data, 0644)
}

func main() {
	loadTodos()
	logger := log.Default()

	http.HandleFunc("/todos", HandleListUsers)
	http.HandleFunc("/todos/create", HandleCreateUsers)
	http.HandleFunc("/todo", HandleGetUser)
	http.HandleFunc("/todos/update", HandleUpdateTodo) 
	http.HandleFunc("/todos/delete", HandleDeleteTodo) 

	logger.Println("API rodando em :8080")
	http.ListenAndServe(":8000", nil)
}

func HandleListUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		json.NewEncoder(w).Encode(domain.CustomError{
			Message:    "Método não permitido",
			StatusCode: http.StatusMethodNotAllowed},
		)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


	json.NewEncoder(w).Encode(todos)
}

func HandleCreateUsers(w http.ResponseWriter, r *http.Request) {
	defer saveTodos()

	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode(domain.CustomError{
			Message:    "Método não permitido",
			StatusCode: http.StatusMethodNotAllowed},
		)
		return
	}

	var todo domain.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	w.WriteHeader(http.StatusCreated)
	todos = append(todos, domain.Todo{ID: len(todos) + 1, Name: todo.Name})

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		json.NewEncoder(w).Encode(domain.CustomError{
			Message:    "Método não permitido",
			StatusCode: http.StatusMethodNotAllowed},
		)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	var todo *domain.Todo
	for i := range todos {
		if todos[i].ID == id {
			todo = &todos[i]
			break
		}
	}

	if todo == nil {
		json.NewEncoder(w).Encode(domain.CustomError{
			Message:    "Id não existe",
			StatusCode: http.StatusNotFound},
		)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	defer saveTodos()

	if r.Method != http.MethodPut {
		json.NewEncoder(w).Encode(domain.CustomError{
			Message:    "Método não permitido",
			StatusCode: http.StatusMethodNotAllowed},
		)
		return
	}

	var new_todo *domain.Todo 
	json.NewDecoder(r.Body).Decode(&new_todo)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		json.NewEncoder(w).Encode(domain.CustomError{
			Message:    "Id invalido",
			StatusCode: http.StatusNotFound},
		)
		return
	}

	found := false
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Name = new_todo.Name
			todos[i].Done = new_todo.Done
			found = true
			break
		}
	}

	if !found {
		json.NewEncoder(w).Encode(domain.CustomError{
			Message:    "Id não existe",
			StatusCode: http.StatusMethodNotAllowed},
		)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func HandleDeleteTodo(w http.ResponseWriter, r *http.Request){
	defer saveTodos()

	if r.Method != http.MethodDelete {
		json.NewEncoder(w).Encode(domain.CustomError{
			Message:    "Método não permitido",
			StatusCode: http.StatusMethodNotAllowed},
		)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		json.NewEncoder(w).Encode(domain.CustomError{
			Message:    "Id invalido",
			StatusCode: http.StatusBadRequest},
		)
		return
	}

	found := false
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		json.NewEncoder(w).Encode(domain.CustomError{
			Message:    "Id não existe",
			StatusCode: http.StatusNotFound},
		)
		return
	}

	json.NewEncoder(w).Encode(todos)
}
