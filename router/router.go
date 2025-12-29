package router

import (
	"KageNoEn/controller"
	"net/http"

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
	r.R.Group(func(public chi.Router) {
		public.Route("/api/v1/login", func(login chi.Router) {
			login.Post("/", r.C.SignIn)
		})
		public.Route("/api/v1/register", func(login chi.Router) {
			login.Post("/", r.C.SignUp)
		})
	})

	r.R.Group(func(private chi.Router) {
		// private.Use(lib.JWTMiddleware)
		private.Route("/api/v1", func(v1 chi.Router) {
			v1.Route("/roles", func(role chi.Router) {
				role.Get("/", r.C.GetAllRoles)
				role.Get("/{id}", r.C.GetRole)
				role.Post("/", r.C.CreateRole)
				role.Delete("/{id}", r.C.DeleteRole)
			})
			v1.Route("/ranks", func(rank chi.Router) {
				rank.Get("/", r.C.GetAllRanks)
				rank.Post("/", r.C.CreateRank)
				rank.Put("/{id}", r.C.UpdateRank)
				rank.Delete("/{id}", r.C.DeleteRank)
			})
			v1.Route("/user-status", func(status chi.Router) {
				status.Get("/", r.C.GetAllUserStatus)
				status.Post("/", r.C.CreateUserStatus)
				status.Put("/{id}", r.C.UpdateUserStatus)
				status.Delete("/{id}", r.C.DeleteUserStatus)
			})
			v1.Route("/usr", func(user chi.Router) {
				user.Get("/", r.C.GetAllUser)
				user.Put("/{id}", r.C.ChangePass)
			})
			v1.Route("/logout", func(logout chi.Router) {
				logout.Put("/{id}", r.C.Logout)
			})
			v1.Route("/friends", func(friend chi.Router) {
				friend.Get("/{id}", r.C.GetAllFriends)
				friend.Get("/invit/{id}", r.C.GetRequest)
				friend.Get("/search/{username}", r.C.GetFiltredSearch)
				friend.Post("/", r.C.AddFriend)
				friend.Put("/confirm/{id}", r.C.ConfirmFriend)
			})
		})
	})
}

func (r *Router) Handler() http.Handler {
	return r.R
}
