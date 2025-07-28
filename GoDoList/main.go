package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"text"`
	Completed bool   `json:"completed"`
}

var (
	todos  []Todo     // 存储所有待办事项
	nextID = 1        // 用于生成唯一ID
	mu     sync.Mutex // 互斥锁，保证并发安全
)

func todosHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	switch r.Method {
	case http.MethodGet:
		// ... 获取所有todos的逻辑 ...
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	case http.MethodPost:
		// ... 创建新todo的逻辑 ...
		var input struct {
			Text string `json:"text"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		todo := Todo{
			ID:        nextID,
			Task:      input.Text,
			Completed: false,
		}
		nextID++
		todos = append(todos, todo)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todo)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// 从URL中解析ID，例如 strings.TrimPrefix(r.URL.Path, "/api/todos/")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// 查找对应的todo
	var todoIndex = -1
	for i, t := range todos {
		if t.ID == id {
			todoIndex = i
			break
		}
	}
	if todoIndex == -1 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodPut:
		// ... 更新todo的逻辑 ...
		var input struct {
			Completed bool `json:"completed"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		todos[todoIndex].Completed = input.Completed

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos[todoIndex])
	case http.MethodDelete:
		// ... 删除todo的逻辑 ...
		todos = append(todos[:todoIndex], todos[todoIndex+1:]...)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/todos", todosHandler)
	http.HandleFunc("/api/todos/", todoHandler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
