package gorse

import (
	"context"

	"github.com/brittonhayes/staffing"
	"github.com/pkg/errors"
	"github.com/zhenghaoz/gorse/client"
)

const (
	ErrGorseUserInsertFailed = "failed to insert user into gorse"
)

type recommendationRepository struct {
	gorse *client.GorseClient
}

func NewRecommendationRepository(connection string, apiKey string) staffing.RecommendationRepository {
	return &recommendationRepository{
		gorse: client.NewGorseClient(connection, apiKey),
	}
}

func (r *recommendationRepository) Close() error {
	return nil
}

func (r *recommendationRepository) CreateUser(ctx context.Context, user *staffing.User) error {
	_, err := r.gorse.InsertUser(ctx, client.User(*user))
	if err != nil {
		return errors.Wrap(err, ErrGorseUserInsertFailed)
	}

	return nil
}
