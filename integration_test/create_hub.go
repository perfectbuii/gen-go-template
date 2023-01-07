package features

import (
	"context"
	"fmt"

	"github.com/perfectbuii/gen-go-template/pb"
	"github.com/perfectbuii/gen-go-template/utils"
)

func (s *Suite) hubMustBeCreated(ctx context.Context) (context.Context, error) {
	stepState := StepStateFromContext(ctx)

	stepState.Response, stepState.ResponseErr = pb.NewHubServiceClient(s.Conn).GetHubByID(ctx, &pb.GetHubByIDRequest{
		HubID: stepState.HubID,
	})

	resp := stepState.Response.(*pb.GetHubByIDResponse)
	if resp.Data == nil || resp.Data.Id != stepState.HubID {
		return StepStateToContext(ctx, stepState), fmt.Errorf("can not get hub")
	}

	return StepStateToContext(ctx, stepState), nil
}

func (s *Suite) userCreateHub(ctx context.Context) (context.Context, error) {
	stepState := StepStateFromContext(ctx)
	stepState.HubID = utils.RandStringBytes(10)
	stepState.HubName = utils.RandStringBytes(10)
	stepState.Response, stepState.ResponseErr = pb.NewHubServiceClient(s.Conn).CreateHub(ctx, &pb.CreateHubRequest{
		Hub: &pb.Hub{
			Id:         stepState.HubID,
			Name:       stepState.HubName,
			LocationId: "XGG",
		},
	})

	return StepStateToContext(ctx, stepState), nil
}
