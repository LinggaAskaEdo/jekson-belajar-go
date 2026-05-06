package repository

import (
	"belajar-go/src/config/query"
	"belajar-go/src/repository/user"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User user.UserRepositoryInterface
}

func InitRepository(sql0 *sqlx.DB, ql *query.LoadQuery) *Repository {
	return &Repository{
		User: user.InitUserRepository(
			sql0,
			ql,
		),
	}
}
