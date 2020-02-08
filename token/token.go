package token

/*
わかりにくくなるので、コメントにて名詞を指す場合には日本語を使う（トークン、型など）
宣言した変数や構造体はそのまま使う（Token、Typeなど）
*/

// TokenType :トークン専用の型を用意
type TokenType string

// Token :型とトークン自体を格納できる構造体
type Token struct {
	Type    TokenType
	Literal string
}

// トークンを設定
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子＋リテラル
	IDENT = "IDENT" // add, hoge, x, y, ...
	INT   = "INT"   // 123456

	// 演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent :入力がkeywordsに該当するかを検索する関数
func LookupIdent(ident string) TokenType {
	// mapは2つ目の返り値として値が存在しているかどうかをTrue, Falseで返してくれる
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT // keywordsではなかったらIDENTを返す
}
