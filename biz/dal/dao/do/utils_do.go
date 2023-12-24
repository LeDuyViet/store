package do

import "fmt"

const (
	BASE_URL_STORE = "store-data.volio.vn"
)

type MetaDo struct {
	CurrentPage int    `json:"current_page"`
	PerPage     int    `json:"per_page"`
	From        int    `json:"from"`
	To          int    `json:"to"`
	Total       int    `json:"total"`
	Path        string `json:"path"`
	LastPage    int    `json:"last_page"`
}

type LinksDo struct {
	First string  `json:"first"`
	Last  string  `json:"last"`
	Prev  *string `json:"prev"`
	Next  *string `json:"next"`
}

func CreateMetaDo(currentPage, perPage, total int, path string) *MetaDo {
	if perPage == -1 {
		return nil
	}

	from := (currentPage-1)*perPage + 1
	LastPage := 1

	if perPage != 0 {
		LastPage = total / perPage
		if total%perPage != 0 {
			LastPage = total/perPage + 1
		}
	}

	to := from + perPage - 1
	if to > total {
		to = total
	}

	return &MetaDo{
		CurrentPage: currentPage,
		Path:        path,
		PerPage:     perPage,
		From:        from,
		To:          to,
		Total:       total,
		LastPage:    LastPage,
	}
}

// example
/**
	currentPage := 2
	perPage:=2
	total:=20
	path:="http://store.volio.vn/api/v2/public/modules"
	query:="app_id=2"

links := p.CreateLinksDo(currentPage, perPage, total, path, query)
*/
func CreateLinksDo(currentPage, perPage, total int, path, query string) *LinksDo {
	var (
		nextPage *string
		prevPage *string
	)

	if perPage == -1 {
		return nil
	}

	baseurl := func(page int) string {
		if perPage == 0 {
			return fmt.Sprintf("%s?%s&page=%v", path, query, page)
		}
		return fmt.Sprintf("%s?%s&per_page=%v&page=%v", path, query, perPage, page)
	}

	LastPage := 1
	if perPage != 0 {
		LastPage = total / perPage
		if total%perPage != 0 {
			LastPage = total/perPage + 1
		}
	}

	if currentPage+1 <= LastPage {
		next := baseurl(currentPage + 1)
		nextPage = &next
	}

	if currentPage > 1 {
		prev := baseurl(currentPage - 1)
		prevPage = &prev
	}

	return &LinksDo{
		First: baseurl(1),
		Last:  baseurl(LastPage),
		Next:  nextPage,
		Prev:  prevPage,
	}
}

type ResPaginate struct {
	Data  interface{} `json:"data,omitempty"`
	Links *LinksDo    `json:"links,omitempty"`
	Meta  *MetaDo     `json:"meta,omitempty"`
}
