package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"belajar-go/src/config/query"
	"belajar-go/src/domain"
	"belajar-go/src/dto"
)

func (d *userRepository) findAllUserFromSql(filter *dto.UserFilter) ([]domain.User, int, error) {
	var results []domain.User
	var totalData int

	template, args, err := d.queryLoader.ExecuteTemplate("FindAllUserData", filter)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to load query template: %w", err)
	}

	queryC, argsC, err := d.queryLoader.ExecuteTemplate("CountDataUser", filter)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to load count template: %w", err)
	}

	if err = d.sql0.Select(&results, template, args...); err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}

	if err = d.sql0.Get(&totalData, queryC, argsC...); err != nil {
		return nil, 0, fmt.Errorf("failed to execute count query: %w", err)
	}

	return results, totalData, nil
}

func (d *userRepository) createUserFromSql(ctx context.Context, users []*domain.UserCreateDomain) ([]*domain.UserCreateDomain, error) {
	const maxBatchSize = 100

	if len(users) == 0 {
		return nil, fmt.Errorf("users cannot be empty")
	}

	if len(users) > maxBatchSize {
		return nil, fmt.Errorf("maximum batch allowed is %d", maxBatchSize)
	}

	begin, err := d.sql0.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			cerr := begin.Rollback()
			if cerr != nil {
				return
			}
			return
		}
		err = begin.Commit()
	}()

	queryStr, ok := d.queryLoader.Get("CreateUser")
	if !ok {
		return nil, errors.New("query CreateUser not found")
	}

	tqlaCompile, args, err := query.CompileTqla(queryStr, users)
	if err != nil {
		return nil, fmt.Errorf("failed to compile query: %w", err)
	}

	_, err = begin.ExecContext(ctx, tqlaCompile, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return users, nil
}
