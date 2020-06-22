package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

type Task struct {
	Name string
	Done bool
}

var tasks []Task = []Task{}

func main() {
	http.HandleFunc("/tasks", GetTaskHandler)
	http.HandleFunc("/tasks/html", TaskHtmlHandler)
	http.ListenAndServe(":8000", nil)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not supported", http.StatusMethodNotAllowed)
		return
	}

	RandomNumber := rand.Intn(100)
	Name := "Tarefa " + strconv.Itoa(RandomNumber)
	Done := false
	if RandomNumber%2 == 0 {
		Done = true
	}
	tasks = append(tasks, Task{Name, Done})

	json.NewEncoder(w).Encode(tasks)
}

func TaskHtmlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not supported", http.StatusMethodNotAllowed)
		return
	}

	tmpl := template.Must(template.ParseFiles("tasks.html"))

	tmpl.Execute(w, tasks)
}
