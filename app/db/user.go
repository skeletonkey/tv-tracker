package db

import (
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

const bcryptCost = 10

func CreateUser(username, email, password string) (string, error) {
	conn := getDb()

	encPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}
	password = string(encPassword)

	userId := uuid.New().String()

	_, err = conn.Exec(`
                INSERT INTO user (uuid, username, email, password)
                VALUES (?,?,?,?)
        `, userId, username, email, password)
	if err != nil {
		return "", err
	}
	return userId, nil
}
