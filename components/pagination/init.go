package pagination

import (
	"github.com/astaxie/beego/context"
	"strconv"
)

type Pagination struct {
	Page  int
	Size  int
	Limit int
}

func initPagination() Pagination {
	var pagination Pagination
	pagination.Page = 1
	pagination.Size = 20
	pagination.setLimit()

	return pagination
}

func InitPagination(ctx *context.Context) Pagination {
	pagination := initPagination()

	page, _ := strconv.Atoi(ctx.Input.Query("page"))
	size, _ := strconv.Atoi(ctx.Input.Query("page_size"))
	if page > 0 {
		pagination.Page = page
	}

	if size > 0 {
		pagination.Size = size
	}

	pagination.setLimit()

	return pagination
}

func (p *Pagination) setLimit() {
	p.Limit = (p.Page - 1) * p.Size
}
