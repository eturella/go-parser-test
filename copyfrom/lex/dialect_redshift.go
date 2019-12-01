package lex

var (
	// RedshiftDialect is a dialect to parse the COPY FROM command
	RedshiftDialect *Dialect = &Dialect{
		Statements: []*Clause{
			//{Token: TokenUpdate, Clauses: SqlUpdateR},
			{Token: TokenCopy, Clauses: RedshiftCopy},
		},
	}

	// RedshiftCopy copy statement
	RedshiftCopy = []*Clause{
		{Token: TokenCopy, Lexer: LexIdentifierOfType(TokenTable)},
		{Token: TokenLeftParenthesis, Lexer: LexColumnNames, Optional: true},
		{Token: TokenFrom, Lexer: LexTableReferences, Clauses: sourceClauses, Name: "copyFrom.source"},
		//{Token: TokenCIamRole, Lexer: LexValue},
		{Token: TokenCRemovequotes, Lexer: LexEmpty, Optional: true},
		{Token: TokenCEmptyasnull, Lexer: LexEmpty, Optional: true},
		{Token: TokenCBlanksasnull, Lexer: LexEmpty, Optional: true},
		{Token: TokenCMaxerror, Lexer: LexValue, Optional: true},
		{Token: TokenCGzip, Lexer: LexEmpty, Optional: true},
		{Token: TokenCEscape, Lexer: LexEmpty, Optional: true},
		{Token: TokenCDelimiter, Lexer: LexSkipAs, Optional: true, Clauses: delimiterClauses},
		{Token: TokenCLzop, Lexer: LexEmpty, Optional: true},
		{Token: TokenCEscape, Lexer: LexEmpty, Optional: true},
		{Token: TokenCExplicitIds, Lexer: LexEmpty, Optional: true},
		{Token: TokenCJson, Lexer: LexValue, Optional: true},
		{Token: TokenCReadratio, Lexer: LexValue, Optional: true},
		// {Token: TokenCManifest, Lexer: LexEmpty, Optional: true},
		{Token: TokenCTimeformat, Lexer: LexValue, Optional: true},
		{Token: TokenCFixedwidth, Lexer: LexValue, Optional: true},
		{Token: TokenCCsv, Lexer: LexEmpty, Optional: true},
		{Token: TokenCQuoteAs, Lexer: LexEmpty, Optional: true},
		{Token: TokenAs, Lexer: LexValue, Optional: true},
		{Token: TokenCFormatAs, Lexer: LexSkipAs, Optional: true},
		{Token: TokenCFormatParquet, Lexer: LexEmpty, Optional: true},
		{Token: TokenCFormatAvro, Lexer: LexIdentityOrValue, Optional: true},
	}

	sourceClauses = []*Clause{
		{Token: TokenCIamRole, Lexer: LexEmpty},
		{Token: TokenValue, Lexer: LexValue},
		{Token: TokenCManifest, Lexer: LexEmpty, Optional: true},
		{Token: TokenCRegion, Lexer: LexEmpty, Optional: true},
		{Token: TokenAs, Lexer: LexValue, Optional: true},
		{Token: TokenValue, Lexer: LexValue, Optional: true},
		{Token: TokenCSSH, Lexer: LexEmpty, Optional: true},
	}

	delimiterClauses = []*Clause{
		{Token: TokenValue, Lexer: LexValue},
	}

	// // SqlUpdateR update statement
	// SqlUpdateR = []*Clause{
	// 	{Token: TokenUpdate, Lexer: LexIdentifierOfType(TokenTable)},
	// }
)

// NewRedshiftLexer creates a new lexer for the input string using RedshiftDialect
// this is sql(ish) compatible parser.
func NewRedshiftLexer(input string) *Lexer {
	return NewLexer(input, RedshiftDialect)
}
