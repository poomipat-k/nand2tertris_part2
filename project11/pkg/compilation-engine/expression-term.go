package compilationEngine

import (
	"fmt"
	"log"

	symbolTable "github.com/poomipat-k/nand2tetris/project11/pkg/symbol-table"
	jackTokenizer "github.com/poomipat-k/nand2tetris/project11/pkg/tokenizer"
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
		e.writeSymbol()

		e.tk.Advance()
		e.CompileExpression()

		if e.tk.Symbol() != ")" {
			log.Fatal("CompileTerm, (expression) expect a closing ), got", e.tk.Token())
		}
		e.writeSymbol()

		// eg. let i = i * (-j); we need to reset skipAdvance to false after j
		e.tk.SetSkipAdvance(false)

	} else if unaryOp[e.tk.Symbol()] {
		e.writeSymbol()
		e.tk.Advance()
		e.CompileTerm()
	} else if tokenType == jackTokenizer.IDENTIFIER {
		prevId := e.tk.Identifier()

		e.tk.Advance()
		if e.tk.Symbol() == "[" {
			// "prevId" is either in subroutineST or classST
			if e.subroutineST.KindOf(prevId) != "" {
				e.writeIdentifier(prevId, "used", e.subroutineST.KindOf(prevId))
			} else if e.classST.KindOf(prevId) != "" {
				e.writeIdentifier(prevId, "used", e.classST.KindOf(prevId))
			} else {
				log.Fatal("CompileTerm, this should be in one of symbol tables, token: ", prevId)
			}

			e.writeIdentifier(e.tk.Identifier(), "used", "")

			e.writeSymbol()

			e.tk.Advance()
			e.CompileExpression()

			if e.tk.Symbol() != "]" {
				log.Fatal("CompileTerm, expect a ']', got: ", e.tk.Token())
			}
			e.writeSymbol()

			// let sum = sum + a[i];
			e.tk.SetSkipAdvance(false)
		} else if e.tk.Symbol() == "(" {
			e.writeIdentifier(prevId, "used", symbolTable.SUBROUTINE)

			e.writeSymbol()

			e.tk.Advance()
			e.CompileExpressionList()

			if e.tk.Symbol() != ")" {
				log.Fatal("CompileTerm, expect a ')'")
			}
			e.writeSymbol()

		} else if e.tk.Symbol() == "." {
			// "prevId" is either a className or a varName
			if e.subroutineST.KindOf(prevId) != "" {
				e.writeIdentifier(prevId, "used", e.subroutineST.KindOf(prevId))
			} else if e.classST.KindOf(prevId) != "" {
				e.writeIdentifier(prevId, "used", e.classST.KindOf(prevId))
			} else {
				e.writeIdentifier(prevId, "used", symbolTable.CLASS)
			}

			// .
			e.writeSymbol()

			e.tk.Advance()
			if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
				log.Fatal("CompileTerm className|varName (identifier) (expect identifier), got:", e.tk.Token())
			}
			e.writeIdentifier(e.tk.Identifier(), "used", symbolTable.SUBROUTINE)

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
			if e.subroutineST.KindOf(prevId) != "" {
				e.writeIdentifier(prevId, "used", e.subroutineST.KindOf(prevId))
			} else if e.classST.KindOf(prevId) != "" {
				e.writeIdentifier(prevId, "used", e.classST.KindOf(prevId))
			} else {
				log.Fatal("CompileTerm, this should be a var name")
			}
			e.tk.SetSkipAdvance(true)
		}
	} else {
		log.Fatal("CompileTerm, unsupported term, got:", e.tk.Token())
	}

	e.WriteString("</term>\n")
	fmt.Println("		END CompileTerm")
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
