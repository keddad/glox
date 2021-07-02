package tests

import (
	"glox/glox/runtime"
	"glox/glox/tokens"
	"reflect"
	"testing"
)

func TestSimple(t *testing.T) {
	str := "var x = 4;"
	exp_tokens := []tokens.Token{{TokenType: 36, Lexeme: "var", InText: 1, Literal: nil}, {TokenType: 19, Lexeme: "x", InText: 1, Literal: nil}, {TokenType: 13, Lexeme: "=", InText: 1, Literal: nil}, {TokenType: 21, Lexeme: "4", InText: 1, Literal: float64(4)}, {TokenType: 8, Lexeme: ";", InText: 1, Literal: nil}}

	state := runtime.NewState()
	gen_tokens := tokens.ScanTokens(str, &state)

	if state.HadError {
		t.Errorf("State had HadError set even though it shouldn't\n")
	}

	if !reflect.DeepEqual(gen_tokens, exp_tokens) {
		t.Errorf("Returned tokens didn't match expected ones\n")
		t.Errorf("Expected: %+v\n", exp_tokens)
		t.Errorf("Got: %+v\n", gen_tokens)
	}
}

func TestCorrectStr(t *testing.T) {
	str := "var x = \"meow\";"
	exp_tokens := []tokens.Token{{TokenType: 36, Lexeme: "var", InText: 1, Literal: nil}, {TokenType: 19, Lexeme: "x", InText: 1, Literal: nil}, {TokenType: 13, Lexeme: "=", InText: 1, Literal: nil}, {TokenType: 20, Lexeme: "meow", InText: 1, Literal: "meow"}, {TokenType: 8, Lexeme: ";", InText: 1, Literal: nil}}

	state := runtime.NewState()
	gen_tokens := tokens.ScanTokens(str, &state)

	if state.HadError {
		t.Errorf("State had HadError set even though it shouldn't\n")
	}

	if !reflect.DeepEqual(gen_tokens, exp_tokens) {
		t.Errorf("Returned tokens didn't match expected ones\n")
		t.Errorf("Expected: %+v\n", exp_tokens)
		t.Errorf("Got: %+v\n", gen_tokens)
	}
}

func TestIncorrectStr(t *testing.T) {
	str := "var x = \"meow"

	state := runtime.NewState()
	_ = tokens.ScanTokens(str, &state)

	if !state.HadError {
		t.Errorf("State hadn't HadError set even though it should\n")
	}
}
