package leaderboard

import "context"

func (l *LocalLeaderboards) RemoveRecordsByRange(ctx context.Context, leaderboardID string, fromRank int, toRank int) error {
	panic("implement me")
}