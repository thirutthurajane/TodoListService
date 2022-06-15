package route

import (
	"github.com/go-chi/chi/v5"
	"todoservice/controller"
)

// var tokenAuth *jwtauth.JWTAuth

func Setup(router *chi.Mux) {
	/*tokenAuth = jwtauth.New("HS256", []byte(config.Configuration.Server.Secret), nil)
	port := ""
	switch configs.Configuration.Server.Env {
	case "develop":
		port = configs.Configuration.Server.Develop.Port
		router.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://localhost"+port+"/swagger/doc.json"), //The url pointing to API definition
		))
	case "production":
		port = config.Configuration.Server.Production.Port
	default:
		port = ":9999"
	}*/

	router.Group(func(router chi.Router) {

		/*	router.Use(jwtauth.Verifier(tokenAuth))
			router.Use(auth.Authenticator)*/

		router.Route("/api/todos", func(r chi.Router) {
			r.Get("/", controller.GetAllTodos)
		})
	})

	router.Group(func(r chi.Router) {
		r.Get("/auth/github/callback", controller.GithubAuthCallback)
	})

	/*router.Group(func(r chi.Router) {
		r.Post("/api/uid/auth/sign", user.Authentication)
	})*/

}
