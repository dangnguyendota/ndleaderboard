package leaderboard

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
)

func (l *LocalLeaderboards) GetAroundMe(ctx context.Context, leaderboardID string, recordID string, pageSize int) ([]*api.Record, error) {
	panic("implement me")
}