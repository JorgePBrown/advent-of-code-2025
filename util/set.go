package util

type HashSet[S comparable] map[S]struct{}

func (hs HashSet[S]) Has(s S) bool {
	_, ok := hs[s]
	return ok
}
func (hs HashSet[S]) Add(s S) {
	hs[s] = struct{}{}
}
