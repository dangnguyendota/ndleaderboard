package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) WriteLeaderboardRecord(context.Context, *api.WriteLeaderboardRecordRequest) (*api.WriteLeaderboardRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteLeaderboardRecord not implemented")
}