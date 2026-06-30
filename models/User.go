package models

type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
}
