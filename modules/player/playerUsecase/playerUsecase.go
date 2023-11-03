package playerusecase

import (
	"context"
	"errors"
	"log"
	"time"

	playerPb "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerPb"

	"github.com/TGRZiminiar/go-mc-kafka/modules/player"
	"github.com/TGRZiminiar/go-mc-kafka/modules/player/playerRepository"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	PlayerUsecaseService interface {
		CreatePlayer(pctx context.Context, req *player.CreatePlayerReq) (*player.PlayerProfile, error)
		FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfile, error)
		AddPlayerMoney(pctx context.Context, req *player.CreatePlayerTransactionReq) (*player.PlayerSavingAccount, error)
		GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error)
		FindOnePlayerCredentail(pctx context.Context, password, email string) (*playerPb.PlayerProfile, error)
		FindOnePlayerProfileToRefresh(pctx context.Context, playerId string) (*playerPb.PlayerProfile, error)
	}

	playerUsecase struct {
		playerRepository playerRepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase(playerRepository playerRepository.PlayerRepositoryService) PlayerUsecaseService {
	return &playerUsecase{playerRepository}
}

func (u *playerUsecase) CreatePlayer(pctx context.Context, req *player.CreatePlayerReq) (*player.PlayerProfile, error) {

	if !u.playerRepository.IsUniquePlayer(pctx, req.Email, req.Username) {
		return nil, errors.New("error: this email or username is already exist")
	}

	// HashPassword
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error: failed to hash password")
	}

	playerId, err := u.playerRepository.InsertOnePlayer(pctx, &player.Player{
		Email:     req.Email,
		Password:  string(hashPassword),
		Username:  req.Username,
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
		PlayerRoles: []player.PlayerRole{
			{
				RoleTitle: "player",
				RoleCode:  0,
			},
		},
	})

	if err != nil {
		return nil, errors.New("error: failed to create new player")
	}

	return u.FindOnePlayerProfile(pctx, playerId.Hex())
}

func (u *playerUsecase) FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfile, error) {
	result, err := u.playerRepository.FindOnePlayerProfile(pctx, playerId)
	if err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Printf("Error: FindOnePlayerProfile: %s", err.Error())
		return nil, errors.New("error: failed to load location")
	}

	return &player.PlayerProfile{
		Id:        result.Id.Hex(),
		Email:     result.Email,
		Username:  result.Username,
		CreatedAt: result.CreatedAt.In(loc),
		UpdatedAt: result.UpdatedAt.In(loc),
	}, nil
}

func (u *playerUsecase) AddPlayerMoney(pctx context.Context, req *player.CreatePlayerTransactionReq) (*player.PlayerSavingAccount, error) {
	_, err := u.playerRepository.InsertOnePlayerTransaction(pctx, &player.PlayerTransaction{
		PlayerId:  req.PlayerId,
		Amount:    req.Amount,
		CreatedAt: utils.LocalTime(),
	})
	if err != nil {
		return nil, err
	}
	return u.playerRepository.GetPlayerSavingAccount(pctx, req.PlayerId)
}

func (u *playerUsecase) GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error) {
	return u.playerRepository.GetPlayerSavingAccount(pctx, playerId)
}

func (u *playerUsecase) FindOnePlayerCredentail(pctx context.Context, password, email string) (*playerPb.PlayerProfile, error) {

	result, err := u.playerRepository.FindOnePlayerCredentail(pctx, email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return nil, errors.New("error: password invalid")
	}

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Printf("Error: FindOnePlayerProfile: %s", err.Error())
		return nil, errors.New("error: failed to load location")
	}

	roleCode := 0
	for _, v := range result.PlayerRoles {
		roleCode += v.RoleCode
	}

	return &playerPb.PlayerProfile{
		Id:        result.Id.Hex(),
		Email:     result.Email,
		RoleCode:  int32(roleCode),
		Username:  result.Username,
		CreatedAt: result.CreatedAt.In(loc).String(),
		UpdatedAt: result.UpdatedAt.In(loc).String(),
	}, nil
}
func (u *playerUsecase) FindOnePlayerProfileToRefresh(pctx context.Context, playerId string) (*playerPb.PlayerProfile, error) {
	result, err := u.playerRepository.FindOnePlayerProfileToRefresh(pctx, playerId)
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Bangkok")

	var rolesCode int = 0
	for _, v := range result.PlayerRoles {
		rolesCode += v.RoleCode
	}

	return &playerPb.PlayerProfile{
		Id:        result.Id.Hex(),
		Email:     result.Email,
		Username:  result.Username,
		RoleCode:  int32(rolesCode),
		CreatedAt: result.CreatedAt.In(loc).String(),
		UpdatedAt: result.UpdatedAt.In(loc).String(),
	}, nil
}
