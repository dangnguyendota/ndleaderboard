package leaderboard

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
)

func (l *LocalLeaderboards) GetRecordsList(ctx context.Context, leaderboardID string, recordsIDs []string) ([]*api.Record, error) {
	panic("implement me")
}