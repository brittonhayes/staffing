package staffing

import (
	"context"
	"errors"
)

var ErrRecommendationNotFound = errors.New("recommendation not found")

type RecommendationRepository interface {
	CreateUser(ctx context.Context, name string) error
	Close() error
}

type RecommendationID string

type Recommendation struct {
	ID   RecommendationID
	Name string
}
