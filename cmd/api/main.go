package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	httpdelivery "server/internal/delivery/http"
	"server/internal/infra/crypto"
	"server/internal/infra/sqlite"
	"server/internal/usecase"
)

func main() {
	db, err := sqlite.NewDB("app.db")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := sqlite.NewUserRepo(db)
	hasher := crypto.BcryptHasher{}

	auth := usecase.NewAuthService(userRepo, hasher)
	handler := httpdelivery.NewHandler(auth)
	router := httpdelivery.NewRouter(handler)

	log.Println("listening on :8080")
	http.ListenAndServe(":8080", router)
}
