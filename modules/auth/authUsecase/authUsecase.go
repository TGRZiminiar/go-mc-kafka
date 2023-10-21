package authusecase

import "github.com/TGRZiminiar/go-mc-kafka/modules/auth/authRepository"

type (
	AuthUseCaseService interface{}

	authUseCase struct {
		authRepository authRepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.AuthRepositoryService) AuthUseCaseService {
	return &authUseCase{authRepository: authRepository}
}
