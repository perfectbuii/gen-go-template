package features

import (
	"context"
	"database/sql"

	"github.com/cucumber/godog"
	"github.com/perfectbuii/gen-go-template/internal/models"
	"google.golang.org/grpc"
)

type Suite struct {
	DB   *models.Queries
	Conn *grpc.ClientConn

	*StepState
}

type StepState struct {
	Response    interface{}
	ResponseErr error
	TeamID      string
	HubID       string
	HubName     string
}

func StepStateFromContext(ctx context.Context) *StepState {
	return ctx.Value(StepState{}).(*StepState)
}

func StepStateToContext(ctx context.Context, state *StepState) context.Context {
	return context.WithValue(ctx, StepState{}, state)
}

func InitializeScenario(sc *godog.ScenarioContext) {
	db, _ := sql.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	conn, _ := grpc.Dial("localhost:5000", grpc.WithInsecure())
	s := &Suite{
		DB:        models.New(db),
		Conn:      conn,
		StepState: &StepState{},
	}

	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		ctx = StepStateToContext(ctx, &StepState{})
		return ctx, nil
	})

	s.RegisterStep(sc)
}

func (s *Suite) GetSteps() map[string]interface{} {
	return map[string]interface{}{
		// example:
		// `^there are (\d+) godogs$`: s.thereAreGodogs,

		`^a signed in "([^"]*)"$`:                s.aSignedIn,
		`^a background$`:                         s.aBackground,
		`^user create hub$`:                      s.userCreateHub,
		`^returns "([^"]*)" status code$`:        s.returnsStatusCode,
		`^hub must be created$`:                  s.hubMustBeCreated,
		`^user search hub$`:                      s.userSearchHub,
		`^our system must return hub correctly$`: s.ourSystemMustReturnHubCorrectly,
		`^user create team$`:                     s.userCreateTeam,
		`^team must be created$`:                 s.teamMustBeCreated,

		/*generate_key*/
	}
}

func (s *Suite) RegisterStep(sc *godog.ScenarioContext) {
	steps := s.GetSteps()
	for step, fn := range steps {
		sc.Step(step, fn)
	}
}

func (s *Suite) aBackground(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (s *Suite) aSignedIn(ctx context.Context, arg1 string) (context.Context, error) {
	return ctx, nil
}

func (s *Suite) returnsStatusCode(ctx context.Context, arg1 string) (context.Context, error) {
	stepState := StepStateFromContext(ctx)
	if stepState.ResponseErr != nil {
		return StepStateToContext(ctx, stepState), stepState.ResponseErr
	}

	return StepStateToContext(ctx, stepState), nil
}
