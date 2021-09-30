package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) GetRecords(context.Context, *api.GetRecordsRequest) (*api.GetRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecords not implemented")
}