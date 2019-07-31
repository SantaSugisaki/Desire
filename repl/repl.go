package repl

import (
	"bufio"
	"fmt"
	"io"
	"../lexer"
	"../token"
)

const PROMPT = ">> "

// 1行ずつプログラムを読み取ってトークンの種類を出力する関数
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for{
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()	// １行だけ読みこんでいる
		if !scanned {
			return
		}

		line := scanner.Text()	// 読み取った内容はscanner.Text()
		l := lexer.New(line) // 字句解析器を生成している

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
