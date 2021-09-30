package leaderboard

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
)

func (l *LocalLeaderboards) GetRecord(ctx context.Context, leaderboardID string, recordID string) (*api.Record, error) {
	panic("implement me")
}