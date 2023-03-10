package repositories

import (
	"context"

	"github.com/perfectbuii/gen-go-template/internal/models"
)

type TeamRepository interface {
	Create(ctx context.Context, team *models.Team) error
	GetByID(ctx context.Context, id string) (*models.Team, error)
	GetList(ctx context.Context, offset, limit int) ([]*models.Team, error)
	Update(ctx context.Context, id string, team *models.Team) error
	Delete(ctx context.Context, id string) error
}

type teamRepository struct {
	db *models.Queries
}

func NewTeamRepository(db models.DBTX) TeamRepository {
	return &teamRepository{
		db: models.New(db),
	}
}

func (u *teamRepository) Create(ctx context.Context, team *models.Team) error {
	if _, err := u.db.CreateTeam(ctx, models.CreateTeamParams{
		ID:         team.ID,
		Name:       team.Name,
		Type:       team.Type,
		LocationID: team.LocationID,
		HubID:      team.HubID,
	}); err != nil {
		return err
	}

	return nil
}

func (u *teamRepository) Update(ctx context.Context, id string, team *models.Team) error {
	return nil
}

func (u *teamRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (u *teamRepository) GetList(ctx context.Context, offset, limit int) ([]*models.Team, error) {
	return nil, nil
}

func (u *teamRepository) GetByID(ctx context.Context, id string) (*models.Team, error) {
	return nil, nil
}
