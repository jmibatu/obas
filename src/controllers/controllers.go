package controllers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"obas/src/config"
	controllers4 "obas/src/controllers/address"
	controllers6 "obas/src/controllers/application"
	controllers5 "obas/src/controllers/demographics"
	controllers7 "obas/src/controllers/documents"
	controllers "obas/src/controllers/home"
	controllers8 "obas/src/controllers/institutions"
	controllers9 "obas/src/controllers/location"
	controllers10 "obas/src/controllers/log"
	controllers3 "obas/src/controllers/registration"
	controllers2 "obas/src/controllers/subjects"
	controllers11 "obas/src/controllers/users"

	"obas/src/controllers/login"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)

	mux.Handle("/", controllers.Home(env))
	mux.Mount("/login", login.Login(env))
	mux.Mount("/users", controllers11.Users(env))
	mux.Mount("/subjects", controllers2.Subjects(env))
	mux.Mount("/registration", controllers3.Registrations(env))
	mux.Mount("/address", controllers4.Addresses(env))
	mux.Mount("/demographics", controllers5.Demographics(env))
	mux.Mount("/application", controllers6.Applications(env))
	mux.Mount("/documents", controllers7.Documents(env))
	mux.Mount("/institution", controllers8.Institutions(env))
	mux.Mount("/location", controllers9.Locations(env))
	mux.Mount("/log", controllers10.Logs(env))

	fileServer := http.FileServer(http.Dir("./src/views/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux

}
