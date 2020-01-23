package delivery

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/maestre3d/go-mq-example/common"
)

// Router HTTP Router
type Router struct {
	Mux *chi.Mux
}

// NewRouter Get an HTTP Router instance
func NewRouter() *Router {
	router := chi.NewRouter()

	routerInstance := new(Router)
	routerInstance.Mux = router

	routerInstance.setCORSPolicy()
	routerInstance.setMiddleware()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, common.Response{Message: "Hello from HTTP REST"})
	})

	return routerInstance
}

// setCORSPolicy() Set CORS Policy to the router instance
func (r *Router) setCORSPolicy() {
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Mux.Use(cors.Handler)
}

// setMiddleware Set required middleware to the router instance
func (r *Router) setMiddleware() {
	// A good base middleware stack
	r.Mux.Use(middleware.RequestID)
	r.Mux.Use(middleware.RealIP)
	r.Mux.Use(middleware.Logger)
	r.Mux.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Mux.Use(middleware.Timeout(60 * time.Second))
}
