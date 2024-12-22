package compilationEngine

import (
	"fmt"
)

func (e *Engine) CompileStatements() {
	fmt.Println("--- CompileStatements ---")
	fmt.Println("token: ", e.tk.Token())

	// no statements
	if !statementKeywords[e.tk.Keyword()] {
		return
	}

	e.writeKeyword()

	e.tk.Advance()
}
