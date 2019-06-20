package lexer

//TokenType represents a token typing
type TokenType = int

const (
	IDBind  TokenType = iota
	IDBool  TokenType = iota
	IDFree  TokenType = iota
	ParenL  TokenType = iota
	ParenR  TokenType = iota
	NumInt  TokenType = iota
	OpFunc  TokenType = iota
	OpSum   TokenType = iota
	OpDiff  TokenType = iota
	OpProd  TokenType = iota
	OpDiv   TokenType = iota
	OpMod   TokenType = iota
	OpEqual TokenType = iota
	OpAnd   TokenType = iota
	OpOr    TokenType = iota
	OpXor   TokenType = iota
	OpNeg   TokenType = iota
)

type Token struct {
	typ TokenType
	str string
}

func evaluate(in []lexeme) ([]Token, error) {
	for _, l := range in {

	}
	return nil, nil
}

func mapOp(in lexeme) (Token, error) {
	tok := Token{str: in.str}
	switch in.str {
	case ".":
		tok.typ = OpFunc
	case "+":
		tok.typ = OpSum
	}
}

func mapNumLit(in lexeme) Token {
	return Token{
		typ: NumInt,
		str: in.str,
	}
}

func mapLSep(in lexeme) (out []Token) {
	for _, r := range in.str {
		out = append(out, Token{
			typ: in.ParenL,
			str: in.str,
		})
	}
	return
}

func mapRSep(in lexeme) (out []Token) {
	for _, r := range in.str {
		out = append(out, Token{
			typ: in.ParenR,
			str: in.str,
		})
	}
	return
}

func mapIdent(in lexeme) Token {
	tok := Token{str: in.str}
	switch in.str {
	case "let":
		tok.typ = IDBind
	case "true", "false":
		tok.typ = IDBool
	default:
		tok.typ = IDFree
	}
	return tok
}
