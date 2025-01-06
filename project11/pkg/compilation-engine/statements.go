package compilationEngine

import (
	"fmt"
	"log"

	vmWriter "github.com/poomipat-k/nand2tetris/project11/pkg/vm-writer"
)

/*
statements: statement*
statement: letStatement | ifStatement | whileStatement | doStatement | returnStatement
*/
func (e *Engine) CompileStatements() {
	fmt.Println("--- CompileStatements ---")

	// no statements
	if !statementKeywords[e.tk.Keyword()] {
		// e.WriteString("<statements>\n")
		// e.WriteString("</statements>\n")
		return
	}

	// e.WriteString("<statements>\n")
	for {
		if e.tk.Keyword() == "let" {
			e.CompileLet()
			e.tk.Advance()
		} else if e.tk.Keyword() == "if" {
			e.CompileIf()
			fmt.Println("----- after if ends, token: ", e.tk.Token())
		} else if e.tk.Keyword() == "while" {
			e.CompileWhile()
			e.tk.Advance()
		} else if e.tk.Keyword() == "do" {
			e.CompileDo()
			e.tk.Advance()
		} else if e.tk.Keyword() == "return" {
			e.CompileReturn()
			e.tk.Advance()
		} else {
			// log.Fatal("CompileStatements, expect a statement keyword (let | if | while | do | return)")
			break
		}

	}
	// e.WriteString("</statements>\n")
}

/* 'let' varName('['expression']')? '=' expression ';' */
func (e *Engine) CompileLet() {
	fmt.Println("--- CompileLet ---")

	// e.WriteString("<letStatement>\n")
	// let
	// e.writeKeyword()

	e.tk.Advance()
	// varName
	if e.tk.Identifier() == "" {
		log.Fatal("CompileLet, expect a varName(identifier), got:", e.tk.Token(), " ", e.tk.TokenType())
	}

	// kind := e.getKindOfIdentifier(e.tk.Identifier())
	// e.writeIdentifier(e.tk.Identifier(), "used", kind)

	e.tk.Advance()
	// '['
	if e.tk.Symbol() == "[" {
		// e.writeSymbol()

		e.tk.Advance()
		// expression
		e.CompileExpression()
		// end expression

		if e.tk.Symbol() != "]" {
			log.Fatal("CompileLet, expect a ]")
		}
		// ']'
		// e.writeSymbol()
		e.tk.Advance()
	}

	if e.tk.Symbol() != "=" {
		log.Fatal("CompileLet, expect = or [")

	}

	// '='
	// e.writeSymbol()

	e.tk.Advance()
	// expression
	e.CompileExpression()
	// end expression

	if e.tk.Symbol() != ";" {
		log.Fatal("CompileLet, expect a ';', got: ", e.tk.Token())
	}
	// ';'
	// e.writeSymbol()
	// e.WriteString("</letStatement>\n")
}

/* 'if' '(' expression ')' '{' statements '}' ('else' '{' statements '}')? */
func (e *Engine) CompileIf() {
	fmt.Println("--- CompileIf ---")
	// if
	// e.WriteString("<ifStatement>\n")

	// e.writeKeyword()

	e.tk.Advance()
	if e.tk.Symbol() != "(" {
		log.Fatal("CompileIf, expect a '('")
	}
	// e.writeSymbol()

	e.tk.Advance()
	e.CompileExpression()

	if e.tk.Symbol() != ")" {
		log.Fatal("CompileIf, expect a ')'")
	}
	// e.writeSymbol()

	e.tk.Advance()
	if e.tk.Symbol() != "{" {
		log.Fatal("CompileIf, expect a '{'")
	}
	// e.writeSymbol()

	e.tk.Advance()
	e.CompileStatements()

	if e.tk.Symbol() != "}" {
		log.Fatal("CompileIf, expect a '}'")
	}
	// e.writeSymbol()

	e.tk.Advance()
	// else
	if e.tk.Keyword() == "else" {
		// e.writeKeyword()

		e.tk.Advance()
		if e.tk.Symbol() != "{" {
			log.Fatal("CompileIf, expect a '{'")
		}
		// e.writeSymbol()

		e.tk.Advance()
		e.CompileStatements()

		if e.tk.Symbol() != "}" {
			log.Fatal("CompileIf, expect a '}'")
		}
		// e.writeSymbol()
		e.tk.Advance()
	}
	// e.WriteString("</ifStatement>\n")

}

