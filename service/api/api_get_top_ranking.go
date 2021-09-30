package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) GetTopRanking(context.Context, *api.GetTopRankingRequest) (*api.GetTopRankingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopRanking not implemented")
}