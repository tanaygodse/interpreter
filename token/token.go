package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	NOT         = "!"
	MULTIPLY    = "*"
	DIVIDE      = "/"
	EQ          = "=="
	NOT_EQ      = "!="
	PLUS_EQ     = "+="
	MINUS_EQ    = "-="
	MULTIPLY_EQ = "*="
	DIVIDE_EQ   = "/="
	EXPONENT    = "**"

	LT    = "<"
	GT    = ">"
	LT_EQ = "<="
	GT_EQ = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
