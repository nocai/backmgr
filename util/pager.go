package util

type Pager struct {
	PageNo   int64
	PageSize int64
	Total    int64
	PageData *[]interface{}
}

func NewPager(pageNo, pageSize, total int64, pageData *[]interface{}) *Pager {
	return &Pager{pageNo, pageSize, total, pageData}
}
