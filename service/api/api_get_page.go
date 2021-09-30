package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LeaderboardService) GetPage(context.Context, *api.GetPageRequest) (*api.GetPageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPage not implemented")
}