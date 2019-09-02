package repositories

import (
	"anonbot/domain"
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"time"
)

type MessageRepository interface {
	GetMessage(messageId int, ctx context.Context) (*domain.Message, error)
	InsertMessage(message *domain.Message, ctx context.Context) (int, error)
	SetDelivered(messageId int, ctx context.Context) error
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db}
}

func (mr *messageRepository) GetMessage(messageId int, ctx context.Context) (*domain.Message, error) {
	q := `
select 
	m.id,
	m.sender_source_system,
	m.sender_source_system_id,
	m.message_body,
	m.received_at,
	m.delivered_at
from message m
where m.id = $1`

	rows, err := mr.db.QueryContext(ctx, q, messageId)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	message := domain.Message {	}

	if rows.Next() {
		if err := rows.Scan(&message.MessageId,
			&message.SenderSourceSystem,
			&message.SenderSourceSystemId,
			&message.MessageBody,
			&message.ReceivedAt,
			&message.DeliveredAt); err != nil {
			return nil, errors.WithStack(err)
		}

		return &message, nil
	}

	return nil, nil
}

func (mr *messageRepository) InsertMessage(message *domain.Message, ctx context.Context) (int, error) {
	q := `
insert into message(sender_source_system, sender_source_system_id, message_body, received_at, delivered_at)
values ($1, $2, $3, $4, $5)
returning id`


	rows, err := mr.db.QueryContext(ctx, q,
		message.SenderSourceSystem,
		message.SenderSourceSystemId,
		message.MessageBody,
		message.ReceivedAt,
		message.DeliveredAt)

	id := 0
	if err != nil {
		return id, errors.WithStack(err)
	}

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return id, errors.WithStack(err)
		}

		return id, nil
	}

	return id, errors.New("Error inserting message")
}

func (mr *messageRepository) SetDelivered(messageId int, ctx context.Context) error {
	q := `
update message m
set
	delivered_at = $2
where m.id = $1`

	result, err := mr.db.ExecContext(ctx, q, messageId, time.Now().UTC())
	if err != nil {
		return errors.WithStack(err)
	}

	if n, err := result.RowsAffected(); err != nil {
		return errors.WithStack(err)
	} else {
		if n < 1 {
			return errors.New("Error setting delivered for messageId")
		}

		return nil
	}
}