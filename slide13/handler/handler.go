package handler

import (
	"fmt"
	"github.com/vadym-98/playground/slide13/middleware"
	"net/http"
	"regexp"
)

type entry struct {
	method  string
	pattern *regexp.Regexp
	handler http.Handler
	m       middleware.Func
}

type RegexMuxer struct {
	entries []entry
}

func NewRegexMuxer() *RegexMuxer {
	return new(RegexMuxer)
}

func (h *RegexMuxer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, e := range h.entries {
		if e.pattern.MatchString(r.URL.Path) && e.method == r.Method {
			if e.m != nil {
				e.m(e.handler).ServeHTTP(w, r)
			} else {
				e.handler.ServeHTTP(w, r)
			}

			return
		}
	}

	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

func (h *RegexMuxer) addEntry(method, pattern string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	h.entries = append(h.entries, entry{
		method:  method,
		pattern: regexp.MustCompile(fmt.Sprintf("^%s$", pattern)),
		handler: http.HandlerFunc(f),
		m:       m,
	})
}

func (h *RegexMuxer) Get(pattern string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	h.addEntry(http.MethodGet, pattern, f, m)
}

func (h *RegexMuxer) Post(pattern string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	h.addEntry(http.MethodPost, pattern, f, m)
}

func (h *RegexMuxer) Put(pattern string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	h.addEntry(http.MethodPut, pattern, f, m)
}

func (h *RegexMuxer) Delete(pattern string, f func(w http.ResponseWriter, r *http.Request), m middleware.Func) {
	h.addEntry(http.MethodDelete, pattern, f, m)
}
