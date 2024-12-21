package compilationEngine

import (
	"fmt"
	"log"
)

// class: 'class' className '{' classVarDec* subroutineDec* '}'
func (e *Engine) CompileClass() {
	fmt.Println("---- compile class ----")
	if e.tk.Keyword() != "class" {
		log.Fatal("current token is not a 'class'")
	}

	e.WriteString("<class>\n")
	e.WriteString(fmt.Sprintf("<keyword> %s </keyword>\n", e.tk.Keyword()))

	e.tk.Advance()
	if e.tk.Identifier() == "" {
		log.Fatal("expect an identifier (className)")
	}
	e.WriteString(fmt.Sprintf("<identifier> %s </identifier>", e.tk.Identifier()))

	e.tk.Advance()
	if e.tk.Symbol() != "{" {
		log.Fatal("expect an '{'")
	}
	e.WriteString(fmt.Sprintf("<symbol> %s </symbol>", e.tk.Symbol()))

	e.tk.Advance()
	e.CompileClassVarDec() // classVarDec*

	e.WriteString("</class>\n")
}

// ('static' | 'field') type varName (',' varName)* ';'
func (e *Engine) CompileClassVarDec() {
	fmt.Println("---- compile classVarDec ----")
	if e.tk.Keyword() != "static" && e.tk.Keyword() != "field" {
		return
	}

	e.WriteString("<classVarDec>\n")

	e.WriteString("</classVarDec>\n")
}
