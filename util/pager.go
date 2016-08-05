package util

type Pager struct {
	PageNo    int64
	PageSize  int64
	PageCount int64
	Total     int64
	PageData  []interface{}
}

func NewPager(pageNo, pageSize, total int64, pageData []interface{}) *Pager {
	// pageCount := total % pageSize == 0 ? total % pageSize :
	pageCount := 0
	if total%pageSize == 0 {
		pageCount = toatal / pageSize
	} else {
		pageCount = total/pageSize + 1
	}

	return &Pager{pageNo, pageSize, pageCount, total, pageData}
}
