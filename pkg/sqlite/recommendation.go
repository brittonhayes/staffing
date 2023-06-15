package sqlite

import (
	"context"

	"github.com/brittonhayes/staffing"
)

type recommendationRepository struct {
}

func NewRecommendationRepository(connection string, inmem bool) staffing.RecommendationRepository {
	return &recommendationRepository{}
}

func (r *recommendationRepository) Close() error {
	return nil
}

func (r *recommendationRepository) CreateUser(ctx context.Context, name string) error {
	return nil
}
