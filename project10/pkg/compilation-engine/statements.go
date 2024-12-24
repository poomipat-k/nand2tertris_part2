package compilationEngine

import (
	"fmt"
	"log"
)

/*
statements: statement*
statement: letStatement | ifStatement | whileStatement | doStatement | returnStatement
*/
func (e *Engine) CompileStatements() {
	fmt.Println("--- CompileStatements ---")

	// no statements
	if !statementKeywords[e.tk.Keyword()] {
		return
	}

	keyword := e.tk.Keyword()
	// TODO: need a for loop
	if keyword == "let" {
		e.CompileLet()
	} else if keyword == "if" {
		e.CompileIf()
	} else if keyword == "while" {
		e.CompileWhile()
	} else if keyword == "do" {
		e.CompileDo()
	} else if keyword == "return" {
		e.CompileReturn()
	} else {
		log.Fatal("CompileStatements, expect a statement keyword (let | if | while | do | return)")
	}
}

/** 'let' varName('['expression']')? '=' expression ';' */
func (e *Engine) CompileLet() {
	fmt.Println("--- CompileLet ---")
	// let
	e.writeKeyword()

	e.tk.Advance()
	// varName
	if e.tk.Identifier() == "" {
		log.Fatal("CompileLet, expect a varName(identifier)")
	}
	e.writeIdentifier()

	e.tk.Advance()
	// '[' or '='
	if e.tk.Symbol() == "=" {

	} else if e.tk.Symbol() == "[" {

	} else {
		log.Fatal("CompileLet, expect = or [")
	}

}

/** 'if' '(' expression ')' '{' statements '}' ('else' '{' statements '}')? */
func (e *Engine) CompileIf() {
	fmt.Println("--- CompileIf ---")
	e.writeKeyword()
}

/** 'while' '(' expression ')' '{' expressions '}' */
func (e *Engine) CompileWhile() {
	fmt.Println("--- CompileWhile ---")
	e.writeKeyword()
}

/** 'do' subroutineCall ';' */
func (e *Engine) CompileDo() {
	fmt.Println("--- CompileDo ---")
	e.writeKeyword()
}

/** 'return' expression? ';' */
func (e *Engine) CompileReturn() {
	fmt.Println("--- CompileReturn ---")
	e.writeKeyword()
}
