package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yousefggg/api-gateway/internal/proxy"
)

type Router struct {
	authProxy *proxy.AuthProxy
	chatProxy *proxy.ChatProxy
	tourProxy *proxy.TourProxy
}

func NewRouter(
	auth *proxy.AuthProxy,
	chat *proxy.ChatProxy,
	tour *proxy.TourProxy,
) *Router {
	return &Router{
		authProxy: auth,
		chatProxy: chat,
		tourProxy: tour,
	}
}

func (r *Router) Init() http.Handler {
	muxRouter := mux.NewRouter()

	muxRouter.PathPrefix("/api/auth/").HandlerFunc(r.authProxy.ServeHTTP)

	muxRouter.PathPrefix("/api/chat/").HandlerFunc(r.chatProxy.ServeHTTP)

	muxRouter.PathPrefix("/api/tours/").HandlerFunc(r.tourProxy.ServeHTTP)

	muxRouter.PathPrefix("/ws").HandlerFunc(r.chatProxy.ServeWebSocket)

	return muxRouter
}