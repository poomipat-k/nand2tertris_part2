package compilationEngine

import (
	"fmt"
	"log"

	jackTokenizer "github.com/poomipat-k/nand2tetris/project10/pkg/tokenizer"
)

/*
statements: statement*
statement: letStatement | ifStatement | whileStatement | doStatement | returnStatement
*/
func (e *Engine) CompileStatements() {
	fmt.Println("--- CompileStatements ---")

	// no statements
	if !statementKeywords[e.tk.Keyword()] {
		e.WriteString("<statements>\n")
		e.WriteString("</statements>\n")
		return
	}

	e.WriteString("<statements>\n")
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

	fmt.Println("	Let before ;, token: ", e.tk.Token(), " type:", e.tk.TokenType(), " skipAdvance: ", e.tk.SkipAdvance())
	if e.tk.Symbol() != ";" {
		log.Fatal("CompileLet, expect a ';', got: ", e.tk.Token())
	}
	// ';'
	e.writeSymbol()
	e.WriteString("</letStatement>\n")
}

/* 'if' '(' expression ')' '{' statements '}' ('else' '{' statements '}')? */
func (e *Engine) CompileIf() {
	fmt.Println("--- CompileIf ---")
	// if
	e.WriteString("<ifStatement>\n")

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
	// else
	if e.tk.Keyword() == "else" {
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
		e.tk.Advance()
	}
	e.WriteString("</ifStatement>\n")

}

/* 'while' '(' expression ')' '{' expressions '}' */
func (e *Engine) CompileWhile() {
	fmt.Println("--- CompileWhile ---")

	e.WriteString("<whileStatement>\n")
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
	e.WriteString("</whileStatement>\n")

}

/* 'do' subroutineCall ';' */
func (e *Engine) CompileDo() {
	fmt.Println("--- CompileDo ---")
	e.WriteString("<doStatement>\n")
	e.writeKeyword()

	e.tk.Advance()
	if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
		log.Fatal("CompileDo, expect an identifier")
	}
	e.writeIdentifier()

	e.tk.Advance()
	if e.tk.Symbol() == "(" {
		e.writeSymbol()

		e.tk.Advance()
		e.CompileExpressionList()

		if e.tk.Symbol() != ")" {
			log.Fatal("CompileDo, expect a ')'")
		}
		e.writeSymbol()

	} else if e.tk.Symbol() == "." {
		// className or varName
		e.writeSymbol()

		e.tk.Advance()
		if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
			log.Fatal("CompileDo className|varName (identifier) (expect identifier), got:", e.tk.Token())
		}
		e.writeIdentifier()

		e.tk.Advance()
		if e.tk.Symbol() != "(" {
			log.Fatal("CompileDo expect '('")
		}
		e.writeSymbol()

		e.tk.Advance()
		e.CompileExpressionList()

		if e.tk.Symbol() != ")" {
			log.Fatal("CompileDo, expect a ')'")
		}
		e.writeSymbol()
	} else {
		log.Fatal("CompileDo not supported token, got: ", e.tk.Token())
	}

	e.tk.Advance()
	if e.tk.Symbol() != ";" {
		log.Fatal("CompileLet, expect a ';', got: ", e.tk.Token())
	}
	e.writeSymbol()
	e.WriteString("</doStatement>\n")
}

/* 'return' expression? ';' */
func (e *Engine) CompileReturn() {
	fmt.Println("--- CompileReturn ---")
	e.WriteString("<returnStatement>\n")

	e.writeKeyword()

	e.tk.Advance()
	if e.tk.Symbol() != ";" {
		e.CompileExpression()
	}

	if e.tk.Symbol() != ";" {
		log.Fatal("CompileLet, expect a ';', got: ", e.tk.Token())
	}
	e.writeSymbol()

	e.WriteString("</returnStatement>\n")

}
