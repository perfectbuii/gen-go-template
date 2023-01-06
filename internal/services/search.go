package services

import (
	"context"

	"github.com/perfectbuii/gen-go-template/internal/models"
	"github.com/perfectbuii/gen-go-template/internal/repositories"
)

type SearchService interface {
	TeamHub(ctx context.Context, q string) ([]*models.Team, []*models.Hub, error)
}

type searchService struct {
	searchRepo repositories.SearchRepository
}

func NewSearchService(searchRepo repositories.SearchRepository) SearchService {
	return &searchService{
		searchRepo: searchRepo,
	}
}

func (s *searchService) TeamHub(ctx context.Context, q string) ([]*models.Team, []*models.Hub, error) {
	teams, users, err := s.searchRepo.TeamHub(ctx, q)
	if err != nil {
		return nil, nil, err
	}

	return teams, users, nil
}
