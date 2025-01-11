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

/*
'let' varName('['expression']')? '=' expression ';'

compileLet

	compileExpression
	output "pop segment i"
*/
func (e *Engine) CompileLet() {
	fmt.Println("--- CompileLet ---")

	// let

	e.tk.Advance()
	// varName
	if e.tk.Identifier() == "" {
		log.Fatal("CompileLet, expect a varName(identifier), got:", e.tk.Token(), " ", e.tk.TokenType())
	}

	varName := e.tk.Identifier()
	kind := e.getKindOfIdentifier(varName)
	segment := e.vmWriter.KindToSegment(kind)
	index := e.getIndexOfIdentifier(varName)
	if kind == "" {
		log.Fatal("CompileLet, varName: ", varName, " is not in any symbol tables, kind: ", kind)
	}

	e.tk.Advance()
	// '[' go on with varName[expression]
	if e.tk.Symbol() == "[" {
		// Array varName
		e.vmWriter.WritePush(segment, index)

		e.tk.Advance()
		e.CompileExpression()

		if e.tk.Symbol() != "]" {
			log.Fatal("CompileLet, expect a ]")
		}
		// ']'

		e.vmWriter.WriteArithmetic("add")

		e.tk.Advance()
		if e.tk.Symbol() != "=" {
			log.Fatal("CompileLet, expect = or [")
		}

		// '='

		e.tk.Advance()
		e.CompileExpression()

		if e.tk.Symbol() != ";" {
			log.Fatal("CompileLet, expect a ';', got: ", e.tk.Token())
		}
		// ';'
		e.vmWriter.WritePop(vmWriter.SEG_TEMP, 0)
		e.vmWriter.WritePop(vmWriter.SEG_POINTER, 1)
		e.vmWriter.WritePush(vmWriter.SEG_TEMP, 0)
		e.vmWriter.WritePop(vmWriter.SEG_THAT, 0)
		return
	}

	// let varName = expression
	if e.tk.Symbol() != "=" {
		log.Fatal("CompileLet, expect = or [")
	}

	// '='

	e.tk.Advance()
	e.CompileExpression()

	if e.tk.Symbol() != ";" {
		log.Fatal("CompileLet, expect a ';', got: ", e.tk.Token())
	}
	// ';'
	e.vmWriter.WritePop(segment, index)
}

/* 'if' '(' expression ')' '{' statements '}' ('else' '{' statements '}')? */
func (e *Engine) CompileIf() {
	fmt.Println("--- CompileIf ---")
	// if

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

	// not
	e.vmWriter.WriteArithmetic("not")
	// if-goto label1
	label1 := generateLabel(e.className)
	e.vmWriter.WriteIf(label1)

	e.tk.Advance()
	e.CompileStatements() // statements within if block

	if e.tk.Symbol() != "}" {
		log.Fatal("CompileIf, expect a '}'")
	}
	// e.writeSymbol()

	e.tk.Advance()
	// else
	label2 := generateLabel(e.className)
	if e.tk.Keyword() == "else" {
		// e.writeKeyword()

		e.tk.Advance()
		if e.tk.Symbol() != "{" {
			log.Fatal("CompileIf, expect a '{'")
		}
		// e.writeSymbol()

		e.vmWriter.WriteGoto(label2)

		e.vmWriter.WriteLabel(label1)

		e.tk.Advance()
		e.CompileStatements() // else statement

		if e.tk.Symbol() != "}" {
			log.Fatal("CompileIf, expect a '}'")
		}
		// e.writeSymbol()

		e.tk.Advance()

		e.vmWriter.WriteLabel(label2)
	} else {
		// no else block
		e.vmWriter.WriteLabel(label1)
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
		log.Fatal("CompileWhile, expect a '('")
	}
	// e.writeSymbol()

	label1 := generateLabel(e.className)
	e.vmWriter.WriteLabel(label1)

	e.tk.Advance()
	e.CompileExpression()

	e.vmWriter.WriteArithmetic("not")
	label2 := generateLabel(e.className)
	e.vmWriter.WriteIf(label2)

	if e.tk.Symbol() != ")" {
		log.Fatal("CompileWhile, expect a ')'")
	}
	// e.writeSymbol()

	e.tk.Advance()
	if e.tk.Symbol() != "{" {
		log.Fatal("CompileWhile, expect a '{'")
	}
	// e.writeSymbol()

	e.tk.Advance()
	e.CompileStatements()

	e.vmWriter.WriteGoto(label1)

	e.vmWriter.WriteLabel(label2)

	if e.tk.Symbol() != "}" {
		log.Fatal("CompileWhile, expect a '}'")
	}
	// e.writeSymbol()
	// e.WriteString("</whileStatement>\n")

}

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
	if e.tk.Symbol() == ";" {
		e.vmWriter.WritePush(vmWriter.SEG_CONSTANT, 0)
		e.vmWriter.WriteReturn()
	} else {
		e.CompileExpression()
		e.vmWriter.WriteReturn()
	}

	if e.tk.Symbol() != ";" {
		log.Fatal("CompileLet, expect a ';', got: ", e.tk.Token())
	}
	// e.writeSymbol()

	// e.WriteString("</returnStatement>\n")

}
