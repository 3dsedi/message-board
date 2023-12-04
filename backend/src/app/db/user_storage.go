package db

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/sedi/messageBoard/model"
)

type UserStorage interface {
	CreateUser(*model.User) error
	GetUsers() ([]*model.User, error)
	GetUserByEmailAndPassword(email, password string) (*model.User, error)
	DeleteUser(userID string) error
}

type PostgresUserStorage struct {
	db *sql.DB
}

func NewPostgresUserStorage() *PostgresUserStorage {
	return &PostgresUserStorage{db: conn}
}

func (s *PostgresUserStorage) CreateUserTable() error {
	query := `
	CREATE EXTENSION IF NOT EXISTS "pgcrypto";
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_name VARCHAR(50) UNIQUE,
		email VARCHAR(50),
		password VARCHAR(50),
		created_at TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	if err != nil {
		log.Fatal("error creating users table: %w", err)
		return model.MessageBoardError{
			Message: "can not create table users",
			Code:    http.StatusInternalServerError,
		}
	}
	log.Println("Users table created successfully!")

	return nil
}

func (s *PostgresUserStorage) Init() error {
	if err := s.CreateUserTable(); err != nil {
		return err
	}

	return nil
}

func (s *PostgresUserStorage) CreateUser(user *model.User) error {
	var resp sql.Result
	var err error

	query := `INSERT INTO users (user_name, email, password, created_at)
			VALUES ($1, $2, $3, $4)`

	resp, err = s.db.Exec(query, user.UserName, user.Email, user.Password, user.CreatedAt)

	if err != nil {
		log.Fatal("error inserting user into database: %w", err)
		//TODO(Sedigheh): should seperate duplicate from the other errors
		return model.DuplicateHanlder(user.Email, "CreateUser")
	}
	rowsAffected, _ := resp.RowsAffected()
	log.Printf("add user successfully. Rows affected: %d\n", rowsAffected)

	return nil
}

func (s *PostgresUserStorage) GetUsers() ([]*model.User, error) {
	query := `SELECT * FROM users`

	rows, err := s.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.NotFoundErrHanlder("users", "", "GetUsers")
		}

		return nil, model.FatalErrorHandler(err, "GetUsers")
	}

	users := []*model.User{}
	for rows.Next() {
		user := new(model.User)
		err := rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, model.FatalErrorHandler(err, "GetUsers")
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("error iterating over users rows: %w", err)
		return nil, model.FatalErrorHandler(err, "GetUsers")
	}

	log.Printf("Retrieved %d users from database", len(users))
	return users, nil
}

func (s *PostgresUserStorage) GetUserByEmailAndPassword(email string, password string) (*model.User, error) {
	query := `SELECT id, user_name, email, password, created_at FROM users WHERE email = $1 AND password = $2`

	user := model.User{}
	err := s.db.QueryRow(query, email, password).Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {

		if err == sql.ErrNoRows {
			log.Printf("User %s not found", email)
			return nil, model.NotFoundErrHanlder("email", email, "GetUserByEmailAndPassword")
		}

		return nil, model.FatalErrorHandler(err, "GetUserByEmailAndPassword")
	}

	return &user, nil
}

func (s *PostgresUserStorage) DeleteUser(userID string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := s.db.Exec(query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User with ID %s not found", userID)
			return model.NotFoundErrHanlder("user", userID, "DeleteUser")
		}
		return model.FatalErrorHandler(err, "DeleteUser")
	}

	log.Printf("Deleted user with ID %s", userID)
	return nil
}
