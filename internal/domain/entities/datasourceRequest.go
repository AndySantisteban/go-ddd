package entities

type DataSourceRequest struct {
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
	Sort     string `form:"sort"`
	Filter   struct {
		Field    string `form:"field"`
		Value    string `form:"value"`
		Operator string `form:"operator"`
	} `form:"filter"`
}

type DataSourceResponse[T any] struct {
	Data  []T
	Total int
}
