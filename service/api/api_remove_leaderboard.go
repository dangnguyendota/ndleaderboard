package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) RemoveLeaderboard(context.Context, *api.RemoveLeaderboardRequest) (*api.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLeaderboard not implemented")
}