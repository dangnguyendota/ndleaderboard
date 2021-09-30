package leaderboard

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"github.com/dangnguyendota/ndleaderboard/service/redisdb"
	"reflect"
	"sync"
	"testing"
)

func TestLocalLeaderboards_WriteRecord(t *testing.T) {
	type fields struct {
		RWMutex  sync.RWMutex
		cache    map[string]*api.Leaderboard
		database *redisdb.RedisDB
	}
	type args struct {
		ctx           context.Context
		leaderboardID string
		recordID      string
		score         int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.Record
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
			got, err := l.WriteRecord(tt.args.ctx, tt.args.leaderboardID, tt.args.recordID, tt.args.score)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WriteRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}
