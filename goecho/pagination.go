package goecho

import (
	"math"
)

type Pagination struct {
	ItemCount   uint
	CurrentPage uint

	PageSize  uint
	PageCount uint
}

func NewPagination(currentpage, itemcount uint) *Pagination {
	return &Pagination{
		CurrentPage: currentpage,
		ItemCount:   itemcount,
	}
}

func (pagination *Pagination) Calc(pagesize uint) *Pagination {
	pagination.PageSize = pagesize
	pagination.PageCount = pagination.getPageCount()
	if pagination.CurrentPage >= pagination.PageCount {
		pagination.CurrentPage = pagination.PageCount - 1
	}
	return pagination
}

func (pagination *Pagination) getPageCount() uint {
	return uint(math.Ceil(float64(pagination.ItemCount+pagination.PageSize-1) / float64(pagination.PageSize)))
}

func (pagination *Pagination) getOffset() uint {
	return pagination.CurrentPage * pagination.PageSize
}

func (pagination *Pagination) ApplyLimit() (limit, offset uint) {
	limit = pagination.PageSize
	offset = pagination.getOffset()
	return
}
