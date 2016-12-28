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

type MatchByOrderAndIndex []*Match

func (m MatchByOrderAndIndex) Len() int {
	return len(m)
}

func (m MatchByOrderAndIndex) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MatchByOrderAndIndex) Less(i, j int) bool {
	if m[i].Order < m[j].Order {
		return true
	} else if m[i].Order > m[j].Order {
		return false
	}
	return m[i].Left < m[j].Left
}
