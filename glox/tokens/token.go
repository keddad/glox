package tokens

import (
	"glox/glox/runtime"
	"strconv"
	"unicode"
)

type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	IDENTIFIER
	STRING
	NUMBER
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Token struct {
	TokenType
	Lexeme  string
	InText  int
	Literal interface{}
}

func (a Token) Equal(b Token) bool {
	return a.TokenType == b.TokenType && a.Literal == b.Literal && a.InText == b.InText && a.Lexeme == b.Lexeme
}

func ScanTokens(in_text string, state *runtime.State) []Token {
	start := 0
	current := 0
	line := 1

	tokens := make([]Token, 0)

	isEnd := func() bool {
		return current >= len(in_text)
	}

	advance := func() byte {
		current += 1
		return in_text[current-1]
	}

	addToken := func(t TokenType, l interface{}) {
		text := in_text[start:current]
		tokens = append(tokens, Token{TokenType: t, Lexeme: text, InText: line, Literal: l})
	}

	match := func(expected byte) bool {
		if isEnd() {
			return false
		}

		if in_text[current] != expected {
			return false
		}

		current += 1
		return true
	}

	peek := func() byte {
		if isEnd() {
			return 0
		}

		return in_text[current]
	}

	peekNext := func() byte {
		if current+1 >= len(in_text) {
			return 0
		}

		return in_text[current+1]
	}

	isAlpha := func(in byte) bool {
		return unicode.IsLetter(rune(in)) || in == '_'
	}

	isAlphaNumeric := func(in byte) bool {
		return isAlpha(in) || unicode.IsDigit(rune(in))
	}

	scanToken := func() {
		c := advance()

		switch c {
		case '(':
			addToken(LEFT_PAREN, nil)
		case ')':
			addToken(RIGHT_PAREN, nil)
		case '{':
			addToken(LEFT_BRACE, nil)
		case '}':
			addToken(RIGHT_BRACE, nil)
		case ',':
			addToken(COMMA, nil)
		case '.':
			addToken(DOT, nil)
		case '-':
			addToken(MINUS, nil)
		case '+':
			addToken(PLUS, nil)
		case ';':
			addToken(SEMICOLON, nil)
		case '*':
			addToken(STAR, nil)
		case '!':
			if match('=') {
				addToken(BANG_EQUAL, nil)
			} else {
				addToken(BANG, nil)
			}
		case '=':
			if match('=') {
				addToken(EQUAL_EQUAL, nil)
			} else {
				addToken(EQUAL, nil)
			}
		case '<':
			if match('=') {
				addToken(LESS_EQUAL, nil)
			} else {
				addToken(LESS, nil)
			}
		case '>':
			if match('=') {
				addToken(GREATER_EQUAL, nil)
			} else {
				addToken(GREATER, nil)
			}
		case '/':
			if match('/') {
				for peek() != '\n' && current < len(in_text) {
					current += 1
				}
			} else {
				addToken(SLASH, nil)
			}

		case ' ':
		case '\r':
		case '\t':
		case '\n':
			line += 1
		case '"':
			for peek() != '"' && !isEnd() {
				if peek() == '\n' {
					line += 1
				}
				current += 1
			}

			if isEnd() {
				runtime.RaiseError(line, state)
				return
			}

			current += 1 // closing "
			addToken(STRING, in_text[start+1:current-1])

		default:
			if unicode.IsDigit(rune(c)) {
				for unicode.IsDigit(rune(peek())) {
					advance()
				}

				if peek() == '.' && unicode.IsDigit(rune(peekNext())) {
					advance()
				}

				for unicode.IsDigit(rune(peek())) {
					advance()
				}

				num, err := strconv.ParseFloat(string(in_text[start:current]), 64)

				if err != nil {
					runtime.RaiseError(line, state)
					return
				}

				addToken(NUMBER, num)

			} else if isAlpha(c) {
				for isAlphaNumeric(peek()) {
					current += 1
				}

				text := in_text[start:current]
				token, exists := keywords[string(text)]

				if !exists {
					token = IDENTIFIER
				}

				addToken(token, nil)
			} else {
				runtime.RaiseError(line, state)
			}
		}

	}

	for current < len(in_text) {
		start = current
		scanToken()
	}

	return tokens
}
