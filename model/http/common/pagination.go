package common

import (
	punqy "github.com/punqy/core"
	logger "github.com/sirupsen/logrus"
	"math"
	"reflect"
	"strconv"
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func PaginationFromReq(r punqy.Request, maxLimit int) Pagination {
	pag := Pagination{Page: 1, Limit: maxLimit}
	p, err := strconv.Atoi(r.Get("page", strconv.Itoa(pag.Page)))
	if err != nil {
		p = 1
		logger.Warnf("pagination invalid arg page %s", r.URI().QueryArgs().Peek("page"))
	}
	pag.Page = p
	l, err := strconv.Atoi(r.Get("limit", strconv.Itoa(maxLimit)))
	if err != nil {
		l = maxLimit
		logger.Warnf("pagination invalid arg page %s", r.URI().QueryArgs().Peek("limit"))
	}
	if maxLimit < l {
		l = maxLimit
	}
	pag.Limit = l
	return pag
}

func (p Pagination) ToStorage() punqy.Pagination {
	var offset int
	if p.Page > 1 {
		offset = p.Limit * (p.Page - 1)
	}
	return punqy.Pagination{
		Limit:  uint32(p.Limit),
		Offset: uint32(offset),
	}
}

type HtmlPagination struct {
	Items       interface{}
	Route       string
	Pages       []int
	PageRange   int
	Page        int
	Limit       int
	Last        int
	First       int
	PageCount   int
	TotalCount  int
	CurrenCount int
	StartPage   int
	EndPage     int
	Previous    int
	Next        int
}

func NewHtmlPagination(pager Pagination, items interface{}, totalCount int, route string) HtmlPagination {
	p := HtmlPagination{
		Page:       pager.Page,
		Limit:      pager.Limit,
		Route:      route,
		Items:      items,
		TotalCount: totalCount,
		PageRange:  5,
	}
	if reflect.TypeOf(p.Items).Kind() == reflect.Slice {
		p.CurrenCount = reflect.ValueOf(p.Items).Len()
	} else {
		logger.Error(punqy.NewLucierror(500, "Pagination requires slice of interface"))
		return p
	}
	pageCount := int(math.Ceil(float64(p.TotalCount) / float64(p.Limit)))
	current := p.Page
	if pageCount < current {
		p.Page = pageCount
	}
	if p.PageRange > pageCount {
		p.PageRange = pageCount
	}
	delta := int(math.Ceil(float64(p.PageRange) / 2))
	var pages []int
	if (current - delta) > (pageCount - p.PageRange) {
		pages = makeRange(pageCount-p.PageRange+1, pageCount)
	} else {
		if current-delta < 0 {
			delta = current
		}
		offset := current - delta
		pages = makeRange(offset+1, offset+p.PageRange)
	}
	proximity := int(math.Floor(float64(p.PageRange) / 2))
	startPage := current - proximity
	endPage := current + proximity
	if startPage < 1 {
		endPage = min(endPage+(1-startPage), pageCount)
	}
	if endPage > startPage {
		startPage = max(startPage-(endPage-pageCount), 1)
		endPage = pageCount
	}
	if current > 1 {
		p.Previous = current - 1
	}
	if current < pageCount {
		p.Next = current + 1
	}
	p.Page = current
	p.Pages = pages
	p.PageCount = pageCount
	return p
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
