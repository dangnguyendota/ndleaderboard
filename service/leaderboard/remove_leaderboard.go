package leaderboard

import (
	"context"
)

func (l *LocalLeaderboards) RemoveLeaderboard(ctx context.Context, leaderboardID string) error {
	_, err := l.GetLeaderboardCache(leaderboardID)
	if err != nil {
		return NewError("RemoveLeaderboard", err)
	}
	if err := l.database.RemoveLeaderboard(ctx, leaderboardID); err != nil {
		return NewError("RemoveLeaderboard", err)
	}
	return nil
}