package migration

import (
	"context"
	"log"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	"github.com/TGRZiminiar/go-mc-kafka/modules/player"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/database"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func playerDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("player_db")
}

func PlayerMigrate(pctx context.Context, cfg *config.Config) {
	db := playerDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("player_transactions")

	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{
			Keys: bson.D{{"_id", 1}},
		},
		{
			Keys: bson.D{{"player_id", 1}},
		},
	})

	for _, index := range indexs {
		log.Printf("Indexs: %s", index)
	}

	col = db.Collection("players")

	indexs, _ = col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{
			Keys: bson.D{{"_id", 1}},
		},
		{
			Keys: bson.D{{"email", 1}},
		},
	})

	for _, index := range indexs {
		log.Printf("Indexs: %s", index)
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)

	// roles data
	document := func() []any {
		roles := []*player.Player{
			{
				Email:    "player001@gmail.com",
				Password: string(hashPassword),
				Username: "player001",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "player002@gmail.com",
				Password: string(hashPassword),
				Username: "player002",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "player003@gmail.com",
				Password: string(hashPassword),
				Username: "player003",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "admin@gmail.com",
				Password: string(hashPassword),
				Username: "admin",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
					{
						RoleTitle: "admin",
						RoleCode:  1,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
		}

		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, document, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate Player Complete", results)

	playerTransaction := make([]any, 0)
	for _, p := range results.InsertedIDs {
		playerTransaction = append(playerTransaction, &player.PlayerTransaction{
			PlayerId:  "player:" + p.(primitive.ObjectID).Hex(),
			Amount:    10000,
			CreatedAt: utils.LocalTime(),
		})
	}
	col = db.Collection("player_transactions")
	results, err = col.InsertMany(pctx, playerTransaction, nil)
	if err != nil {
		panic(err)
	}

	log.Println("Migrate Player Transaction Complete", results)

	col = db.Collection("player_transaction_queue")
	result, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate Player Transaction Queue Complete", result)

}
