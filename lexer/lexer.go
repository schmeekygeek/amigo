package lexer

import (
  "amigo/token"
)

type Lexer struct {
  input        string
  position     int  // current position in input (points to current char index)
  readPosition int  // current reading position in input (+1 after current char)
  ch           byte // current char under examination
}

func New(input string) *Lexer {
  l := &Lexer{input: input}
  l.advance()
  return l
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
  return token.Token{
    Type:    tokenType,
    Literal: string(literal),
  }
}

func (l *Lexer) advance() {

  if l.readPosition >= len(l.input) {
    l.ch = 0
  } else {
    l.ch = l.input[l.readPosition]
  }

  l.position = l.readPosition
  l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {

  var tok token.Token

  l.skipWhitespace()

  switch l.ch {
  case '=':
    if l.peekChar() == '=' {
      ch := l.ch
      l.advance()
      tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
    } else {
      tok = newToken(token.ASSIGN, l.ch)
    }
  case '+':
    tok = newToken(token.PLUS, l.ch)
  case '-':
    tok = newToken(token.MINUS, l.ch)
  case '!':
    if l.peekChar() == '=' {
      ch := l.ch
      l.advance()
      tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
    } else {
      tok = newToken(token.BANG, l.ch)
    }
  case '/':
    tok = newToken(token.SLASH, l.ch)
  case '*':
    tok = newToken(token.ASTERISK, l.ch)
  case '<':
    tok = newToken(token.LT, l.ch)
  case '>':
    tok = newToken(token.GT, l.ch)
  case ',':
    tok = newToken(token.COMMA, l.ch)
  case ';':
    tok = newToken(token.SEMICOLON, l.ch)
  case '(':
    tok = newToken(token.LPAREN, l.ch)
  case ')':
    tok = newToken(token.RPAREN, l.ch)
  case '{':
    tok = newToken(token.LBRACE, l.ch)
  case '}':
    tok = newToken(token.RBRACE, l.ch)
  default:
    if isLetter(l.ch) {
      tok.Literal = l.readIdentifier()
      tok.Type = token.LookupIdentifier(tok.Literal)
      return tok
    } else if isDigit(l.ch) {
      tok.Literal = l.readNumber()
      tok.Type = token.INT
      return tok
    } else {
      tok = newToken(token.ILLEGAL, l.ch)
    }
  case 0:
    tok.Literal = ""
    tok.Type = token.EOF
}

  l.advance()

  return tok
}

func (l *Lexer) readIdentifier() string {

  position := l.position

  for isLetter(l.ch) {
    l.advance()
  }
  return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {

  position := l.position

  for isDigit(l.ch) {
    l.advance()
  }
  return l.input[position:l.position]
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
  return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
  for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
    l.advance()
  }
}

func (l *Lexer) peekChar() byte {
  if l.readPosition >= len(l.input) {
    return 0
  } else {
    return l.input[l.readPosition]
  }
}
