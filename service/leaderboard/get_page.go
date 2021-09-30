package leaderboard

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
)

func (l *LocalLeaderboards) GetPage(ctx context.Context, leaderboardID string, pageNumber int, pageSize int) ([]*api.Record, error) {
	panic("implement me")
}