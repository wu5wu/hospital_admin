package models

type Paging struct {
	Total int64
	Page  int `json:"pageNum"`
	Limit int `json:"pageSize"`
}
