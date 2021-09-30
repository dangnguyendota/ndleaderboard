package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) GetAroundMe(context.Context, *api.GetAroundMeRecordsRequest) (*api.GetAroundMeRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAroundMe not implemented")
}