package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main()  {
	fmt.Println("Hello, world!")

	godotenv.Load()

	// godotenv.Load(".env")
	portString := os.Getenv("PORT")
	
	if portString == "" {
        log.Fatal("could not find port")
    }

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	// v1Router.HandleFunc("/healthz", handler_readiness) this servers get and post both
	v1Router.Get("/healthz", handler_readiness)
	v1Router.Get("/err", handler_err)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	fmt.Println("Port: ", portString)

	err := srv.ListenAndServe()

	if err!= nil {
        log.Fatal(err)
    }
}