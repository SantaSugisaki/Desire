package parser

import (
	"../ast"
	"../lexer"
	"../token"
)

type Parser struct {
	// 字句解析器
	l *lexer.Lexer
	
	curToken token.Token // 現在のトークン
	peekToken token.Token // 次のトークン
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つのトークンを読み込む。curTokenとpeekTokenの両方がセットされる。
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

