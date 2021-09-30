package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) GetRank(context.Context, *api.GetRankRequest) (*api.GetRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRank not implemented")
}