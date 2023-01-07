package features

import (
	"context"
	"fmt"

	"github.com/perfectbuii/gen-go-template/pb"
)

func (s *Suite) userSearchHub(ctx context.Context) (context.Context, error) {
	stepState := StepStateFromContext(ctx)
	stepState.Response, stepState.ResponseErr = pb.NewSearchServiceClient(s.Conn).GetTeamHub(ctx, &pb.SearchTeamHubRequest{
		Q: stepState.HubName[:8],
	})

	return StepStateToContext(ctx, stepState), nil
}

func (s *Suite) ourSystemMustReturnHubCorrectly(ctx context.Context) (context.Context, error) {
	stepState := StepStateFromContext(ctx)
	resp := stepState.Response.(*pb.SearchTeamHubResponse)
	if resp == nil || resp.Hubs[0].Id != stepState.HubID {
		return StepStateToContext(ctx, stepState), fmt.Errorf("can not search hub")
	}

	return StepStateToContext(ctx, stepState), nil
}
