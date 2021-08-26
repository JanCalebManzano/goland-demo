package datastructs

import (
	"time"
)

// User is the data structure for users
type User struct {
	ID           int32      `json:"id" bson:"_id"`
	Username     string     `json:"username,omitempty" bson:"username"`
	FullName     string     `json:"full_name,omitempty" bson:"full_name"`
	MobileNumber string     `json:"mobile_number,omitempty" bson:"mobile_number"`
	EmailAddress string     `json:"email_address,omitempty" bson:"email_address"`
	CreatedAt    *time.Time `json:"created_at,omitempty" bson:"created_at"`
	IsActive     bool       `json:"is_active,omitempty" bson:"is_active"`
}
