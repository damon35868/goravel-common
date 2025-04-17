package requests

import (
	"github.com/damon35868/goravel-common/utils"
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
		"page":     "required|integer",
		"pageSize": "required|integer",
	}
}

func (r *PageRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PageRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *PageRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	if err := utils.FormatRequest("page", data); err != nil {
		return err
	}
	if err := utils.FormatRequest("pageSize", data); err != nil {
		return err
	}
	return nil
}

func (r *PageRequest) GetPageInfo() (int, int) {
	return r.Page, r.PageSize
}
