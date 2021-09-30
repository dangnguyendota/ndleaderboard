package leaderboard

import (
	"context"
	"fmt"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"github.com/dangnguyendota/ndleaderboard/service/redisdb"
	"sync"
)

type Error struct {
	f   string
	err error
}

func (e Error) Error() string {
	return fmt.Sprintf("error on function %s: %s", e.f, e.err)
}

func NewError(f string, err error) *Error {
	return &Error{
		f:   f,
		err: err,
	}
}

type Leaderboards interface {
	HealthCheck(ctx context.Context) error
	Init(ctx context.Context, leaderboards []*api.Leaderboard) error
	WriteRecord(ctx context.Context, leaderboardID string, recordID string, score int) (*api.Record, error)
	GetRecord(ctx context.Context, leaderboardID string, recordID string) (*api.Record, error)
	GetRecordsList(ctx context.Context, leaderboardID string, recordsIDs []string) ([]*api.Record, error)
	GetRank(ctx context.Context, leaderboardID string, recordID string) (int, error)
	GetPage(ctx context.Context, leaderboardID string, pageNumber int, pageSize int) ([]*api.Record, error)
	GetAroundMe(ctx context.Context, leaderboardID string, recordID string, pageSize int) ([]*api.Record, error)
	GetTopRanking(ctx context.Context, leaderboardID string, pageSize int) ([]*api.Record, error)
	RemoveLeaderboard(ctx context.Context, leaderboardID string) error
	RemoveRecord(ctx context.Context, leaderboardID string, recordID string) error
	RemoveRecords(ctx context.Context, leaderboardID string, recordsIDs []string) error
	RemoveRecordsByRange(ctx context.Context, leaderboardID string, fromRank int, toRank int) error
	TotalRecords(ctx context.Context, leaderboardID string) (int, error)
	UpdateRecord(ctx context.Context, leaderboardID string, recordID string, score int) error
}

type LocalLeaderboards struct {
	sync.RWMutex
	cache    map[string]*api.Leaderboard
	database *redisdb.RedisDB
}

func NewLeaderboards(database *redisdb.RedisDB) *LocalLeaderboards {
	return &LocalLeaderboards{
		cache:    map[string]*api.Leaderboard{},
		database: database,
	}
}

func (l *LocalLeaderboards) GetLeaderboardCache(leaderboardID string) (*api.Leaderboard, error) {
	l.RLock()
	defer l.RUnlock()
	if leaderboard, ok := l.cache[leaderboardID]; ok {
		return leaderboard, nil
	}
	return nil, fmt.Errorf("nil leaderboard")
}

func (l *LocalLeaderboards) RemoveLeaderboardCache(leaderboardID string) {
	l.Lock()
	defer l.Unlock()
	if _, ok := l.cache[leaderboardID]; ok {
		delete(l.cache, leaderboardID)
	}
}