package helper

import (
	"github.com/damon35868/goravel-common/common"
	"github.com/damon35868/goravel-common/dto"
	"github.com/damon35868/goravel-common/requests"
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

func PageBuilder[T any](pageDto requests.PageRequest, queryBuilderActions ...func(qb orm.Query) orm.Query) dto.PageRespDto[T] {
	var dataVal dto.PageRespDto[T]
	queryBuilder := facades.Orm().Query()
	if len(queryBuilderActions) > 0 {
		queryBuilderAction := queryBuilderActions[0]
		queryBuilder = queryBuilderAction(queryBuilder)
	}

	queryBuilder.Paginate(pageDto.Page, pageDto.PageSize, &dataVal.Items, &dataVal.TotalCount)
	dataVal.HasNextPage = common.HasNextPage(int64(pageDto.Page), int64(pageDto.PageSize), dataVal.TotalCount)

	return dataVal
}
