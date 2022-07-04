package db

import (
	"authorization-service/internal/entity"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func CreateToken(ctx context.Context, db *sql.DB, token string, userID int64, expired time.Time) error {

	var idExists bool
	isUserIDExistsStmt, err := db.Prepare(`SELECT EXISTS(SELECT id FROM users WHERE id = $1)`)
	if err != nil {
		return err
	}
	row := isUserIDExistsStmt.QueryRowContext(ctx, userID)

	err = row.Scan(&idExists)
	if err != nil {
		return err
	}

	if idExists {
		createTokenStmt, err := db.Prepare(`INSERT INTO tokens (token, user_id, expired) VALUES($1, $2, $3)`)
		if err != nil {
			return err
		}

		_, err = createTokenStmt.ExecContext(ctx, token, userID, expired)
		if err != nil {
			return err
		}

	} else {
		return err
	}

	return nil
}

func ReadToken(ctx context.Context, db *sql.DB, tok string) (entity.Token, error) {
	var token entity.Token

	readTokenStmt, err := db.Prepare(`SELECT * FROM tokens WHERE token = $1`)
	if err != nil {
		return token, err
	}

	row := readTokenStmt.QueryRowContext(ctx, tok)

	err = row.Scan(&tok)
	if err != nil {
		return token, fmt.Errorf("unable to read token: %w", err)
	}

	return token, nil
}

func DeleteToken(ctx context.Context, db *sql.DB, token string) error {
	deleteTokenStmt, err := db.Prepare(`DELETE FROM tokens WHERE token = $1`)
	if err != nil {
		return err
	}

	_, err = deleteTokenStmt.ExecContext(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
