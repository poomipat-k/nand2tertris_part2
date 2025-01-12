package compilationEngine

/* expression: term (op term)* */
func (e *Engine) CompileExpression() {
	// fmt.Println("--- CompileExpression ---, token: ", e.tk.Token())

	// term
	e.CompileTerm()
	// end term

	if !e.tk.SkipAdvance() {
		e.tk.Advance()
	}
	// (op term)*
	_, isOp := opSymbol[e.tk.Symbol()]
	for isOp {
		// op
		op := e.tk.Symbol()

		// e.writeSymbol()

		e.tk.Advance()
		e.CompileTerm()

		// write op
		e.vmWriter.WriteArithmetic(opSymbol[op])

		if !e.tk.SkipAdvance() {
			e.tk.Advance()
		}

		_, isOp = opSymbol[e.tk.Symbol()]
	}
}

/* expressionList: (expression(',' expression)*)? */
func (e *Engine) CompileExpressionList() int {
	// fmt.Println("--- CompileExpressionList ---, token: ", e.tk.Token())

	nArgs := 0
	if !e.isTerm() {
		return nArgs
	}
	e.CompileExpression()
	nArgs++

	for e.tk.Symbol() == "," {

		e.tk.Advance()
		e.CompileExpression()
		nArgs++
	}
	e.tk.SetSkipAdvance(false)
	return nArgs
}
