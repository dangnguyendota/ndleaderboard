package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
)

func (s *LeaderboardService) HealthCheck(context.Context, *api.HealthCheckRequest) (*api.HealthCheckResponse, error) {
	return &api.HealthCheckResponse{
		Status: "leaderboard server using redis, developed by Nguyễn Đăng Nguyên",
	}, nil
}