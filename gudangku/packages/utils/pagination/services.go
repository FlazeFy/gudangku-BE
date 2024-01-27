package pagination

import (
	"strconv"
)

func BuildPaginationResponse(page, pageSize, total, totalPages int, path string) PaginationResponse {
	var baseURL = "http://127.0.0.1:1234/"
	pagination := PaginationResponse{
		FirstPageURL: "",
		From:         0,
		LastPage:     totalPages,
		LastPageURL:  "",
		Links:        []Link{},
		NextPageURL:  "",
		Path:         "",
		PrevPageURL:  "",
		To:           0,
	}

	from := ((page - 1) * pageSize) + 1
	to := from + pageSize - 1
	if to > total {
		to = total
	}

	firstPageURL := baseURL + path + "?page=1"
	pagination.FirstPageURL = firstPageURL
	lastPageURL := baseURL + path + "?page=" + strconv.Itoa(totalPages)
	pagination.LastPageURL = lastPageURL

	var nextPageURL string
	if page < totalPages {
		nextPageURL = baseURL + path + "?page=" + strconv.Itoa(page+1)
	}
	pagination.NextPageURL = nextPageURL
	var prevPageURL string
	if page > 1 {
		prevPageURL = baseURL + path + "?page=" + strconv.Itoa(page-1)
	}
	pagination.PrevPageURL = prevPageURL

	links := []Link{}
	for i := 1; i <= totalPages; i++ {
		link := Link{
			URL:    baseURL + path + "?page=" + strconv.Itoa(i),
			Label:  strconv.Itoa(i),
			Active: i == page,
		}
		links = append(links, link)
	}
	pagination.Links = links

	pagination.From = from
	pagination.To = to
	pagination.Path = baseURL + path

	return pagination
}
