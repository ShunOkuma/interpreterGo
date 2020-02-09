package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/ShunOkuma/interpreterGo/lexer"
	"github.com/ShunOkuma/interpreterGo/token"
)

// PROMPT :行の先頭に表示させる
const PROMPT = ">> "

// Start :
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan() // scanできるかどうかをboolで返してくれる
		if !scanned {
			return
		}

		line := scanner.Text() // 一行読み出して
		l := lexer.New(line)   // 字句解析器にセット

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
