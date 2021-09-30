package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) UpdateRecordScore(context.Context, *api.UpdateRecordScoreRequest) (*api.UpdateRecordScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecordScore not implemented")
}
