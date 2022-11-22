package main

import (
	"fmt"
	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
	"github.com/namsral/flag"
	"log"
	"net/http"
	"boilerplate/app"
	authApi "boilerplate/gateway/api/auth"
	"boilerplate/gateway/httpx"
	"boilerplate/gateway/state"
	"boilerplate/registry"
	"boilerplate/users"
	"boilerplate/utils"
)

func main() {
	flag.Parse()

	r := mux.NewRouter()
	srv := app.NewServer(registry.ServiceAddress("gateway"), r)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{*utils.ClientUrl},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	deps := state.New()

	r.Handle("/internal/signup", httpx.Internal(deps, authApi.Signup(deps, users.UserTypeInternal)))

	r.Handle("/v1/login", httpx.Base(authApi.Login(deps)))
	r.Handle("/v1/signup", httpx.Base(authApi.Signup(deps, users.UserTypeCustomer)))
	r.Handle("/v1/signout", httpx.Authenticated(deps, authApi.Signout(deps)))

	r.Handle("/v1/reset_password", httpx.Base(authApi.ResetPassword(deps, users.UserTypeCustomer)))

	r.Handle("/me", httpx.Authenticated(deps,
		func(w http.ResponseWriter, r *httpx.Request) {
			_, err := fmt.Fprintf(w, "authenticated route accessed by %d", r.UserId)
			if err != nil {
				return
			}
		}))

	go srv.ListenAndServeForever()
	log.Printf("Server listening at %s", srv.Address())

	srv.WaitForShutdown()
}
