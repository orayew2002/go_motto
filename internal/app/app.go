package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/orayew2002/go_motto/internal/domains"
	"github.com/orayew2002/go_motto/internal/handler"
	"github.com/orayew2002/go_motto/internal/repository"
	"github.com/orayew2002/go_motto/internal/service"
)

func Run(appDependencies domains.AppDependencies) {
	repo := repository.NewRepository()
	service := service.NewService(repo, appDependencies.Logger)
	handler := handlers.NewHandler(service)

	appRoutes := handlers.Routes(handler)

	addr := "127.0.0.1:8080"
	srv := &http.Server{
		Addr:         addr,
		Handler:      appRoutes,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Printf("🚀 Server running at http://%s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("❌ Server error: %v", err)
	}
}
