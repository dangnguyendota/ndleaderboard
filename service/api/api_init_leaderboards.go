package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) InitLeaderboards(ctx context.Context, request *api.InitLeaderboardsRequest) (*api.Status, error) {

	return nil, status.Errorf(codes.Unimplemented, "method InitLeaderboards not implemented")
}
