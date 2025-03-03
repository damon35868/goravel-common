package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type PageRequest struct {
	Page     int `form:"page" json:"page" validate:"required"`
	PageSize int `form:"pageSize" json:"pageSize" validate:"omitempty"`
}

func (r *PageRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *PageRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PageRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"page":     "required",
		"pageSize": "required",
	}
}

func (r *PageRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PageRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PageRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
