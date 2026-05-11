package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ChatProxy struct {
	proxy *httputil.ReverseProxy
	wsURL string
}

func NewChatProxy(target string) (*ChatProxy, error) {

	url, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	return &ChatProxy{
		proxy: httputil.NewSingleHostReverseProxy(url),
		wsURL: target,
	}, nil
}

func (p *ChatProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/chat")
	p.proxy.ServeHTTP(w, r)
}
func (p *ChatProxy) ServeWebSocket(w http.ResponseWriter, r *http.Request) {
	p.proxy.ServeHTTP(w, r)
}