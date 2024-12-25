package compilationEngine

import (
	"fmt"
	"log"

	jackTokenizer "github.com/poomipat-k/nand2tetris/project10/pkg/tokenizer"
)

/* expression: term (op term)* */
func (e *Engine) CompileExpression() {
	fmt.Println("--- CompileExpression ---")
	e.WriteString("<expression>\n")

	// term
	e.CompileTerm()
	// end term

	if !e.tk.SkipAdvance() {
		e.tk.Advance()
	}
	// (op term)*
	for opSymbol[e.tk.Symbol()] {
		e.writeSymbol()

		e.tk.Advance()
		e.CompileTerm()

		if !e.tk.SkipAdvance() {
			e.tk.Advance()
		}
	}

	e.WriteString("</expression>\n")
}

/*
term:

	integerConstant | stringConstant | keywordConstant | varName |
	varName'['expression']' | subroutineCall | '('expression')' | unaryOp term
*/
func (e *Engine) CompileTerm() {
	fmt.Println("--- CompileTerm ---")

	// if !e.skipTermTag {
	// 	e.WriteString("<term>\n")
	// }
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
	} else if unaryOp[e.tk.Symbol()] {
		fmt.Println("==unaryOp")
		e.writeSymbol()
		e.tk.Advance()
		e.CompileTerm()
	} else if tokenType == jackTokenizer.IDENTIFIER {
		e.writeIdentifier()

		e.tk.Advance()
		if e.tk.Symbol() == "[" {
			e.writeSymbol()

			e.tk.Advance()
			e.CompileExpression()

			e.tk.Advance()
			if e.tk.Symbol() != "]" {
				log.Fatal("CompileTerm, expect a ']'")
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
			// varName do nothing
			// fmt.Println("===varName, then set skipAdvance because it is advanced already")
			e.tk.SetSkipAdvance(true)
		}
	} else {
		log.Fatal("CompileTerm, unsupported term, got:", e.tk.Token())
	}
	// if !e.skipTermTag {
	// 	e.WriteString("</term>\n")
	// }
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

/* expressionList: (expression(',' expression)*)? */
func (e *Engine) CompileExpressionList() {
	fmt.Println("--- CompileExpressionList ---")
	e.WriteString("<expressionList>\n")
	if !e.isTerm() {
		e.WriteString("</expressionList>\n")
		return
	}
	e.CompileExpression()

	for e.tk.Symbol() == "," {
		e.writeSymbol()

		e.tk.Advance()
		e.CompileExpression()
	}
	e.WriteString("</expressionList>\n")
}
