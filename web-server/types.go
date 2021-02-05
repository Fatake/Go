package main

import "encoding/json"

// User User
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// IsValid IsValid
func (u *User) IsValid() bool {
	if u.Email == "" {
		return false
	}
	if u.Name == "" {
		return false
	}
	if u.Phone == "" {
		return false
	}
	return true
}

// ToJson ToJson
func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
