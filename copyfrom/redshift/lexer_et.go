package redshift

// ConsumeWord lets move position to consume given word
// SkipWhiteSpaces Skips white space characters in the input.
// PeekWord grab the next word (till whitespace, without consuming)
// Emit passes an token back to the client.
// Next returns the next rune in the input
// NextToken returns the next token from the input.

// // SkipAs find ad skip the "as" token eventually present af "format" and "delimiter"
// func SkipAs(l *lex.Lexer) lex.StateFn {
// 	l.SkipWhiteSpaces()
// 	if l.IsEnd() {
// 		return nil
// 	}
// 	word := strings.ToLower(l.PeekWord())
// 	//fmt.Printf("LexSkipAs looking for operator:  word=%q\n", word)
// 	if word == "as" {
// 		l.ConsumeWord(word)
// 		//l.Emit(lex.TokenAs)
// 		//l.Push("LexSkipAs", SkipAs)
// 	}
// 	// switch l.lastToken.T {
// 	// // case TokenCFormatAs:
// 	// // 	return LexEmpty
// 	// // case TokenCDelimiter:
// 	// // 	return LexValue
// 	// // case TokenCQuoteAs:
// 	// // 	return LexValue
// 	// default:
// 	// 	return LexEmpty
// 	// }
// 	return lex.LexEmpty
// }

// // GesQuotedString find a string with single quote
// func GesQuotedString(l *lex.Lexer) lex.StateFn {
// 	l.SkipWhiteSpaces()
// 	if l.IsEnd() {
// 		return nil
// 	}
// 	r := l.Peek()
// 	fmt.Println(r)
// 	if r != '\'' {
// 		return nil
// 	}
// 	//val := l.NextToken()
// 	l.Emit(lex.TokenNil)

// 	return lex.LexValue
// }
