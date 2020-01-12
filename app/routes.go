package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/lucasstettner/api-boilerplate/app/features/status"
	"github.com/lucasstettner/api-boilerplate/app/utils/responses"
	"github.com/lucasstettner/api-boilerplate/config"
	_ "github.com/lucasstettner/api-boilerplate/docs"
	"github.com/lucasstettner/api-boilerplate/pkg/logger"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(c *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		logger.Logger(c.Logger),                       // Log API request calls
		middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
		middleware.Recoverer,                          // Recover from panics without crashing server
	)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		responses.NewResponse(w, http.StatusOK, nil, "Hello my new api!")
	})

	// Mount routes
	router.Mount("/status", status.Routes())

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition"
	))

	return router
}
