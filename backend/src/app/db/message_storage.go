package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/sedi/message-board/model"
)

type MessageStorage interface {
	GetMessageByUserID(UserID string) ([]*model.Message, error)
	CreateMessage(*model.Message) error
	DeleteMessage(id int) error
	UpdateMessage(*model.Message) error
	GetMainMessages() ([]*model.Message, error)
	GetMessageReplies(id int) ([]*model.Message, error)
}

type PostgresMessageStorage struct {
	db *sql.DB
}

func NewPostgresMessageStorage() *PostgresMessageStorage {
	return &PostgresMessageStorage{db: conn}
}

func (s *PostgresMessageStorage) Init() error {
	if err := s.CreateMessageTable(); err != nil {
		return err
	}

	return nil
}

func (s *PostgresMessageStorage) CreateMessageTable() error {
	query := `CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		user_id UUID REFERENCES users(id) ON DELETE CASCADE,
		content TEXT,
		parent_id INT REFERENCES messages(id) ON DELETE CASCADE,
		created_at TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	if err != nil {
		log.Fatalf("error creating message table: %s", err)
		return model.FatalErrorHandler(err, "CreateMessageTable")
	}
	log.Println("Messages table created successfully!")

	return nil
}

func (s *PostgresMessageStorage) CreateMessage(message *model.Message) error {
	var resp sql.Result
	var err error

	log.Printf("Inserting message into database: %v\n", message)

	if message.ParentID != nil {
		query := `INSERT INTO messages (user_id, content, parent_id, created_at)
			VALUES ($1, $2, $3, $4)`

		resp, err = s.db.Exec(query, message.UserID, message.Content, *message.ParentID, message.CreatedAt)
	} else {
		query := `INSERT INTO messages (user_id, content, created_at)
			VALUES ($1, $2, $3)`

		resp, err = s.db.Exec(query, message.UserID, message.Content, message.CreatedAt)
	}

	if err != nil {
		log.Fatal("error inserting message into database: %w", err)
		return model.FatalErrorHandler(err, "CreateMessageTable")

	}
	rowsAffected, _ := resp.RowsAffected()
	log.Printf("Inserted message successfully. Rows affected: %d\n", rowsAffected)

	return nil
}

func (s *PostgresMessageStorage) GetMessageByUserID(userID string) ([]*model.Message, error) {
	query := `SELECT m.id, m.user_id, u.user_name, m.content, m.parent_id, m.created_at 
	FROM messages m
	INNER JOIN users u ON m.user_id = u.id
	WHERE m.user_id = $1`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no message found")
			return nil, model.NotFoundErrHanlder("user", userID, "GetMessageByUserID")
		}
		return nil, model.FatalErrorHandler(err, "GetMessageByUserID")
	}
	messages := []*model.Message{}
	for rows.Next() {
		msg := new(model.Message)
		err := rows.Scan(&msg.ID, &msg.UserID, &msg.UserName, &msg.Content, &msg.ParentID, &msg.CreatedAt)
		if err != nil {
			return nil, model.FatalErrorHandler(err, "GetMessageByUserID")
		}
		messages = append(messages, msg)
	}
	if err := rows.Err(); err != nil {
		return nil, model.FatalErrorHandler(err, "GetMessageByUserID")
	}

	log.Printf("Retrieved %d message replies for message ID: %s", len(messages), userID)
	return messages, nil	

}

func (s *PostgresMessageStorage) GetMainMessages() ([]*model.Message, error) {
	query := `SELECT m.id, m.user_id, u.user_name, m.content, m.parent_id, m.created_at 
              FROM messages m
              INNER JOIN users u ON m.user_id = u.id
              WHERE m.parent_id IS NULL `

	rows, err := s.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("no message found")
			return nil, model.NotFoundErrHanlder("main", "", "GetMainMessages")
		}
		return nil, model.FatalErrorHandler(err, "GetMainMessages")
	}
	messages := []*model.Message{}
	for rows.Next() {
		message := new(model.Message)
		err := rows.Scan(&message.ID, &message.UserID, &message.UserName, &message.Content, &message.ParentID, &message.CreatedAt)
		if err != nil {
			return nil, model.FatalErrorHandler(err, "GetMainMessages")
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, model.FatalErrorHandler(err, "GetMainMessages")
	}

	log.Printf("Retrieved %d messages from database", len(messages))
	return messages, nil
}

func (s *PostgresMessageStorage) GetMessageReplies(id int) ([]*model.Message, error) {
	//query := `SELECT id,user_id, content, parent_id, created_at FROM messages WHERE parent_id = $1`
	query := `SELECT m.id, m.user_id, u.user_name, m.content, m.parent_id, m.created_at 
	FROM messages m
	INNER JOIN users u ON m.user_id = u.id
	WHERE m.parent_id = $1`

	rows, err := s.db.Query(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.NotFoundErrHanlder("id", fmt.Sprint(id), "GetMessageReplies")
		}
		return nil, model.FatalErrorHandler(err, "GetMessageReplies")
	}

	replies := []*model.Message{}
	for rows.Next() {
		message := new(model.Message)
		err := rows.Scan(&message.ID, &message.UserID, &message.UserName, &message.Content, &message.ParentID, &message.CreatedAt)
		if err != nil {
			return nil, model.FatalErrorHandler(err, "GetMessageReplies")
		}
		replies = append(replies, message)
	}
	if err := rows.Err(); err != nil {
		return nil, model.FatalErrorHandler(err, "GetMessageReplies")
	}

	log.Printf("Retrieved %d message replies for message ID: %d", len(replies), id)
	return replies, nil
}

func (s *PostgresMessageStorage) DeleteMessage(id int) error {
	query := `DELETE FROM messages WHERE id = $1`
	result, err := s.db.Exec(query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			model.NotFoundErrHanlder("id", fmt.Sprint(id), "UpdateMessage")
		}
		model.FatalErrorHandler(err, "UpdateMessage")
	}

	numRowsAffected, _ := result.RowsAffected()
	if numRowsAffected == 0 {
		return fmt.Errorf("message with ID %d not found", id)
	}

	log.Printf("Deleted message with ID %d. Rows affected: %d", id, numRowsAffected)
	return nil
}

func (s *PostgresMessageStorage) UpdateMessage(message *model.Message) error {
	query := `UPDATE messages SET content = $1 WHERE id = $2`
	result, err := s.db.Exec(query, message.Content, message.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			model.NotFoundErrHanlder("id", fmt.Sprint(message.ID), "UpdateMessage")
		}
		model.FatalErrorHandler(err, "UpdateMessage")
	}

	numRowsAffected, _ := result.RowsAffected()
	if numRowsAffected == 0 {
		return fmt.Errorf("message with ID %d not found", message.ID)
	}
	log.Printf("Updated message with ID %d. Rows affected: %d", message.ID, numRowsAffected)
	return nil
}
