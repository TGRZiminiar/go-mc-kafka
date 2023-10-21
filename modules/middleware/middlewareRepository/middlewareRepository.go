package middlewarerepository

type (
	MiddlewareRepositoryService interface{}

	MiddlewareRepository struct {
	}
)

func NewMiddlewareRepository() MiddlewareRepositoryService {
	return &MiddlewareRepository{}
}
