package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type AuthProxy struct {
	proxy *httputil.ReverseProxy
}

func NewAuthProxy(target string) (*AuthProxy, error) {

	url, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	return &AuthProxy{
		proxy: httputil.NewSingleHostReverseProxy(url),
	}, nil
}

func (p *AuthProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/auth")

	p.proxy.ServeHTTP(w, r)
}