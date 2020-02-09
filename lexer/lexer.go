package lexer

import "github.com/ShunOkuma/interpreterGo/token"

// Lexer :字句解析器の情報を格納するための構造体
// input文字をchにbyte型として格納するためASCII文字のみ利用できる。
// 日本語などのUTF-8に対応させるためにはchをruneで設定する必要がある上に、l.readPositionの仕組みを使えないため、ここでは対象外としている。
type Lexer struct {
	input        string
	position     int  // 現在値
	readPosition int  // これから読み込む値
	ch           byte // 現在値のinput文字
}

// readChar :input文字列から1つ読み込んでchにセットし、readPositionを1つ進める
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// New :Lexerのポインタを返すNew関数
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // 一度呼んでおくことで position, readPosition, chの初期化になる
	return l
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// NextToken :
func (l *Lexer) NextToken() token.Token {
	var tok token.Token // Token型として設定

	l.skipWhitespace()

	// switchでchをチェックし、chに合ったconstとch自体をToken型として格納する
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
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
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
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
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default: // 識別子/キーワード
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // 早めにreturnしているのはreadIdentifier()でreadCharしているため
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

// readIdentifier :識別子やキーワードの文字列を1文字ずつ読み取ってpositionを使って取り出す
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) { // isLetterがtrueである限り、つまりa~z, A~Z, _である限り、loopしてreadCharし続ける
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter :chが a~z, A~Z, _ のどれかだったらその文字をそのまま返す、どれでもない場合は False を返す
// こうすることで簡単に何を英字とするかを設定できる。ここでは_も英字として扱っている。
func isLetter(ch byte) bool {
	return 'a' <= ch && ch >= 'z' || 'A' <= ch && ch >= 'Z' || ch == '_'
}

// skipWhitespace :空白、タブ、改行、復帰？が来たらスキップする
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// peekChar :次の文字の中身を確認する（positionは進めない、確認するだけ）
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) { //まずはreadPositionがinputを超えていないかチェック
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
