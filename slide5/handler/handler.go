package handler

import (
	"fmt"
	"net/http"
	"regexp"
)

type entry struct {
	method  string
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexMuxer struct {
	entries []entry
}

func NewRegexMuxer() *RegexMuxer {
	return &RegexMuxer{}
}

func (h *RegexMuxer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, e := range h.entries {
		if e.pattern.MatchString(r.URL.Path) && e.method == r.Method {
			e.handler.ServeHTTP(w, r)
			return
		}
	}

	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

func (h *RegexMuxer) addEntry(method, pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	h.entries = append(h.entries, entry{
		method:  method,
		pattern: regexp.MustCompile(fmt.Sprintf("^%s$", pattern)),
		handler: http.HandlerFunc(f),
	})
}

func (h *RegexMuxer) Get(pattern string, f func(w http.ResponseWriter, r *http.Request)) {
	h.addEntry(http.MethodGet, pattern, f)
}
