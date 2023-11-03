package middlewareusecase

import (
	"errors"

	"github.com/TGRZiminiar/go-mc-kafka/config"
	middlewarerepository "github.com/TGRZiminiar/go-mc-kafka/modules/middleware/middlewareRepository"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/jwtauth"
	"github.com/TGRZiminiar/go-mc-kafka/pkg/rbac"
	"github.com/labstack/echo/v4"
)

type (
	MiddlewareUsecaseService interface {
		JwtAuthorization(c echo.Context, cfg *config.Config, accessToken string) (echo.Context, error)
		RbacAuthorization(c echo.Context, cfg *config.Config, expected []int) (echo.Context, error)
		PlayerIdParamValidation(c echo.Context) (echo.Context, error)
	}

	MiddlewareUsecase struct {
		middlewareRepository middlewarerepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewarerepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &MiddlewareUsecase{middlewareRepository: middlewareRepository}
}

func (u *MiddlewareUsecase) JwtAuthorization(c echo.Context, cfg *config.Config, accessToken string) (echo.Context, error) {

	ctx := c.Request().Context()

	cliams, err := jwtauth.ParseToken(cfg.Jwt.AccessSecretKey, accessToken)
	if err != nil {
		return nil, err
	}

	if err := u.middlewareRepository.AccessTokenSearch(ctx, cfg.Grpc.AuthUrl, accessToken); err != nil {
		return nil, err
	}

	c.Set("player_id", cliams.PlayerId)
	c.Set("role_code", cliams.RoleCode)

	return c, nil
}

func (u *MiddlewareUsecase) RbacAuthorization(c echo.Context, cfg *config.Config, expected []int) (echo.Context, error) {
	ctx := c.Request().Context()

	playerRoleCode := c.Get("role_code").(int)

	rolesCount, err := u.middlewareRepository.RolesCount(ctx, cfg.Grpc.AuthUrl)
	if err != nil {
		return nil, err
	}

	playerRoleBinary := rbac.IntToBinary(playerRoleCode, int(rolesCount))
	for i := 0; i < int(rolesCount); i++ {
		if playerRoleBinary[i]&expected[i] == 1 {
			return c, nil
		}
	}

	return nil, errors.New("error: permission deny")
}

func (u *MiddlewareUsecase) PlayerIdParamValidation(c echo.Context) (echo.Context, error) {
	playerIdReq := c.Param("player_id")
	playerIdToken := c.Get("player_id").(string)

	if playerIdToken == "" {
		return nil, errors.New("error: playerId is required")
	}

	if playerIdToken != playerIdReq {
		return nil, errors.New("error: playerId not match")
	}

	return c, nil
}
