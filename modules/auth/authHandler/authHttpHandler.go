package authhandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	"github.com/TGRZiminiar/go-mc-kafka/modules/auth"
	authusecase "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authUsecase"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/request"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	AuthHttpHandlerService interface {
		Login(c echo.Context) error
		RefreshToken(c echo.Context) error
		Logout(c echo.Context) error
	}

	authHttpHandler struct {
		cfg         *config.Config
		authUseCase authusecase.AuthUseCaseService
	}
)

func NewAuthHttpHandler(pcfg *config.Config, authUseCase authusecase.AuthUseCaseService) AuthHttpHandlerService {
	return &authHttpHandler{
		authUseCase: authUseCase,
		cfg:         pcfg,
	}
}

func (h *authHttpHandler) Login(c echo.Context) error {

	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(auth.PlayerLoginReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUseCase.Login(ctx, h.cfg, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *authHttpHandler) RefreshToken(c echo.Context) error {

	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(auth.RefreshTokenReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUseCase.RefreshToken(ctx, h.cfg, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *authHttpHandler) Logout(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(auth.LogoutReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUseCase.Logout(ctx, req.CredentialId)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, &response.MsgResponse{
		Message: fmt.Sprintf("Deleted Count %d", res),
	})

}
