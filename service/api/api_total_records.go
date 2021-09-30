package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) TotalRecords(context.Context, *api.TotalRecordsRequest) (*api.TotalRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalRecords not implemented")
}