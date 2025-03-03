package dto

// req
type (
	PageReqDto struct {
		Page     int64 `json:"page"`
		PageSize int64 `json:"pageSize"`
	}

	PageWhereReqDto[T any] struct {
		Page     int64 `json:"page"`
		PageSize int64 `json:"pageSize"`
		Where    *T    `json:"where,omitempty"`
	}
)

// res
type (
	PageRespDto[T any] struct {
		Items       []T   `json:"items"`
		TotalCount  int64 `json:"totalCount"`
		HasNextPage bool  `json:"hasNextPage"`
	}
)
