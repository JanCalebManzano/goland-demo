package datastructs

import (
	"time"
)

// User is the data structure for users
type User struct {
	ID           string     `json:"id,omitempty"`
	Username     string     `json:"username,omitempty"`
	FullName     string     `json:"full_name,omitempty"`
	MobileNumber string     `json:"mobile_number,omitempty"`
	EmailAddress string     `json:"email_address,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	IsActive     bool       `json:"is_active,omitempty"`
}
