package compilationEngine

import "fmt"

func (e *Engine) writeSymbol() {
	e.WriteString(fmt.Sprintf("<symbol> %s </symbol>\n", e.tk.Symbol()))
}

func (e *Engine) writeKeyword() {
	e.WriteString(fmt.Sprintf("<keyword> %s </keyword>\n", e.tk.Keyword()))
}

func (e *Engine) writeIdentifier() {
	e.WriteString(fmt.Sprintf("<identifier> %s </identifier>\n", e.tk.Identifier()))
}

func (e *Engine) writeIntegerConst() {
	e.WriteString(fmt.Sprintf("<integerConstant> %d </integerConstant>\n", e.tk.IntVal()))
}

func (e *Engine) writeStringConst() {
	e.WriteString(fmt.Sprintf("<stringConstant> %s </stringConstant>\n", e.tk.StringVal()))
}