/* 'while' '(' expression ')' '{' expressions '}' */
func (e *Engine) CompileWhile() {
	fmt.Println("--- CompileWhile ---")

	// e.WriteString("<whileStatement>\n")
	// e.writeKeyword()

	e.tk.Advance()
	if e.tk.Symbol() != "(" {
		log.Fatal("CompileIf, expect a '('")
	}
	// e.writeSymbol()

	e.tk.Advance()
	e.CompileExpression()

	if e.tk.Symbol() != ")" {
		log.Fatal("CompileIf, expect a ')'")
	}
	// e.writeSymbol()

	e.tk.Advance()
	if e.tk.Symbol() != "{" {
		log.Fatal("CompileIf, expect a '{'")
	}
	// e.writeSymbol()

	e.tk.Advance()
	e.CompileStatements()

	if e.tk.Symbol() != "}" {
		log.Fatal("CompileIf, expect a '}'")
	}
	// e.writeSymbol()
	// e.WriteString("</whileStatement>\n")

}

// /* 'do' subroutineCall ';' */
// func (e *Engine) CompileDo() {
// 	fmt.Println("--- CompileDo ---")
// 	// e.WriteString("<doStatement>\n")
// 	// e.writeKeyword()

// 	e.tk.Advance()
// 	if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
// 		log.Fatal("CompileDo, expect an identifier")
// 	}

// 	// could be [subroutineName, (className | varName).subroutineName]
// 	prevId := e.tk.Identifier()

// 	e.tk.Advance()
// 	if e.tk.Symbol() == "(" {
// 		// then "prevId" is a subroutine Name

// 		// e.writeIdentifier(prevId, "used", symbolTable.SUBROUTINE)

// 		// e.writeSymbol()

// 		e.tk.Advance()
// 		e.CompileExpressionList()

// 		if e.tk.Symbol() != ")" {
// 			log.Fatal("CompileDo, expect a ')'")
// 		}
// 		// e.writeSymbol()

// 	} else if e.tk.Symbol() == "." {
// 		// "prevId" is either a className or a varName
// 		if e.subroutineST.KindOf(prevId) != "" {
// 			// e.writeIdentifier(prevId, "used", e.subroutineST.KindOf(prevId))
// 		} else if e.classST.KindOf(prevId) != "" {
// 			// e.writeIdentifier(prevId, "used", e.classST.KindOf(prevId))
// 		} else {
// 			// e.writeIdentifier(prevId, "used", symbolTable.CLASS)
// 		}

// 		// .
// 		// e.writeSymbol()

// 		e.tk.Advance()
// 		if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
// 			log.Fatal("CompileDo className|varName (identifier) (expect identifier), got:", e.tk.Token())
// 		}
// 		// e.writeIdentifier(e.tk.Identifier(), "used", symbolTable.SUBROUTINE)

// 		e.tk.Advance()
// 		if e.tk.Symbol() != "(" {
// 			log.Fatal("CompileDo expect '('")
// 		}
// 		// e.writeSymbol()

// 		e.tk.Advance()
// 		e.CompileExpressionList()

// 		if e.tk.Symbol() != ")" {
// 			log.Fatal("CompileDo, expect a ')'")
// 		}
// 		// e.writeSymbol()
// 	} else {
// 		log.Fatal("CompileDo not supported token, got: ", e.tk.Token())
// 	}

// 	e.tk.Advance()
// 	if e.tk.Symbol() != ";" {
// 		log.Fatal("CompileDo, expect a ';', got: ", e.tk.Token())
// 	}
// 	// e.writeSymbol()
// 	// e.WriteString("</doStatement>\n")

// 	e.vmWriter.WritePop("temp", 0)
// }

/* 'do' subroutineCall ';' */
func (e *Engine) CompileDo() {
	fmt.Println("--- CompileDo ---")
	e.tk.Advance()

	e.CompileExpression()
	e.vmWriter.WritePop("temp", 0)
}

/* 'return' expression? ';' */
func (e *Engine) CompileReturn() {
	fmt.Println("--- CompileReturn ---")
	// e.WriteString("<returnStatement>\n")

	// e.writeKeyword()

	e.tk.Advance()
	if e.tk.Symbol() != ";" {
		e.CompileExpression()
		e.vmWriter.WriteReturn()
	} else {
		e.vmWriter.WritePush(vmWriter.SEG_CONSTANT, 0)
		e.vmWriter.WriteReturn()
	}

	if e.tk.Symbol() != ";" {
		log.Fatal("CompileLet, expect a ';', got: ", e.tk.Token())
	}
	// e.writeSymbol()

	// e.WriteString("</returnStatement>\n")

}
