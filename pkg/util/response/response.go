package response

import 	"go-base-structure/model/base"


type Meta struct {
	Success bool                 `json:"success" default:"true"`
	Message string               `json:"message" default:"true"`
	Info    *base.PaginationInfo `json:"info"`
}
