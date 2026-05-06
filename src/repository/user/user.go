package user

import (
	"context"

	"belajar-go/src/config/query"
	"belajar-go/src/domain"
	"belajar-go/src/dto"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryInterface interface {
	FindAll(filter *dto.UserFilter) ([]domain.User, int, error)
	Create(ctx context.Context, user []*domain.UserCreateDomain) ([]*domain.UserCreateDomain, error)
}

type userRepository struct {
	sql0        *sqlx.DB
	queryLoader *query.LoadQuery
}

func InitUserRepository(sql0 *sqlx.DB, ql *query.LoadQuery) UserRepositoryInterface {
	return &userRepository{
		sql0:        sql0,
		queryLoader: ql,
	}
}
