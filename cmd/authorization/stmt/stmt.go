package stmt

import (
	"context"
	"database/sql"
	"log"
)

var User struct {
	Create *sql.Stmt
	Read   *sql.Stmt
	Update *sql.Stmt
	Delete *sql.Stmt
}

var Token struct {
	Create *sql.Stmt
	Read   *sql.Stmt
	Check  *sql.Stmt
	Delete *sql.Stmt
}

func mustPrepare(ctx context.Context, db *sql.DB, sql string) *sql.Stmt {
	stmt, err := db.PrepareContext(ctx, sql)
	if err != nil {
		log.Panicf("unable to prepare db statemtent: %s", err)
	}
	return stmt
}

func PrepareStatements(ctx context.Context, db *sql.DB) {
	// User
	User.Create = mustPrepare(ctx, db, `INSERT INTO users VALUES($1, $2)`)
	User.Read = mustPrepare(ctx, db, `SELECT * FROM users WHERE id = $1`)
	User.Update = mustPrepare(ctx, db, `UPDATE users SET password_hash = $1 WHERE login = $2`)
	User.Delete = mustPrepare(ctx, db, `DELETE FROM users WHERE login = $1`)

	// Tokens
	Token.Create = mustPrepare(ctx, db, `INSERT INTO tokens (token, user_id, expired) VALUES($1, $2, $3)`)
	Token.Check = mustPrepare(ctx, db, `SELECT EXISTS(SELECT id FROM users WHERE id = $1)`)
	Token.Read = mustPrepare(ctx, db, `SELECT * FROM tokens WHERE token = $1`)
	Token.Delete = mustPrepare(ctx, db, `DELETE FROM tokens WHERE token = $1`)
}
