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

func PageBuilder[T any](pageDto Pageable, queryBuilderActions ...func(qb orm.Query) orm.Query) (data dto.PageRespDto[T]) {
	queryBuilder := facades.Orm().Query()
	if len(queryBuilderActions) > 0 {
		queryBuilderAction := queryBuilderActions[0]
		queryBuilder = queryBuilderAction(queryBuilder)
	}

	page, pageSize := pageDto.GetPageInfo()
	queryBuilder.OrderByDesc("id").Paginate(page, pageSize, &data.Items, &data.TotalCount)
	data.HasNextPage = common.HasNextPage(int64(page), int64(pageSize), data.TotalCount)
	return
}
