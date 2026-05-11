package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type TourProxy struct {
	proxy *httputil.ReverseProxy
}

func NewTourProxy(target string) (*TourProxy, error) {

	url, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	return &TourProxy{
		proxy: httputil.NewSingleHostReverseProxy(url),
	}, nil
}

func (p *TourProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/tours")
	p.proxy.ServeHTTP(w, r)
}