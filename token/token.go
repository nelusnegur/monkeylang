package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Illegal = "ILLEGAL"
	Eof     = "EOF"

	Identifier = "IDENT"
	Int        = "INT"

	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Bang     = "!"
	Asterisk = "*"
	Slash    = "/"

	LessThan    = "<"
	GreaterThan = ">"
	Equal       = "=="
	NotEqual    = "!="

	Comma     = ","
	Semicolon = ";"

	LeftParen  = "("
	RightParen = ")"
	LeftBrace  = "{"
	RightBrace = "}"

	Function = "FUNCTION"
	Let      = "LET"
	True     = "TRUE"
	False    = "FALSE"
	If       = "IF"
	Else     = "ELSE"
	Return   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return Identifier
}
