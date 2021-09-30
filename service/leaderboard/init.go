package leaderboard

import (
	"context"
	"fmt"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
)

func (l *LocalLeaderboards) Init(ctx context.Context, leaderboards []*api.Leaderboard) error {
	l.RWMutex.Lock()
	defer l.Unlock()
	l.cache = map[string]*api.Leaderboard{}
	for _, leaderboard := range leaderboards {
		if _, ok := l.cache[leaderboard.LeaderboardId]; ok {
			return NewError("Init", fmt.Errorf("duplicated leaderboard id %s", leaderboard.LeaderboardId))
		}
		l.cache[leaderboard.LeaderboardId] = leaderboard
	}
	return nil
}