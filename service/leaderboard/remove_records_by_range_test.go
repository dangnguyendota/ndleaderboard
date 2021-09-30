package leaderboard

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"github.com/dangnguyendota/ndleaderboard/service/redisdb"
	"sync"
	"testing"
)

func TestLocalLeaderboards_RemoveRecordsByRange(t *testing.T) {
	type fields struct {
		RWMutex  sync.RWMutex
		cache    map[string]*api.Leaderboard
		database *redisdb.RedisDB
	}
	type args struct {
		ctx           context.Context
		leaderboardID string
		fromRank      int
		toRank        int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LocalLeaderboards{
				RWMutex:  tt.fields.RWMutex,
				cache:    tt.fields.cache,
				database: tt.fields.database,
			}
			if err := l.RemoveRecordsByRange(tt.args.ctx, tt.args.leaderboardID, tt.args.fromRank, tt.args.toRank); (err != nil) != tt.wantErr {
				t.Errorf("RemoveRecordsByRange() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
