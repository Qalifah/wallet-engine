package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"wallet-engine/database/mongodb"
	"wallet-engine/handler"
	"wallet-engine/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()
	databaseURI := os.Getenv("DATABASE_URI")
	dbClient, err := mongodb.New(ctx, databaseURI)
	if err != nil {
		log.Fatalf("unable to create database client %v", err)
	}
	defer dbClient.Disconnect(ctx)

	db := dbClient.Database("wallet_engine")
	walletRepo := mongodb.NewWalletRepository(db)
	ctrl := handler.New(walletRepo)

	r := router.Init(ctrl)
	log.Fatal(r.Run(os.Getenv("PORT")))
}
