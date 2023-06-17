package recommend

import (
	"context"

	"github.com/brittonhayes/staffing"
	"github.com/brittonhayes/staffing/proto/pb"
	"github.com/pkg/errors"
)

var _ Service = (*service)(nil)

var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides recommendation methods.
type Service interface {
	CreateUser(ctx context.Context, command *pb.RecommendationCreateUserCommand) error
}

type service struct {
	recommendations staffing.RecommendationRepository
}

func (s *service) CreateUser(ctx context.Context, command *pb.RecommendationCreateUserCommand) error {

	if err := command.ValidateAll(); err != nil {
		return err
	}

	err := s.recommendations.CreateUser(ctx, command.GetUserId())
	if err != nil {
		return errors.Wrap(err, "service failed to create user")
	}

	return nil
}

// NewService creates a new project service with the necessary dependencies.
func NewService(recommendations staffing.RecommendationRepository) Service {
	return &service{
		recommendations: recommendations,
	}
}
