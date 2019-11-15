// Package condition provides library to parse sigma rule condition field into rule tokens
package condition

import (
	"fmt"

	"github.com/markuskont/go-sigma-rule-engine/pkg/match"
	"github.com/markuskont/go-sigma-rule-engine/pkg/types"
)

type Item struct {
	T   Token
	Val string
}

// TODO - perhaps we should invoke parse only if we actually need to parse the query statement and simply instantiate a single-branch rule otherwise
func Parse(s types.Detection) (*match.Tree, error) {
	if s == nil {
		return nil, types.ErrMissingDetection{}
	}
	switch len(s) {
	case 0:
		return nil, types.ErrMissingDetection{}
	case 1:
		// Simple case - should have only one search field, but should not have a condition field
		if c, ok := s["condition"].(string); ok {
			return nil, types.ErrIncompleteDetection{Condition: c}
		}
	case 2:
		// Simple case - one condition statement comprised of single IDENT that matches the second field name
		if c, ok := s["condition"].(string); !ok {
			return nil, types.ErrIncompleteDetection{Condition: "MISSING"}
		} else {
			if _, ok := s[c]; !ok {
				return nil, types.ErrIncompleteDetection{
					Condition: c,
					Msg:       fmt.Sprintf("Field %s defined in condition missing from map.", c),
					Keys:      s.FieldSlice(),
				}
			}
		}
		delete(s, "condition")
	default:
		// Complex case, time to build syntax tree out of condition statement
		raw, ok := s["condition"].(string)
		if !ok {
			return nil, types.ErrMissingCondition{}
		}
		delete(s, "condition")
		p := &parser{
			lex:       lex(raw),
			sigma:     s,
			tokens:    make([]Item, 0),
			previous:  TokBegin,
			condition: raw,
		}
		if err := p.run(); err != nil {
			return nil, err
		}
		return nil, nil
	}
	// Should only have one element as complex scenario is handled separately
	rx := s.Fields()
	ast := &match.Tree{}
	root, err := newRuleMatcherFromIdent(<-rx, false)
	if err != nil {
		return nil, err
	}
	ast.Root = root
	return ast, nil
}
