package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi"
	"github.com/howters/gopack/pkg/config"
	"github.com/howters/gopack/pkg/handler"
)

func routes(app *config.AppConfig) http.Handler{
	// mux := pat.New();
	// mux.Get("/", http.HandlerFunc(handler.Repo.Home));
	// mux.Get("/about", http.HandlerFunc(handler.Repo.About));

	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoader)

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)

	return mux
}
