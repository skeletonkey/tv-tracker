package db

import (
	_ "github.com/mattn/go-sqlite3"

	"github.com/google/uuid"
)

func createUser(username, email, password string) (string, error) {
	conn := getDb()

	userId := uuid.New().String()
	_, err := conn.Exec(`
                INSERT INTO users (uuid, username, email, password)
                VALUES (?,?,?,?)
        `, userId, username, email, password)
	if err != nil {
		return "", err
	}
	return userId, nil
}
