package datastructs

import (
	"time"
)

// Todo is the data structure for todos
type Todo struct {
	ID          int32      `json:"id" bson:"_id"`
	UserID      int32      `json:"user_id" bson:"user_id"`
	Content     string     `json:"content,omitempty" bson:"content"`
	IsCompleted bool       `json:"is_completed,omitempty" bson:"is_completed"`
	CreatedAt   *time.Time `json:"created_at,omitempty" bson:"created_at"`
	IsActive    bool       `json:"is_active,omitempty" bson:"is_active"`
}
