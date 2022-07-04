package db

import (
	"authorization-service/internal/entity"
	"context"
	"database/sql"
	"fmt"
)

func CreateUser(ctx context.Context, db *sql.DB, login, passwordHash string) error {
	createUserStmt, err := db.Prepare(`INSERT INTO users VALUES($1, $2)`)
	if err != nil {
		return err
	}

	_, err = createUserStmt.ExecContext(ctx, login, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func ReadUser(ctx context.Context, db *sql.DB, id int64) (entity.User, error) {
	var user entity.User

	readUserStmt, err := db.Prepare(`SELECT * FROM users WHERE id = $1`)
	if err != nil {
		return user, err
	}

	row := readUserStmt.QueryRowContext(ctx, id)

	err = row.Scan(&user)
	if err != nil {
		return user, fmt.Errorf("unable to read user: %w", err)
	}

	return user, nil
}

func UpdateUser(ctx context.Context, db *sql.DB, login, passwordHash string) error {
	updateUserStmt, err := db.Prepare(`UPDATE users SET password_hash = $1 WHERE login = $2`)
	if err != nil {
		return err
	}

	_, err = updateUserStmt.ExecContext(ctx, passwordHash, login)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(ctx context.Context, db *sql.DB, login string) error {
	deleteUserStmt, err := db.Prepare(`DELETE FROM users WHERE login = $1`)
	if err != nil {
		return err
	}

	_, err = deleteUserStmt.ExecContext(ctx, login)
	if err != nil {
		return err
	}

	return nil
}
