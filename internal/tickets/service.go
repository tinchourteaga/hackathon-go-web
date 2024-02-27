package tickets

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTotalTickets(ctx *gin.Context, destination string) (int, error)
	AverageDestination(ctx *gin.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx *gin.Context, destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, err
	}

	return len(tickets), nil
}

func (s *service) AverageDestination(ctx *gin.Context, destination string) (float64, error) {
	allTickets, err := s.repository.GetAll(ctx)

	if err != nil {
		return 0, err
	}

	dstTickets, err := s.repository.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, err
	}

	return float64(len(dstTickets)) / float64(len(allTickets)), nil
}
