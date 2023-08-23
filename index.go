package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var db *sql.DB

func main() {
	// Connect to the database
	db = setupDB()
	defer db.Close()

	// Define API routes
	http.HandleFunc("/todos", handleTodos)
	http.HandleFunc("/todos/", handleTodo)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupDB() *sql.DB {
	server := "localhost"
	port := 63122
	user := "hamza"
	password := "6761795h"
	database := "todo"

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w)
	case http.MethodPost:
		createTodo(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		deleteTodo(w, r)
	} else {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter) {
	rows, err := db.Query("SELECT ID, title, status FROM todo_list")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO todo_list (ID, title, status) VALUES (?, ?, ?)",
		newTodo.ID, newTodo.Title, newTodo.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoIDStr := r.URL.Path[len("/todos/"):]
	todoID, err := strconv.Atoi(todoIDStr)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM todo_list WHERE ID = ?", todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
