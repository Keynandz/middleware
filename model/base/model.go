package base

import (
	"time"
)

type Entity struct {
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy int        `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *int       `json:"updated_by"`
}
