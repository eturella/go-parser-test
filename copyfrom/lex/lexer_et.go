package lex

import (
	"strings"
)

// ConsumeWord lets move position to consume given word
// SkipWhiteSpaces Skips white space characters in the input.
// PeekWord grab the next word (till whitespace, without consuming)
// Emit passes an token back to the client.
// Next returns the next rune in the input
// NextToken returns the next token from the input.

// LexSkipAs find ad skip the "as" token eventually present af "format" and "delimiter"
func LexSkipAs(l *Lexer) StateFn {
	l.SkipWhiteSpaces()
	if l.IsEnd() {
		return nil
	}
	word := strings.ToLower(l.PeekWord())
	//fmt.Printf("LexSkipAs looking for operator:  word=%q\n", word)
	if word == "as" {
		l.ConsumeWord(word)
		l.Emit(TokenAs)
		l.Push("LexSkipAs", LexSkipAs)
	}
	switch l.lastToken.T {
	// case TokenCFormatAs:
	// 	return LexEmpty
	// case TokenCDelimiter:
	// 	return LexValue
	// case TokenCQuoteAs:
	// 	return LexValue
	default:
		return LexEmpty
	}
}
