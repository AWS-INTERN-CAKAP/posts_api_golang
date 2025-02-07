package main

import (
	"github.com/Davidafdal/e-commerce/config"
	"github.com/Davidafdal/e-commerce/internal/builder"
	"github.com/Davidafdal/e-commerce/pkg/postgres"
	"github.com/Davidafdal/e-commerce/pkg/server"
	"github.com/Davidafdal/e-commerce/pkg/token"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	db, err := postgres.InitPostgres(&cfg.Postgres)
	checkError(err)


	tokenUseCase := token.NewTokenUseCase(cfg.JWTSecretKey)

	publicRoutes := builder.BuildAppPublicRoutes(db, tokenUseCase)
	privateRoutes := builder.BuildAppPrivateRoutes(db, tokenUseCase)

	srv := server.NewServer(publicRoutes, privateRoutes, cfg.JWTSecretKey, tokenUseCase)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}