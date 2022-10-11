package tickets

import (
	"context"
	"desafio-go-web-leidyroldan/internal/domain"
)

type Service interface {
	GetTicketsByCountry(ctx context.Context, country string) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, country string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTicketsByCountry(ctx context.Context, country string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, country)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *service) GetTotalTickets(ctx context.Context) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tickets, nil

}

func (s *service) AverageDestination(ctx context.Context, country string) (float64, error) {
	totalTickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	ticketsByCountry, err := s.repository.GetTicketByDestination(ctx, country)
	if err != nil {
		return 0, err
	}

	average := len(ticketsByCountry) / len(totalTickets)

	return float64(average), nil

}
