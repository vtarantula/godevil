package set

type Set struct {
	m_integer map[int]bool
}

func New() *Set {
	set := &Set{
		m_integer: make(map[int]bool),
	}
	return set
}

func (set *Set) Size() int {
	return len(set.m_integer)
}

func (set *Set) Contains(val int) bool {
	if _, ok := set.m_integer[val]; ok {
		return true
	}
	return false
}

func (set *Set) Add(val int) {
	set.m_integer[val] = true
}

func (set *Set) Delete(val int) {
	if set.Contains(val) {
		delete(set.m_integer, val)
	}
}

func (set *Set) Union(otherset *Set) *Set {
	unionset := New()
	for v := range set.m_integer {
		unionset.Add(v)
	}
	for v := range otherset.m_integer {
		unionset.Add(v)
	}
	return unionset
}

func (set *Set) Intersect(otherset *Set) *Set {
	intersectset := New()
	for v := range set.m_integer {
		if otherset.Contains(v) {
			intersectset.Add(v)
		}
	}
	return intersectset
}
