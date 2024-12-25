package compilationEngine

import (
	"fmt"
	"log"

	jackTokenizer "github.com/poomipat-k/nand2tetris/project10/pkg/tokenizer"
)

/*
term:

	integerConstant | stringConstant | keywordConstant | varName |
	varName'['expression']' | subroutineCall | '('expression')' | unaryOp term
*/
func (e *Engine) CompileTerm() {
	fmt.Println("--- CompileTerm ---")

	e.WriteString("<term>\n")

	tokenType := e.tk.TokenType()

	if tokenType == jackTokenizer.INT_CONST {
		e.writeIntegerConst()
	} else if tokenType == jackTokenizer.STRING_CONST {
		e.writeStringConst()
	} else if keywordConstant[e.tk.Keyword()] {
		e.writeKeyword()
	} else if e.tk.Symbol() == "(" {
		fmt.Println("==(expression)")
		e.writeSymbol()

		e.tk.Advance()
		e.CompileExpression()

		if e.tk.Symbol() != ")" {
			log.Fatal("CompileTerm, expect a closing ), got", e.tk.Token())
		}
		e.writeSymbol()

		// eg. let i = i * (-j); we need to reset skipAdvance to false after j
		e.tk.SetSkipAdvance(false)
		fmt.Println("==end (expression) token: ", e.tk.Token(), " skipAdvance: ", e.tk.SkipAdvance())

	} else if unaryOp[e.tk.Symbol()] {
		fmt.Println("==unaryOp")
		e.writeSymbol()
		e.tk.Advance()
		e.CompileTerm()
		fmt.Println("==end unaryOp, token: ", e.tk.Token())
	} else if tokenType == jackTokenizer.IDENTIFIER {
		e.writeIdentifier()

		e.tk.Advance()
		if e.tk.Symbol() == "[" {
			e.writeSymbol()

			e.tk.Advance()
			e.CompileExpression()

			if e.tk.Symbol() != "]" {
				log.Fatal("CompileTerm, expect a ']', got: ", e.tk.Token())
			}
			e.writeSymbol()
		} else if e.tk.Symbol() == "(" {
			e.writeSymbol()

			e.tk.Advance()
			e.CompileExpressionList()

			if e.tk.Symbol() != ")" {
				log.Fatal("CompileTerm, expect a ')'")
			}
			e.writeSymbol()

		} else if e.tk.Symbol() == "." {
			// className or varName
			e.writeSymbol()

			e.tk.Advance()
			if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
				log.Fatal("CompileTerm className|varName (identifier) (expect identifier), got:", e.tk.Token())
			}
			e.writeIdentifier()

			e.tk.Advance()
			if e.tk.Symbol() != "(" {
				log.Fatal("CompileTerm expect '('")
			}
			e.writeSymbol()

			e.tk.Advance()
			e.CompileExpressionList()

			if e.tk.Symbol() != ")" {
				log.Fatal("CompileTerm, expect a ')'")
			}
			e.writeSymbol()
		} else {
			// skip advance if the current is varName
			e.tk.SetSkipAdvance(true)
			fmt.Println("		set skipAdvance true, token: ", e.tk.Token())
		}
	} else {
		log.Fatal("CompileTerm, unsupported term, got:", e.tk.Token())
	}

	e.WriteString("</term>\n")
}

func (e *Engine) isTerm() bool {
	tokenType := e.tk.TokenType()
	return tokenType == jackTokenizer.INT_CONST ||
		tokenType == jackTokenizer.STRING_CONST ||
		keywordConstant[e.tk.Keyword()] ||
		e.tk.Symbol() == "(" ||
		unaryOp[e.tk.Symbol()] ||
		tokenType == jackTokenizer.IDENTIFIER
}
