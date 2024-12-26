package compilationEngine

import (
	"fmt"
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
		// op
		e.writeSymbol()

		e.tk.Advance()
		e.CompileTerm()

		if !e.tk.SkipAdvance() {
			e.tk.Advance()
		}
	}

	e.WriteString("</expression>\n")
	fmt.Println("	END CompileExpression")
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
	e.tk.SetSkipAdvance(false)
	fmt.Println("	END CompileExpressionList")
}
