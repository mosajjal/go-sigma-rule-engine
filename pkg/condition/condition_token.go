package condition

type Token int

const (
	TokErr Token = iota
	TokUnsupp

	// user-defined word
	Identifier
	IdentifierWithWildcard
	IdentifierAll

	// Literals
	LitEof

	// Separators
	SepLpar
	SepRpar
	SepPipe

	// Operators
	OpEq
	OpGt
	OpGte
	OpLt
	OpLte

	// Keywords
	KeywordAnd
	KeywordOr
	KeywordNot
	KeywordAgg

	// Statements
	StOne
	StAll
)

func (t Token) String() string {
	switch t {
	case Identifier, IdentifierWithWildcard:
		return "IDENT"
	case IdentifierAll:
		return "THEM"
	case SepLpar:
		return "LPAR"
	case SepRpar:
		return "RPAR"
	case SepPipe:
		return "PIPE"
	case OpEq:
		return "EQ"
	case OpGt:
		return "GT"
	case OpGte:
		return "GTE"
	case OpLt:
		return "LT"
	case OpLte:
		return "LTE"
	case KeywordAnd:
		return "AND"
	case KeywordOr:
		return "OR"
	case KeywordNot:
		return "NOT"
	case StAll:
		return "ALL"
	case StOne:
		return "ONE"
	case KeywordAgg:
		return "AGG"
	case LitEof:
		return "EOF"
	default:
		return "Err"
	}
}

func (t Token) Literal() string {
	switch t {
	case Identifier, IdentifierWithWildcard:
		return "keywords"
	case IdentifierAll:
		return "them"
	case SepLpar:
		return "("
	case SepRpar:
		return ")"
	case SepPipe:
		return "|"
	case OpEq:
		return "="
	case OpGt:
		return ">"
	case OpGte:
		return ">="
	case OpLt:
		return "<"
	case OpLte:
		return "<="
	case KeywordAnd:
		return "and"
	case KeywordOr:
		return "or"
	case KeywordNot:
		return "not"
	case StAll:
		return "all of"
	case StOne:
		return "1 of"
	case LitEof:
		return ""
	default:
		return "Err"
	}
}

func (t Token) Rune() rune {
	switch t {
	case SepLpar:
		return '('
	case SepRpar:
		return ')'
	case SepPipe:
		return '|'
	default:
		return eof
	}
}
