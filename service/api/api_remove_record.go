package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) RemoveRecord(context.Context, *api.RemoveRecordRequest) (*api.RemoveRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRecord not implemented")
}