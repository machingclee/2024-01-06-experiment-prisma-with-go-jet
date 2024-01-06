package main

import (
	"authentication/pkg/pgsql"
	"authentication/pkg/user"
	"log"
	"net/http"
)

func main() {
	db := pgsql.NewDB()
	user.NewHandler(db)

	srv := http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
