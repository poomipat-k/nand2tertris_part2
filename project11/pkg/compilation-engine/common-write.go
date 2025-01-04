package compilationEngine

import (
	"fmt"
	"log"
	"strings"
)

var declareFunctionCall = map[string]bool{
	"CompileClass":         true,
	"CompileSubroutineDec": true,
	"CompileClassVarDec":   true,
	"CompileVarDec":        true,
}

func (e *Engine) writeSymbol() {
	e.WriteString(fmt.Sprintf("<symbol> %s </symbol>\n", e.tk.Symbol()))
}

func (e *Engine) writeKeyword() {
	e.WriteString(fmt.Sprintf("<keyword> %s </keyword>\n", e.tk.Keyword()))
}

func (e *Engine) writeIntegerConst() {
	e.WriteString(fmt.Sprintf("<integerConstant> %d </integerConstant>\n", e.tk.IntVal()))
}

func (e *Engine) writeStringConst() {
	e.WriteString(fmt.Sprintf("<stringConstant> %s </stringConstant>\n", e.tk.StringVal()))
}

func (e *Engine) writeIdentifier(calledFrom string) {
	if e.tk.Identifier() == "" {
		log.Fatal("writeIdentifier, identifier should not empty")
	}
	// tag format <identifier_(dec | used)_kind(_runningNumber)?>

	role := "dec"
	if !declareFunctionCall[calledFrom] {
		role = "used"
	}
	var kind string
	token := e.tk.Identifier()
	kind = e.subroutineST.KindOf(token)
	if kind == "" {
		kind = e.classST.KindOf(token)
	}
	if kind == "" {
		// either class or subroutine
		if calledFrom == "CompileClass" {
			kind = "CLASS"
		} else {
			kind = "SUBROUTINE"
		}
	}

	tag := fmt.Sprintf("identifier_%s_%s", role, strings.ToLower(kind))

	e.WriteString(fmt.Sprintf("<%s> %s </%s>\n", tag, e.tk.Identifier(), tag))
}
