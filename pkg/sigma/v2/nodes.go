package sigma

// NodeSimpleAnd is a list of matchers connected with logical conjunction
type NodeSimpleAnd []Branch

// Match implements Matcher
func (n NodeSimpleAnd) Match(e Event) bool {
	for _, b := range n {
		if !b.Match(e) {
			return false
		}
	}
	return true
}

// Reduce cleans up unneeded slices
// Static structures can be used if node only holds one or two elements
// Avoids pointless runtime loops
func (n NodeSimpleAnd) Reduce() Branch {
	if len(n) == 1 {
		return n[0]
	}
	if len(n) == 2 {
		return &NodeAnd{L: n[0], R: n[1]}
	}
	return n
}

// NodeSimpleAnd is a list of matchers connected with logical disjunction
type NodeSimpleOr []Branch

// Reduce cleans up unneeded slices
// Static structures can be used if node only holds one or two elements
// Avoids pointless runtime loops
func (n NodeSimpleOr) Reduce() Branch {
	if len(n) == 1 {
		return n[0]
	}
	if len(n) == 2 {
		return &NodeOr{L: n[0], R: n[1]}
	}
	return n
}

// Match implements Matcher
func (n NodeSimpleOr) Match(e Event) bool {
	for _, b := range n {
		if b.Match(e) {
			return true
		}
	}
	return false
}

// NodeNot negates a branch
type NodeNot struct {
	B Branch
}

// Match implements Matcher
func (n NodeNot) Match(e Event) bool {
	return !n.B.Match(e)
}

// NodeAnd is a two element node of a binary tree with Left and Right branches
// connected via logical conjunction
type NodeAnd struct {
	L, R Branch
}

// Match implements Matcher
func (n NodeAnd) Match(e Event) bool {
	return n.L.Match(e) && n.R.Match(e)
}

// NodeAnd is a two element node of a binary tree with Left and Right branches
// connected via logical disjunction
type NodeOr struct {
	L, R Branch
}

// Match implements Matcher
func (n NodeOr) Match(e Event) bool {
	return n.L.Match(e) || n.R.Match(e)
}

func newNodeNotIfNegated(b Branch, negated bool) (Branch, bool) {
	if negated {
		return &NodeNot{B: b}, false
	}
	return b, false
}
