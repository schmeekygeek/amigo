package token

type TokenType string

type Token struct {
  Type    TokenType
  Literal string
}

var keywords = map[string]TokenType {
  "fn":     FUNCTION,
  "let":    LET,
}

const (
  ILLEGAL = "ILLEGAL"
  EOF     = "EOF"

  // Identifiers + literals
  IDENT = "IDENT" // add, foobar
  INT   = "INT" // 123141

  ASSIGN = "="
  PLUS   = "+"

  COMMA     = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  FUNCTION = "FUNCTION"
  LET      = "LET"
)

func LookupIdentifier(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}
