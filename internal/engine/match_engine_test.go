package engine

import (
	"Match-Dating/internal/model"
	"reflect"
	"testing"
)

func TestMatchEngine_RemoveProfile(t *testing.T) {
	type fields struct {
		profileMap map[uint]*model.Profile
		male       queueior
		female     queueior
	}
	type args struct {
		p *model.Profile
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &MatchEngine{
				profileMap: tt.fields.profileMap,
				male:       tt.fields.male,
				female:     tt.fields.female,
			}
			e.RemoveProfile(tt.args.p)
		})
	}
}

func TestMatchEngine_QueryResult(t *testing.T) {
	type fields struct {
		profileMap map[uint]*model.Profile
		male       queueior
		female     queueior
	}
	type args struct {
		c int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.QueryResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &MatchEngine{
				profileMap: tt.fields.profileMap,
				male:       tt.fields.male,
				female:     tt.fields.female,
			}
			if got := e.QueryResult(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchEngine.QueryResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchEngine_AddProfile(t *testing.T) {
	type fields struct {
		profileMap map[uint]*model.Profile
		male       queueior
		female     queueior
	}
	type args struct {
		p *model.Profile
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &MatchEngine{
				profileMap: tt.fields.profileMap,
				male:       tt.fields.male,
				female:     tt.fields.female,
			}
			e.AddProfile(tt.args.p)
		})
	}
}
