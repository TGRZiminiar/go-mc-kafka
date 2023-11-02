package authRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/TGRZiminiar/go-mc-kafka/modules/auth"
	playerPb "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerPb"
	grpcconn "github.com/TGRZiminiar/go-mc-kafka/pkg/grpcConn"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthRepositoryService interface {
		InsertOnePlayerCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error)
		CredentialSearch(pctx context.Context, grpcUrl string, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error)
	}

	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{db}
}

func (r *authRepository) authDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}

func (r *authRepository) CredentialSearch(pctx context.Context, grpcUrl string, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := grpcconn.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: Grpc Conn Error %s", err.Error())
		return nil, errors.New("error: grpc connection failed")
	}
	// Calling CredentialSearch in file playerGrpcHandler
	result, err := conn.Player().CredentialSearch(ctx, req)
	if err != nil {
		log.Printf("Error: Credential Search Error %s", err.Error())
		return nil, errors.New("error: credentail search failed")

	}

	return result, nil
}

func (r *authRepository) InsertOnePlayerCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: Insert One Player Credential Error %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one player credential failed")
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
