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
	l := lexer.New(input)
	p := New(l)

	program := ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i,tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t,stmt,tt.expectedIdentifier){
			return false
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s not *ast.LetStatement. got=%T",s)
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	
}
