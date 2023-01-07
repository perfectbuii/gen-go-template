package features

import (
	"context"

	"github.com/cucumber/godog"
)

func (s *Suite) teamMustBeCreated(ctx context.Context) (context.Context, error) {
	stepState := StepStateFromContext(ctx)
	return StepStateToContext(ctx, stepState), godog.ErrPending
}

func (s *Suite) userCreateTeam(ctx context.Context) (context.Context, error) {
	stepState := StepStateFromContext(ctx)
	return StepStateToContext(ctx, stepState), godog.ErrPending
}