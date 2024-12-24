package compilationEngine

import (
	"fmt"

	jackTokenizer "github.com/poomipat-k/nand2tetris/project10/pkg/tokenizer"
)

/* expression: term (op term)* */
func (e *Engine) CompileExpression() {
	fmt.Println("--- CompileExpression ---")
	e.WriteString("<expression>\n")

	e.WriteString("</expression>\n")

}

/*
term:

	integerConstant | stringConstant | keywordConstant | varName |
	varName'['expression']' | subroutineCall | '('expression')' | unaryOp term
*/
func (e *Engine) CompileTerm() {
	fmt.Println("--- CompileTerm ---")

	if e.isTerm() {
		fmt.Println("== isTerm")
		e.WriteString("<term>\n")

		e.WriteString("</term>\n")
	}

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
