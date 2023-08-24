package models

import (
	"errors"

	"github.com/vickyshaw29/events/database"
	"github.com/vickyshaw29/events/utils"
)

type User struct {
	ID       int64  `binding:"required"`
	email    string `binding:"required"`
	password string `binding:"required"`
}

func (u User) Create() error {
	query := `INSERT INTO users(email, password) VALUES(?,?)`
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT password FROM users WHERE email = ?`
	row := database.DB.QueryRow(query, u.email)
	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}
	isValid := utils.CheckPasswordHash(u.password, retrievedPassword)
	if !isValid {
		return errors.New("invalid credentials")
	}
	return nil
}
