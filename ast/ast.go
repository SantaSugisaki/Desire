package ast

import (
	"../token"
	"bytes"
)

func (p *Program) String() string {
	var out bytes.Buffer // バッファを作成

	for _, s := range p.Statements { // プログラムをバッファに入れる
		out.WriteString(s.String())
	}

	return out.String() // プログラム全体を文字列として返す
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (i *Identifier) String() string { return i.Value }

// ノード型を定義している。識別子をフィールドに持つ
type Node interface {
	TokenLiteral() string
	String() string
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
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let構文
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier
	Value Expression
}

// return文
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression
}

// 式だけの文
type ExpressionStatement struct {
	Token      token.Token //式の最初のトークン
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal } // let構文の識別子を返すメソッド

// 識別子
type Identifier struct {
	Token token.Token // token.Ident トークン
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {return il.Token.Literal}
func (il *IntegerLiteral) String() string {return il.Token.Literal}
