package lexer

// Lexer :字句解析器の情報を格納するための構造体
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New :Lexerのポインタを返すNew関数
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
