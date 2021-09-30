package leaderboard

import "context"

func (l *LocalLeaderboards) GetRank(ctx context.Context, leaderboardID string, recordID string) (int, error) {
	panic("implement me")
}