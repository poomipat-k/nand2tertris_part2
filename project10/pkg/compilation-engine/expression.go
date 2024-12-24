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

	e.tk.Advance()
	// (op term)*
	for opSymbol[e.tk.Symbol()] {
		e.writeSymbol()

		e.tk.Advance()
		e.CompileTerm()

		e.tk.Advance()
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

	if !e.isTerm() {
		log.Fatal("CompileTerm, expect a term, got:", e.tk.Token())
	}
	fmt.Println("== isTerm")
	e.WriteString("<term>\n")
	tokenType := e.tk.TokenType()
	if tokenType == jackTokenizer.INT_CONST {
		e.writeIntegerConst()
	} else if tokenType == jackTokenizer.STRING_CONST {
		e.writeStringConst()
	} else if keywordConstant[e.tk.Keyword()] {
		e.writeKeyword()
	} else if tokenType == jackTokenizer.IDENTIFIER {
		e.writeIdentifier()
	} else {
		log.Fatal("CompileTerm, unsupported term, got:", e.tk.Token())
	}

	e.WriteString("</term>\n")

}

/* expressionList: (expression(',' expression)*)? */
func (e *Engine) CompileExpressionList() {
	fmt.Println("--- CompileExpressionList ---")

}

func (e *Engine) isTerm() bool {
	tokenType := e.tk.TokenType()
	if tokenType == jackTokenizer.INT_CONST ||
		tokenType == jackTokenizer.STRING_CONST ||
		keywordConstant[e.tk.Keyword()] ||
		e.tk.Identifier() != "" {
		return true
	}
	return false
}
