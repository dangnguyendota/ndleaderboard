package leaderboard

import "context"

func (l *LocalLeaderboards) HealthCheck(ctx context.Context) error {
	err := l.database.HealthCheck(ctx)
	if err != nil {
		return NewError("HealthCheck", err)
	}
	return nil
}