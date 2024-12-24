package compilationEngine

import "fmt"

/** expression: term (op term)* */
func (e *Engine) CompileExpression() {
	fmt.Println("--- CompileExpression ---")

}

/*
integerConstant | stringConstant | keywordConstant | varName |
varName'['expression']' | subroutineCall | '('expression')' | unaryOp term
*/
func (e *Engine) CompileTerm() {
	fmt.Println("--- CompileTerm ---")

}

/** expressionList: (expression(',' expression)*)? */
func (e *Engine) CompileExpressionList() {
	fmt.Println("--- CompileExpressionList ---")

}
