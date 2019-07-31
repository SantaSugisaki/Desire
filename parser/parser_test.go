package parser

import (
	"testing"
	"../ast"
	"../lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input) // 字句解析器を生成している
	p := New(l) // パーサを生成している

	program := p.ParseProgram() //プログラムをパースしている
	checkParserErrors(t,p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")	// プログラムが存在しなかった
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements)) // プログラムが全部で何行あったかを表示する
	}

	// あるべき識別子をtestsに登録している
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	// i:インデックス, tt:識別子
	for i,tt := range tests {
		stmt := program.Statements[i]	// stmtはstatementの略
		if !testLetStatement(t,stmt,tt.expectedIdentifier){ // let構文であるかをチェックしている
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input :=`
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",len(program.Statements))
	}

	for _,stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",returnStmt.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0{
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _,msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {	// let構文ではなかった場合
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false // let構文ではない
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T",s) // 何構文かを表示する
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s",name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",name,letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
