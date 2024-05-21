package engine

import (
	"Match-Dating/internal/model"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type RBTQueue struct {
	tree rbt.Tree
}

func NewRBTQueue(gender model.Gender) *RBTQueue {
	return &RBTQueue{
		tree: *rbt.NewWith(GetComparator(gender)),
	}
}

func GetComparator(gender model.Gender) func(a, b interface{}) int {
	if gender == model.GenderMale {
		return func(a, b interface{}) int {
			aP := a.(*model.Profile)
			bP := b.(*model.Profile)
			if aP == bP {
				return 0
			}
			if aP.Height <= bP.Height {
				return 1
			}

			return -1
		}
	} else if gender == model.GenderFemale {
		return func(a, b interface{}) int {
			aP := a.(*model.Profile)
			bP := b.(*model.Profile)
			if aP == bP {
				return 0
			}

			if aP.Height >= bP.Height {
				return 1
			}

			return -1
		}
	}
	return func(a, b interface{}) int { return 0 }

}

func (t *RBTQueue) AddProfile(p *model.Profile) {
	t.tree.Put(p, p)
}

func (t *RBTQueue) RemoveProfile(p *model.Profile) {
	t.tree.Remove(p)
}

func (t *RBTQueue) ListProfilesWithHeight(p *model.Profile) []*model.Profile {
	result := make([]*model.Profile, 0)

	it := t.tree.Iterator()
	for it.Next() {
		c := it.Value().(*model.Profile)
		if p.Gender == model.GenderMale {
			if p.Height <= c.Height {
				break
			}
			result = append(result, c)
		} else if p.Gender == model.GenderFemale {
			if p.Height >= c.Height {
				break
			}
			result = append(result, c)
		}
	}

	return result
}

func (t *RBTQueue) ListProfilesWithHeightAndQty(p *model.Profile, qty int) []*model.Profile {
	result := make([]*model.Profile, 0, qty)

	it := t.tree.Iterator()
	for it.Next() {
		if qty <= 0 {
			break
		}
		c := it.Value().(*model.Profile)
		if p.Gender == model.GenderMale {
			if p.Height <= c.Height {
				break
			}
			result = append(result, c)
		} else if p.Gender == model.GenderFemale {
			if p.Height >= c.Height {
				break
			}
			result = append(result, c)
		}
		qty--
	}
	return result

}

func (t *RBTQueue) Match(p *model.Profile) *model.Profile {
	// 隨機取一個且符合規定 且返回
	panic("not implement")
}

func (t *RBTQueue) GetProfile(p *model.Profile) *model.Profile {

	if val, found := t.tree.Get(p); found {
		return val.(*model.Profile)
	}
	return nil
}
