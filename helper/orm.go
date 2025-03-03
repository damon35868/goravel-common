package helper

import (
	"github.com/damon35868/goravel-common/common"
	"github.com/damon35868/goravel-common/dto"
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
)

type Pageable interface {
	GetPageInfo() (int, int)
}

func PageBuilder[T any, TDTO Pageable](pageDto TDTO, queryBuilderActions ...func(qb orm.Query) orm.Query) dto.PageRespDto[T] {
	var dataVal dto.PageRespDto[T]
	queryBuilder := facades.Orm().Query()
	if len(queryBuilderActions) > 0 {
		queryBuilderAction := queryBuilderActions[0]
		queryBuilder = queryBuilderAction(queryBuilder)
	}

	page, pageSize := pageDto.GetPageInfo()
	queryBuilder.Paginate(page, pageSize, &dataVal.Items, &dataVal.TotalCount)
	dataVal.HasNextPage = common.HasNextPage(int64(page), int64(pageSize), dataVal.TotalCount)

	return dataVal
}
