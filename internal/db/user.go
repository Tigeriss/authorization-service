package db

import (
	"authorization-service/cmd/authorization/stmt"
	"authorization-service/internal/entity"
	"context"
	"fmt"
)

func CreateUser(ctx context.Context, login, passwordHash string) error {

	_, err := stmt.User.Create.ExecContext(ctx, login, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func ReadUser(ctx context.Context, id int64) (entity.User, error) {
	var user entity.User

	row := stmt.User.Read.QueryRowContext(ctx, id)

	err := row.Scan(&user)
	if err != nil {
		return user, fmt.Errorf("unable to read user: %w", err)
	}

	return user, nil
}

func UpdateUser(ctx context.Context, login, passwordHash string) error {

	_, err := stmt.User.Update.ExecContext(ctx, passwordHash, login)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(ctx context.Context, login string) error {

	_, err := stmt.User.Delete.ExecContext(ctx, login)
	if err != nil {
		return err
	}

	return nil
}
