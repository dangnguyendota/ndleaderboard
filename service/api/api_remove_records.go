package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) RemoveRecords(context.Context, *api.RemoveRecordsRequest) (*api.RemoveRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRecords not implemented")
}