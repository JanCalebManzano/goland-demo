package datastructs

import (
	"time"
)

// Todo is the data structure for todos
type Todo struct {
	ID          int64      `json:"id,omitempty"`
	Content     string     `json:"content,omitempty"`
	IsCompleted bool       `json:"is_completed,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	IsActive    bool       `json:"is_active,omitempty"`
}
