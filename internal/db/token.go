package db

import (
	"authorization-service/cmd/authorization/stmt"
	"authorization-service/internal/entity"
	"context"
	"fmt"
	"time"
)

func CreateToken(ctx context.Context, token string, userID int64, expired time.Time) error {

	var idExists bool
	row := stmt.Token.Read.QueryRowContext(ctx, userID)

	err := row.Scan(&idExists)
	if err != nil {
		return err
	}

	if idExists {
		_, err = stmt.Token.Create.ExecContext(ctx, token, userID, expired)
		if err != nil {
			return err
		}

	} else {
		return err
	}

	return nil
}

func ReadToken(ctx context.Context, tok string) (entity.Token, error) {
	var token entity.Token

	row := stmt.Token.Read.QueryRowContext(ctx, tok)

	err := row.Scan(&token)
	if err != nil {
		return token, fmt.Errorf("unable to read token: %w", err)
	}

	return token, nil
}

func DeleteToken(ctx context.Context, token string) error {

	_, err := stmt.Token.Delete.ExecContext(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
