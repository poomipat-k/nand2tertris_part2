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
	fmt.Println("type: ", e.tk.TokenType())
	fmt.Println("token: ", e.tk.Token())

	// no statements
	if !statementKeywords[e.tk.Keyword()] {
		return
	}

	// TODO: need a for loop
	e.WriteString("<statements>\n")

	if e.tk.Keyword() == "let" {
		e.CompileLet()
	} else if e.tk.Keyword() == "if" {
		e.CompileIf()
	} else if e.tk.Keyword() == "while" {
		e.CompileWhile()
	} else if e.tk.Keyword() == "do" {
		e.CompileDo()
	} else if e.tk.Keyword() == "return" {
		e.CompileReturn()
	} else {
		log.Fatal("CompileStatements, expect a statement keyword (let | if | while | do | return)")
	}

	e.WriteString("</statements>\n")

}

/* 'let' varName('['expression']')? '=' expression ';' */
func (e *Engine) CompileLet() {
	fmt.Println("--- CompileLet ---")

	e.WriteString("<letStatement>\n")
	// let
	e.writeKeyword()

	e.tk.Advance()
	// varName
	if e.tk.Identifier() == "" {
		log.Fatal("CompileLet, expect a varName(identifier), got:", e.tk.Token(), " ", e.tk.TokenType())
	}
	e.writeIdentifier()

	e.tk.Advance()
	// '['
	if e.tk.Symbol() == "[" {
		e.writeSymbol()

		e.tk.Advance()
		// expression
		e.CompileExpression()
		// end expression

		if e.tk.Symbol() != "]" {
			log.Fatal("CompileLet, expect a ]")
		}
		// ']'
		e.writeSymbol()
		e.tk.Advance()
	}

	if e.tk.Symbol() != "=" {
		log.Fatal("CompileLet, expect = or [")

	}

	// '='
	e.writeSymbol()

	e.tk.Advance()
	// expression
	e.CompileExpression()
	// end expression

	if e.tk.Symbol() != ";" {
		log.Fatal("CompileLet, expect a ';'")
	}
	// ';'
	e.writeSymbol()
	e.WriteString("</letStatement>\n")
}

/* 'if' '(' expression ')' '{' statements '}' ('else' '{' statements '}')? */
func (e *Engine) CompileIf() {
	fmt.Println("--- CompileIf ---")
	// if
	e.writeKeyword()

	e.tk.Advance()
	if e.tk.Symbol() != "(" {
		log.Fatal("CompileIf, expect a '('")
	}
	e.writeSymbol()

	e.tk.Advance()
	e.CompileExpression()

	if e.tk.Symbol() != ")" {
		log.Fatal("CompileIf, expect a ')'")
	}
	e.writeSymbol()

	e.tk.Advance()
	if e.tk.Symbol() != "{" {
		log.Fatal("CompileIf, expect a '{'")
	}
	e.writeSymbol()

	e.tk.Advance()
	e.CompileStatements()

	if e.tk.Symbol() != "}" {
		log.Fatal("CompileIf, expect a '}'")
	}
	e.writeSymbol()

	e.tk.Advance()
	// check if there is an else clause
	if e.tk.Keyword() != "else" {
		return
	}
	e.writeKeyword()

	e.tk.Advance()
	if e.tk.Symbol() != "{" {
		log.Fatal("CompileIf, expect a '{'")
	}
	e.writeSymbol()

	e.tk.Advance()
	e.CompileStatements()

	if e.tk.Symbol() != "}" {
		log.Fatal("CompileIf, expect a '}'")
	}
	e.writeSymbol()

}

/* 'while' '(' expression ')' '{' expressions '}' */
func (e *Engine) CompileWhile() {
	fmt.Println("--- CompileWhile ---")
	e.writeKeyword()
}

/* 'do' subroutineCall ';' */
func (e *Engine) CompileDo() {
	fmt.Println("--- CompileDo ---")
	e.writeKeyword()

	e.tk.Advance()
	e.CompileExpression()
}

/* 'return' expression? ';' */
func (e *Engine) CompileReturn() {
	fmt.Println("--- CompileReturn ---")
	e.writeKeyword()
}
