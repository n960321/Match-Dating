package engine

import (
	"Match-Dating/internal/model"
	"sync"
)

var GetID = func() func() uint {
	var id uint
	return func() uint {
		id++
		return id
	}
}()

var mux sync.Mutex

type MatchEngine struct {
	profileMap map[uint]*model.Profile
	male       queueior
	female     queueior
}

type queueior interface {
	AddProfile(p *model.Profile)
	GetProfile(p *model.Profile) *model.Profile
	RemoveProfile(p *model.Profile)
	ListProfilesWithHeight(p *model.Profile) []*model.Profile
	Match(p *model.Profile) *model.Profile
}

func NewMatchEngine() *MatchEngine {
	return &MatchEngine{
		profileMap: make(map[uint]*model.Profile),
		male:       NewRBTQueue(model.GenderMale),
		female:     NewRBTQueue(model.GenderFemale),
	}
}

func (e *MatchEngine) AddProfile(p *model.Profile) {
	p.ID = GetID()
	e.profileMap[p.ID] = p

	if p.Gender == model.GenderMale {
		e.male.AddProfile(p)
	} else if p.Gender == model.GenderFemale {
		e.female.AddProfile(p)
	}
}

func (e *MatchEngine) RemoveProfile(p *model.Profile) {

	if _p, ok := e.profileMap[p.ID]; !ok {
		return
	} else {
		p = _p
	}

	if p.Gender == model.GenderMale {
		e.male.RemoveProfile(p)
	} else if p.Gender == model.GenderFemale {
		e.female.RemoveProfile(p)
	}

	delete(e.profileMap, p.ID)
}

func (e *MatchEngine) GetProfile(p *model.Profile) *model.Profile {
	if val, ok := e.profileMap[p.ID]; ok {
		if val.Gender == model.GenderFemale {
			val = e.female.GetProfile(val)
		} else if val.Gender == model.GenderMale {
			val = e.male.GetProfile(val)
		}

		if val == nil {
			delete(e.profileMap, p.ID)
		} else {
			return val
		}
	}
	return nil
}

func (e *MatchEngine) GetAllMatch(p *model.Profile) *model.MatchResult {
	result := new(model.MatchResult)
	if p.Gender == model.GenderMale {
		result.Candidates = e.female.ListProfilesWithHeight(p)
	} else if p.Gender == model.GenderFemale {
		result.Candidates = e.male.ListProfilesWithHeight(p)
	}

	return result
}

func (e *MatchEngine) QueryResult(c int) *model.QueryResult {
	result := new(model.QueryResult)

	for _, p := range e.profileMap {
		if c <= 0 {
			break
		}
		// profile := p
		pair := new(model.DatingInfo)
		if p.Gender == model.GenderMale {
			candidate := e.female.Match(p)
			if candidate == nil {
				continue
			}
			pair.Male = p
			pair.Female = candidate
		} else if p.Gender == model.GenderFemale {
			candidate := e.male.Match(p)
			if candidate == nil {
				continue
			}
			pair.Female = p
			pair.Male = candidate
		}
		pair.Male.WantedDates--
		pair.Female.WantedDates--

		if pair.Male.WantedDates <= 0 {
			e.male.RemoveProfile(pair.Male)
		}

		if pair.Female.WantedDates <= 0 {
			e.female.RemoveProfile(pair.Female)
		}

		result.Infos = append(result.Infos, pair)
		c--
	}

	return result
}
