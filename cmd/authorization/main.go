package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var dbConnStr = "postgres://authent_service:608011@localhost/auth_service?sslmode=disable"

type User struct {
	ID           int64  `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
}

type Token struct {
	Token   string    `json:"token"`
	UserID  int64     `json:"user_id"`
	Scope   string    `json:"scope"`
	Created time.Time `json:"created"`
	Expired time.Time `json:"expired"`
}

func heading(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("X-Powered-By", "Tigeriss' Engine")
		f(writer, request)
	}
}

func private(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("X-Token")
		if token != "MagicKey" {
			writer.WriteHeader(http.StatusForbidden)
			log.Println("wrong X-Token")
			return
		}
		f(writer, request)
	}
}

func lowercaseHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusForbidden)
		log.Println("Should ba a POST Method")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	receivedBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Can not read body of request")
		return
	}
	textToChange := string(receivedBytes)
	result := strings.ToLower(textToChange)
	log.Println(result)

	w.Write([]byte(result))
	return
}

func uppercaseHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusForbidden)
		log.Println("Should ba a POST Method")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	receivedBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Can not read body of request")
		return
	}
	textToChange := string(receivedBytes)
	result := strings.ToUpper(textToChange)
	log.Println(result)

	w.Write([]byte(result))
	return
}

func createUser(db *sql.DB, login, password_hash string) error {
	createUserStmt := `INSERT INTO users VALUES($1, $2)`
	_, err := db.Exec(createUserStmt, login, password_hash)

	if err != nil {
		log.Println("Error create user " + err.Error())
		return err
	}
	return err
}
func readUser(db *sql.DB, login string) (User, error) {
	var user User
	readUserStmt := `SELECT * FROM users WHERE login = $1`

	rows, err := db.Query(readUserStmt, login)
	if err != nil {
		log.Println("Error get user " + err.Error())
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user)
		if err != nil {
			log.Println("Error get data from row " + err.Error())
			return user, err
		}
	}
	return user, nil
}

func updateUser(db *sql.DB, login, password_hash string) error {
	updateUserStmt := `UPDATE users SET password_hash = $1 WHERE login = $2`
	_, err := db.Exec(updateUserStmt, password_hash, login)

	if err != nil {
		log.Println("Error update user " + err.Error())
		return err
	}
	return err
}

func deleteUser(db *sql.DB, login string) error {
	deleteUserStmt := `DELETE FROM users WHERE login = $1`

	_, err := db.Exec(deleteUserStmt, login)

	if err != nil {
		log.Println("Error delete user " + err.Error())
		return err
	}

	return err
}

func createToken(db *sql.DB, token string, userID int64, expired time.Time) error {
	createTokenStmt := `INSERT INTO tokens (token, user_id) VALUES($1, $2)`
	_, err := db.Exec(createTokenStmt, token, userID)

	if err != nil {
		log.Println("Error create token " + err.Error())
		return err
	}
	return err
}
func readToken(db *sql.DB, tok string) (Token, error) {
	var token Token
	readTokenStmt := `SELECT * FROM tokens WHERE token = $1`

	rows, err := db.Query(readTokenStmt, tok)
	if err != nil {
		log.Println("Error get token " + err.Error())
		return token, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tok)
		if err != nil {
			log.Println("Error get data from row " + err.Error())
			return token, err
		}
	}
	return token, nil
}
func deleteToken(db *sql.DB, token string) error {
	deleteTokenStmt := `DELETE FROM tokens WHERE token = $1`

	_, err := db.Exec(deleteTokenStmt, token)

	if err != nil {
		log.Println("Error delete token " + err.Error())
		return err
	}

	return err
}

func main() {

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Println("Error open db connection: " + err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Error ping db connection: " + err.Error())
	}
	log.Println("Connect to db successfully")

	http.HandleFunc("/api/private/lowercase", private(heading(lowercaseHandle)))
	http.HandleFunc("/api/uppercase", heading(uppercaseHandler))

	log.Fatal(http.ListenAndServe(":9090", nil))
}
