package leaderboard

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
)

func (l *LocalLeaderboards) WriteRecord(ctx context.Context, leaderboardID string, recordID string, score int) (*api.Record, error) {
	panic("implement me")
}