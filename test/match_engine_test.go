package engine

import (
	"Match-Dating/internal/engine"
	"Match-Dating/internal/model"
	"errors"
	"math/rand"
	"strconv"
	"testing"
)

// 測試AddProfile 正確性
func TestMatchEngine_AddProfile(t *testing.T) {
	type args struct {
		p      *model.Profile
		engine *engine.MatchEngine
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "加入一個profile，並查看是否存在於engine",
			args: args{
				p:      createProfile(0, model.GenderUnknow, 100, 200, 3, 10),
				engine: engine.NewMatchEngine(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.engine.AddProfile(tt.args.p)
			if r := tt.args.engine.GetProfile(tt.args.p); r != nil {
				t.Errorf("Add Profile failed ,do not exist in engine")
			}
		})
	}
}

// 測試GetAllMatch的正確性
func TestMatchEngine_GetAllMatch(t *testing.T) {
	type args struct {
		p          *model.Profile
		engine     *engine.MatchEngine
		candidates []*model.Profile
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "隨機建立profile，加進搓合引擎並檢查函示返回值正確與否",
			args: args{
				engine:     engine.NewMatchEngine(),
				p:          createProfile(0, model.GenderUnknow, 100, 200, 3, 10),
				candidates: createProfiles(int(RandomInt(1, 10000)), model.GenderUnknow, 100, 200, 3, 10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.engine.AddProfile(tt.args.p)
			for _, c := range tt.args.candidates {
				tt.args.engine.AddProfile(c)
			}

			result := tt.args.engine.GetAllMatch(tt.args.p)
			// t.Logf("p: %+v", tt.args.p)
			for _, c := range result.Candidates {
				// t.Logf("c: %+v", c)
				if err := check(tt.args.p, c); err != nil {
					t.Errorf("Check get failed a: %+v, b:%+v", tt.args.p, c)
				}
			}
		})
	}
}

// 測試刪除Profile的正確性
func TestMatchEngine_RemoveProfile(t *testing.T) {

	type args struct {
		engine     *engine.MatchEngine
		candidates []*model.Profile
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "建立多個Profile，然後刪除一個後，檢查engine中還有無存在",
			args: args{
				engine:     engine.NewMatchEngine(),
				candidates: createProfiles(int(RandomInt(1, 10000)), model.GenderUnknow, 100, 200, 3, 10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, c := range tt.args.candidates {
				tt.args.engine.AddProfile(c)
			}
			p := tt.args.candidates[0]
			tt.args.engine.RemoveProfile(p)
			if r := tt.args.engine.GetProfile(p); r != nil {
				t.Errorf("Remove Profile failed , still exist in engine")
			}
		})
	}
}

// 測試QueryResult 的正確性
func TODO_TestMatchEngine_QueryResult(t *testing.T) {
	type args struct {
		engine     *engine.MatchEngine
		candidates []*model.Profile
		count      int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "建立多個Profile，呼叫QueryResult，並檢查結果正確性",
			args: args{
				engine:     engine.NewMatchEngine(),
				candidates: createProfiles(int(RandomInt(1, 10000)), model.GenderUnknow, 100, 200, 3, 10),
				count:      int(RandomInt(10, 100)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, c := range tt.args.candidates {
				tt.args.engine.AddProfile(c)
			}

			result := tt.args.engine.QueryResult(tt.args.count)

			if len(result.Infos) > tt.args.count {
				t.Fatalf("The result is more then count")
			}
			for _, info := range result.Infos {
				if err := check(info.Male, info.Female); err != nil {
					t.Errorf("Query Match failed")
				}
			}
		})
	}
}

func createProfile(seed int, gender model.Gender, heightMin, heightMax, wantDatesMin, wantDatesMax int) *model.Profile {
	g := gender
	if g == model.GenderUnknow {
		g = model.Gender(RandomInt(1, 2))
	}

	return &model.Profile{
		Name:        "test_" + strconv.Itoa(seed),
		Height:      RandomInt(heightMin, heightMax),
		WantedDates: RandomInt(wantDatesMin, wantDatesMax),
		Gender:      g,
	}
}

func createProfiles(q int, gender model.Gender, heightMin, heightMax, wantDatesMin, wantDatesMax int) []*model.Profile {
	profiles := make([]*model.Profile, 0, q)
	for i := 0; i < q; i++ {
		profiles = append(profiles, createProfile(i, gender, heightMin, heightMax, wantDatesMin, wantDatesMax))
	}
	return profiles
}

func RandomInt(min, max int) uint {
	return uint(min + rand.Intn(max-min+1))
}

func check(a, b *model.Profile) error {
	if a.Gender == model.GenderFemale {
		a, b = b, a
	}

	// a male, b female

	if a.Height <= b.Height {
		return errors.New("Male can not short then female.")
	}
	return nil
}
