package repo

import (
	"context"

	"github.com/vingarcia/ddd-go-template/advanced/domain"
)

// User represents the operations we use for
// retrieving a user from a persistent storage
type Users interface {
	GetUser(ctx context.Context, userID int) (domain.User, error)
	UpsertUser(ctx context.Context, user domain.User) (userID int, err error)
}