package staffing

import (
	"context"
	"errors"

	"github.com/zhenghaoz/gorse/client"
)

var ErrRecommendationNotFound = errors.New("recommendation not found")

type RecommendationRepository interface {
	CreateUser(ctx context.Context, user *User) error
	Close() error
}

type User client.User
