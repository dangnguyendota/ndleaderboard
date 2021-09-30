package api

import (
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"github.com/dangnguyendota/ndleaderboard/service/leaderboard"
)

const (
	StatusDuplicatedLeaderboard uint32 = iota + 100
)

type LeaderboardService struct {
	api.UnimplementedNDLeaderboardServer
	leaderboards leaderboard.Leaderboards
}

func NewLeaderboardService(leaderboards leaderboard.Leaderboards) *LeaderboardService {
	return &LeaderboardService{
		leaderboards: leaderboards,
	}
}

