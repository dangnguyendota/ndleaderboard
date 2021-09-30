package leaderboard

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
)

func (l *LocalLeaderboards) GetTopRanking(ctx context.Context, leaderboardID string, pageSize int) ([]*api.Record, error) {
	panic("implement me")
}