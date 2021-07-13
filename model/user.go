package model

import "time"

type User struct {
	ID        string    `json:"-" db:"id, omitempty", `
	Email     string    `json:"email,omitempty" db:"email,omitempty"`
	Phone     string    `json:"phone,omitempty" db:"phone, omitempty"`
	Password  string    `json:"password,omitempty" db:"password,omitempty"`
	Address   string    `json:"address, omitempty" db:"address,omitempty"`
	FullName  string    `json:"full_name, omitempty" db:"full_name,omitempty"`
	Avatar    string    `json:"avatar,omitempty" db:"avatar,omitempty"`
	Role      string    `json:"-,omitempty" db:"role,omitempty"`
	Token     string    `json:"token,omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAT time.Time `json:"-" db:"updated_at,omitempty"`
}
