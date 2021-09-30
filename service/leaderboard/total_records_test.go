package leaderboard

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"github.com/dangnguyendota/ndleaderboard/service/redisdb"
	"sync"
	"testing"
)

func TestLocalLeaderboards_TotalRecords(t *testing.T) {
	type fields struct {
		RWMutex  sync.RWMutex
		cache    map[string]*api.Leaderboard
		database *redisdb.RedisDB
	}
	type args struct {
		ctx           context.Context
		leaderboardID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
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
			got, err := l.TotalRecords(tt.args.ctx, tt.args.leaderboardID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TotalRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TotalRecords() got = %v, want %v", got, tt.want)
			}
		})
	}
}
