package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sidharth461/crud-mysql-go/database"
	"github.com/Sidharth461/crud-mysql-go/models"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUsers(w, r)
	case http.MethodPost:
		PostUsers(w, r)
	case http.MethodPut:
		UpdateUsers(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}
func GetUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GetUsers called")
	rows, err := database.DB.Query("SELECT id, name, email, contact_number FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var u models.User

		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.ContactNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(u.Name) == "" {
		http.Error(w, "Name is Required", http.StatusBadRequest)
		return
	}
	result, err := database.DB.Exec(
		"INSERT INTO users (name, email, contact_number) VALUES (?, ?, ?)",
		u.Name,
		u.Email,
		u.ContactNumber,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"message": "User created successfully",
		"id":      id,
	}

	json.NewEncoder(w).Encode(response)

}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var u models.User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	result, err := database.DB.Exec(
		"UPDATE users SET name = ?, email = ?, contact_number = ? WHERE id = ?",
		u.Name,
		u.Email,
		u.ContactNumber,
		idInt,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"message": "User updated successfully",
		"id":      idInt,
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := database.DB.Exec(
		"DELETE FROM users WHERE id= ?",
		idInt,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "user Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"message": "User deleted successfully",
		"id":      idInt,
	}

	json.NewEncoder(w).Encode(response)
}
