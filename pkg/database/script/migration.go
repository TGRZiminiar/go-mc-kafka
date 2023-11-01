package main

import (
	"context"
	"log"
	"os"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/database/migration"
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

	switch cfg.App.Name {
	case "auth":
		migration.AuthMigrate(ctx, &cfg)
	case "player":
		migration.PlayerMigrate(ctx, &cfg)

	case "item":
		migration.ItemMigrate(ctx, &cfg)
	case "inventory":
		migration.InventoryMigrate(ctx, &cfg)
	case "payment":
		migration.PaymentMigrate(ctx, &cfg)

	}

}
