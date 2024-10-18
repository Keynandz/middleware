package base

import (
	"time"
)

type Entity struct {
	Created   time.Time  `json:"created"`
	Createdby int        `json:"createdby"`
	Updated   *time.Time `json:"updated"`
	Updatedby *int       `json:"updatedby"`
}
