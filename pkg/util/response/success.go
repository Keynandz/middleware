package response

import (
	"net/http"

	"go-base-structure/model/base"

	"github.com/labstack/echo/v4"
)

type successConstant struct {
	OK Success
}

type successConstantWithTotal struct {
	OK SuccessWithTotal
}

var SuccessConstant successConstant = successConstant{
	OK: Success{
		Response: successResponse{
			Meta: Meta{
				Success: true,
				Message: "Request successfully proceed",
			},
			Data: nil,
		},
		Code: http.StatusOK,
	},
}

var SuccessConstantWithTotal successConstantWithTotal = successConstantWithTotal{
	OK: SuccessWithTotal{
		Response: successResponseWithTotal{
			Meta: Meta{
				Success: true,
				Message: "Request successfully proceed",
			},
			Data: nil,
			Total: 0,
		},
		Code: http.StatusOK,
	},
}

type successResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type successResponseWithTotal struct {
	Meta  Meta        `json:"meta"`
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

type SuccessWithTotal struct {
	Response successResponseWithTotal `json:"response"`
	Code     int                      `json:"code"`
}

type Success struct {
	Response successResponse `json:"response"`
	Code     int             `json:"code"`
}

func SuccessBuilder(res *Success, data interface{}) *Success {
	res.Response.Data = data
	return res
}

func SuccessBuilderWithTotal(res *SuccessWithTotal, data interface{}, total int) *SuccessWithTotal {
	res.Response.Data = data
	res.Response.Total = total
	return res
}

func CustomSuccessBuilder(code int, data interface{}, message string, info *base.PaginationInfo) *Success {
	return &Success{
		Response: successResponse{
			Meta: Meta{
				Success: true,
				Message: message,
				Info:    info,
			},
			Data: data,
		},
		Code: code,
	}
}

func SuccessResponse(data interface{}) *Success {
	return SuccessBuilder(&SuccessConstant.OK, data)
}

func SuccessResponseWithTotal(data interface{}, total int) *SuccessWithTotal {
	return SuccessBuilderWithTotal(&SuccessConstantWithTotal.OK, data, total)
}

func (s *Success) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}

func (s *SuccessWithTotal) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}
