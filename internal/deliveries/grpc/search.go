package deliveries

import (
	"context"
	"fmt"

	"github.com/perfectbuii/gen-go-template/internal/services"
	"github.com/perfectbuii/gen-go-template/pb"
	"github.com/perfectbuii/gen-go-template/transform"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type searchDelivery struct {
	searchService services.SearchService
	pb.UnimplementedSearchServiceServer
}

func NewSearchDelivery(searchService services.SearchService) pb.SearchServiceServer {
	return &searchDelivery{
		searchService: searchService,
	}
}

func (d *searchDelivery) GetTeamHub(ctx context.Context, req *pb.SearchTeamHubRequest) (*pb.SearchTeamHubResponse, error) {
	teams, hubs, err := d.searchService.TeamHub(ctx, req.GetQ())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("d.searchService.TeamHub: %v", err).Error())
	}

	return &pb.SearchTeamHubResponse{
		Teams: transform.TeamToPbPtrList(teams),
		Hubs:  transform.HubToPbPtrList(hubs),
	}, nil
}
