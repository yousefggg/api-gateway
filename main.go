package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/yousefggg/api-gateway/internal/config"
	"github.com/yousefggg/api-gateway/internal/proxy"
	"github.com/yousefggg/api-gateway/internal/router"
	"github.com/yousefggg/api-gateway/internal/transport"
)

func main() {

	_ = godotenv.Load()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	authProxy, err := proxy.NewAuthProxy(cfg.AuthServiceURL)
	if err != nil {
		log.Fatalf("auth proxy error: %v", err)
	}

	chatProxy, err := proxy.NewChatProxy(cfg.ChatServiceURL)
	if err != nil {
		log.Fatalf("chat proxy error: %v", err)
	}

	tourProxy, err := proxy.NewTourProxy(cfg.TourServiceURL)
	if err != nil {
		log.Fatalf("tour proxy error: %v", err)
	}

	r := router.NewRouter(authProxy, chatProxy, tourProxy)

	server := transport.NewServer(
		":"+cfg.Port,
		r,
		cfg.JWTSecret,
	)

	log.Println("API Gateway started on port", cfg.Port)

	if err := server.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}