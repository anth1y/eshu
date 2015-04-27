package sql

type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOTA
	WS

	// Literals
	IDENT // fields, table_name

	// Misc characters
	ASTERISK // *
	COMMA    // ,

	// keywords
	SELECT
	FROM
)
