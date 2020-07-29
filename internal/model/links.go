package model

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

//Links ...
type Links struct {
	Prev *Href `json:"prev,omitempty"`
	Next *Href `json:"next,omitempty"`
}

//Href ...
type Href struct {
	Href string `json:"href,omitempty"`
}

func (h *Href) setHref(u string, q url.Values) {
	h.Href = fmt.Sprintf("%v?%v", u, q.Encode())
}

//NewLinks ...
func NewLinks(r *http.Request, count, pageNum, pageSize int) *Links {

	links := &Links{}
	q := r.URL.Query()
	u := r.URL.Path
	var totalPage = count / pageSize
	if count%pageSize > 0 {
		totalPage++
	}
	prev, next := findPrevNext(totalPage, pageNum, pageSize)
	if prev != 0 {
		q.Set("page_num", strconv.Itoa(prev))
		prev := &Href{}
		prev.setHref(u, q)
		links.Prev = prev
	}
	if next != 0 {
		q.Set("page_num", strconv.Itoa(next))
		next := &Href{}
		next.setHref(u, q)
		links.Next = next
	}
	return links
}

func findPrevNext(totalPages, pageNum, pageSize int) (int, int) {
	prev := pageNum - 1
	if prev < 1 {
		prev = 0
	}
	next := pageNum + 1
	if next > totalPages {
		next = 0
	}
	return prev, next
}
