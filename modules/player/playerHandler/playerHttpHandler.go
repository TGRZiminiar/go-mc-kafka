package playerhandler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	"github.com/TGRZiminiar/go-mc-kafka/modules/player"
	playerusecase "github.com/TGRZiminiar/go-mc-kafka/modules/player/playerUsecase"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/request"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	PlayerHttpHandlerService interface {
		CreatePlayer(c echo.Context) error
		FindOnePlayerProfile(c echo.Context) error
		AddPlayerMoney(c echo.Context) error
	}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewPlayerHttpHandler(cfg *config.Config, playerUsecase playerusecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{cfg, playerUsecase}
}

func (h *playerHttpHandler) CreatePlayer(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(player.CreatePlayerReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.playerUsecase.CreatePlayer(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, map[string]any{"msg": "ok", "userId": res})
}

func (h *playerHttpHandler) FindOnePlayerProfile(c echo.Context) error {
	ctx := context.Background()

	// player:(playerId)
	playerId := strings.TrimPrefix(c.Param("player_id"), "player:")

	res, err := h.playerUsecase.FindOnePlayerProfile(ctx, playerId)
	if err != nil {
		return response.SuccessResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *playerHttpHandler) AddPlayerMoney(c echo.Context) error {
	ctx := context.Background()
	fmt.Println("hello world")
	wrapper := request.NewContextWrapper(c)

	req := new(player.CreatePlayerTransactionReq)

	if err := wrapper.Bind(req); err != nil {
		fmt.Println("Something Weonge", err)
		fmt.Println("Fuckerrrr")
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	err := h.playerUsecase.AddPlayerMoney(ctx, req)
	if err != nil {
		fmt.Println("Something Weong2e", err)
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, map[string]any{"msg": "success"})
}
