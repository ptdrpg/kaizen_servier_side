
package router

import (
	"net/http"
	"KageNoEn/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Router struct {
	R *chi.Mux
	C *controller.Controller
}

func NewRouter(c *controller.Controller) *Router {
	r := chi.NewRouter()

	// Middleware CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	return &Router{
		R: r,
		C: c,
	}
}

func (r *Router) RegisterRouter() {
	// r.R.Group(func(public chi.Router) {
	// 	public.Route("/api/v1/login", func(login chi.Router) {
	// 		login.Post("/", r.C.Login)
	// 	})
	// })

	r.R.Group(func(private chi.Router) {
		// private.Use(lib.JWTMiddleware)
		private.Route("/api/v1", func(v1 chi.Router) {
			v1.Route("/roles", func(role chi.Router) {
				role.Get("/", r.C.GetAllRoles)
				role.Get("/{id}", r.C.GetRole)
				role.Post("/", r.C.CreateRole)
				role.Delete("/{id}", r.C.DeleteRole)
			})
		})
	})
}

func (r *Router) Handler() http.Handler {
	return r.R
}
	
	