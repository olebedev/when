package rules

type MatchByIndex []*Match

func (m MatchByIndex) Len() int {
	return len(m)
}

func (m MatchByIndex) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MatchByIndex) Less(i, j int) bool {
	return m[i].Left < m[j].Left
}

type MatchByOrder []*Match

func (m MatchByOrder) Len() int {
	return len(m)
}

func (m MatchByOrder) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MatchByOrder) Less(i, j int) bool {
	return m[i].Order < m[j].Order
}
