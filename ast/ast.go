package ast

import (
	"../token"
)

// ノード型を定義している。識別子をフィールドに持つ
type Node interface {
	TokenLiteral() string
}

// 構文として意味を持つノードの型を定義している。
type Statement interface {
	Node
	statementNode()
}

// 式(値)
type Expression interface {
	Node
	expressionNode()
}

// プログラム全体
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len (p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let文
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal}

// 識別子
type Identifier struct {
	Token token.Token // token.Ident トークン
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal}
