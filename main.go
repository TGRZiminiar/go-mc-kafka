package main

import (
	"context"
	"log"
	"os"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/database"
)

func main() {
	ctx := context.Background()
	_ = ctx

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 1 {
			log.Fatal("Error: .env path is invalid")
		}
		return os.Args[1]
	}())

	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	log.Println(db)
}
