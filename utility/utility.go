package utility

import (
	"encoding/json"
	"net/http"
)

type Todos struct {
	Id          int    `json:"id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Completed   bool   `json:"Completed"`
	CreatedAt   string `json:"CreatedAt"`
}

func RespondWithJson(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}

func DecodeJsonQuery(w http.ResponseWriter, r *http.Request) (ret int, retTodos Todos) {
	var newtodos Todos

	err := json.NewDecoder(r.Body).Decode(&newtodos)
	if err != nil {
		return 0, newtodos
	}
	return 1, newtodos
}

func InitializeTodos(id int, title string, description string, completed bool, createdat string) (todos Todos) {
	var newtodos Todos
	newtodos.Id = id
	newtodos.Title = title
	newtodos.Description = description
	newtodos.Completed = completed
	newtodos.CreatedAt = createdat
	return newtodos
}
