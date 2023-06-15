package staffing

import (
	"context"
	"errors"

	"github.com/uptrace/bun"
)

var ErrRecommendationNotFound = errors.New("recommendation not found")

type RecommendationRepository interface {
	CreateUser(ctx context.Context, name string) error
	Close() error
}

type RecommendationID string

type Recommendation struct {
	bun.BaseModel `bun:"table:recommendations"`

	ID   RecommendationID `json:"id" bun:",pk"`
	Name string
}
