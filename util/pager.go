package util

type Pager struct {
	PageNo    int
	PageSize  int
	PageCount int64
	Total     int64
	PageData  []interface{}
}

func NewPager(pageNo, pageSize int, total int64, pageData []interface{}) *Pager {
	var pageCount int64
	if total%int64(pageSize) == 0 {
		pageCount = total / int64(pageSize)
	} else {
		pageCount = total/int64(pageSize) + 1
	}

	return &Pager{pageNo, pageSize, pageCount, total, pageData}
}
