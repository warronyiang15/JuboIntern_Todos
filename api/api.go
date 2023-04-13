package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"todos/db"
	"todos/utility"

	"github.com/gorilla/mux"
)

func GetAllTodos(mydb *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ret, todoArray := db.GetAllTodosFromDB(mydb)
		if ret == 0 {
			// something wrong
			utility.RespondWithJson(w, r, http.StatusInternalServerError, []byte("{}"))
			return
		}

		// Response with json
		utility.RespondWithJson(w, r, http.StatusOK, todoArray)
	}
}

func GetTodosWithID(mydb *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			// Something went wrong when decoding params
			utility.RespondWithJson(w, r, http.StatusBadRequest, []byte("{}"))
			return
		}

		ret, todos := db.GetTodosWithIDFromDB(mydb, id)
		if ret == 0 {
			// something went wrong
			utility.RespondWithJson(w, r, http.StatusNotFound, []byte("{}"))
			return
		}

		// Response with JSON
		utility.RespondWithJson(w, r, http.StatusOK, todos)
	}
}

func CreateTodos(mydb *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ret, newtodos := utility.DecodeJsonQuery(w, r)
		if ret == 0 {
			//Something went wrong when decoding json query
			utility.RespondWithJson(w, r, http.StatusBadRequest, []byte("{}"))
			return
		}
		ID := int(db.CreateTodosToDB(mydb, newtodos))
		if ID == 0 {
			// Something went wrong when insert into Database
			utility.RespondWithJson(w, r, http.StatusInternalServerError, []byte("{}"))
			return
		}

		ret, newtodos = db.GetTodosWithIDFromDB(mydb, ID)
		if ret == 0 {
			// Something went wrong when refetching
			utility.RespondWithJson(w, r, http.StatusInternalServerError, []byte("{}"))
			return
		}

		// Response with JSON
		utility.RespondWithJson(w, r, http.StatusOK, newtodos)
	}
}

func UpdateTodos(mydb *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			// Something went wrong when decoding request id
			utility.RespondWithJson(w, r, http.StatusBadRequest, []byte("{}"))
			return
		}

		ret, newtodos := utility.DecodeJsonQuery(w, r)
		if ret == 0 {
			// Something went wrong when decoding json query
			utility.RespondWithJson(w, r, http.StatusBadRequest, []byte("{}"))
			return
		}

		// Set the ID as PUT id
		newtodos.Id = id

		retID := db.UpdateTodosToDB(mydb, newtodos)
		if retID == 0 {
			// Something went wrong
			utility.RespondWithJson(w, r, http.StatusInternalServerError, []byte("{}"))
			return
		}

		ret, newtodos = db.GetTodosWithIDFromDB(mydb, retID)
		if ret == 0 {
			// Something went wrong when refetching
			utility.RespondWithJson(w, r, http.StatusInternalServerError, []byte("{}"))
			return
		}

		utility.RespondWithJson(w, r, http.StatusOK, newtodos)
	}
}

func DeleteTodos(mydb *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			// Something went wrong when decoding request id
			utility.RespondWithJson(w, r, http.StatusBadRequest, []byte("{}"))
			return
		}

		ret := db.DeleteTodosFromDB(mydb, id)

		if ret == 0 {
			// Something went wrong when delete todos from database
			utility.RespondWithJson(w, r, http.StatusInternalServerError, []byte("{}"))
			return
		}

		// Success!
		utility.RespondWithJson(w, r, http.StatusOK, []byte("{}"))
	}
}
