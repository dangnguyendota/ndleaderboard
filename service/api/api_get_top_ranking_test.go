package api

import (
	"context"
	api "github.com/dangnguyendota/ndleaderboard/proto/ndleaderboard/api"
	"github.com/dangnguyendota/ndleaderboard/service/leaderboard"
	"reflect"
	"testing"
)

func TestLeaderboardService_GetTopRanking(t *testing.T) {
	type fields struct {
		UnimplementedNDLeaderboardServer api.UnimplementedNDLeaderboardServer
		leaderboards                     leaderboard.Leaderboards
	}
	type args struct {
		in0 context.Context
		in1 *api.GetTopRankingRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.GetTopRankingResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &LeaderboardService{
				UnimplementedNDLeaderboardServer: tt.fields.UnimplementedNDLeaderboardServer,
				leaderboards:                     tt.fields.leaderboards,
			}
			got, err := s.GetTopRanking(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopRanking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopRanking() got = %v, want %v", got, tt.want)
			}
		})
	}
}