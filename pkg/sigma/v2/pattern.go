package sigma

import "strings"

type TextPatternModifier int

const (
	TextPatternContains TextPatternModifier = iota
	TextPatternPrefix
	TextPatternSuffix
)

// StringMatcher is an atomic pattern that could implement glob, literal or regex matchers
type StringMatcher interface {
	// StringMatch implements StringMatcher
	StringMatch(string) bool
}

// StringMatchers holds multiple atomic matchers
// Patterns are meant to be list of possibilities
// thus, objects are joined with logical disjunctions
type StringMatchers []StringMatcher

// StringMatch implements StringMatcher
func (s StringMatchers) StringMatch(msg string) bool {
	for _, m := range s {
		if m.StringMatch(msg) {
			return true
		}
	}
	return false
}

// ContentPattern is a token for literal content matching
type ContentPattern struct {
	Token     string
	Lowercase bool
}

// StringMatch implements StringMatcher
func (c ContentPattern) StringMatch(msg string) bool {
	return strings.Contains(msg, func() string {
		if c.Lowercase {
			return strings.ToLower(c.Token)
		}
		return c.Token
	}())
}

// PrefixPattern is a token for literal content matching
type PrefixPattern struct {
	Token     string
	Lowercase bool
}

// StringMatch implements StringMatcher
func (c PrefixPattern) StringMatch(msg string) bool {
	return strings.HasPrefix(msg, func() string {
		if c.Lowercase {
			return strings.ToLower(c.Token)
		}
		return c.Token
	}())
}

// SuffixPattern is a token for literal content matching
type SuffixPattern struct {
	Token     string
	Lowercase bool
}

// StringMatch implements StringMatcher
func (c SuffixPattern) StringMatch(msg string) bool {
	return strings.HasSuffix(msg, func() string {
		if c.Lowercase {
			return strings.ToLower(c.Token)
		}
		return c.Token
	}())
}
