package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) RemoveRecordsByRange(context.Context, *api.RemoveRecordsByRangeRequest) (*api.RemoveRecordsByRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRecordsByRange not implemented")
}