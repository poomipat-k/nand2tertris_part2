package compilationEngine

import (
	"fmt"
	"log"

	jackTokenizer "github.com/poomipat-k/nand2tetris/project11/pkg/tokenizer"
)

/* class: 'class' className '{' classVarDec* subroutineDec* '}' */
func (e *Engine) CompileClass() {
	fmt.Println("---- compile class ----")
	if e.tk.Keyword() != "class" {
		log.Fatal("current token is not a 'class'")
	}

	e.classST.Reset()

	e.WriteString("<class>\n")
	e.writeKeyword()

	e.tk.Advance()
	if e.tk.Identifier() == "" {
		log.Fatal("expect an identifier (className)")
	}
	e.writeIdentifier()

	e.tk.Advance()
	if e.tk.Symbol() != "{" {
		log.Fatal("expect an '{'")
	}

	e.writeSymbol()

	e.tk.Advance()
	e.CompileClassVarDec()   // classVarDec*
	e.CompileSubroutineDec() // subRoutine*

	if e.tk.Symbol() != "}" {
		log.Fatal("expect a '}' at the end of a class, got: ", e.tk.Token())
	}
	e.writeSymbol()

	e.WriteString("</class>\n")
}

/* (('static' | 'field') type varName (',' varName)* ';')* */
func (e *Engine) CompileClassVarDec() {
	fmt.Println("---- compile classVarDec ----")
	if !classVarScope[e.tk.Keyword()] {
		return
	}

	// at least one classVarDec exist
	for {
		if e.tk.Symbol() == "}" || subroutineDec[e.tk.Keyword()] {
			break
		}
		e.WriteString("<classVarDec>\n")
		e.compileOneClassVarDec()
		e.WriteString("</classVarDec>\n")
		e.tk.Advance()
	}
}

// ('static' | 'field') type varName (',' varName)* ';'
func (e *Engine) compileOneClassVarDec() {
	if !classVarScope[e.tk.Keyword()] {
		log.Fatal("expect static or field")
	}
	// static | field
	e.writeKeyword()

	e.tk.Advance()
	// type
	if e.tk.TokenType() == jackTokenizer.KEYWORD && jackType[e.tk.Keyword()] {
		e.writeKeyword()
	} else if e.tk.TokenType() == jackTokenizer.IDENTIFIER {
		e.writeIdentifier()
	} else {
		log.Fatal("expect 'int' | 'char' | 'boolean' | className(identifier)")
	}

	e.tk.Advance()
	// varName
	if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
		log.Fatal("expect identifier")
	}
	e.writeIdentifier()

	e.tk.Advance()
	if e.tk.TokenType() != jackTokenizer.SYMBOL || (e.tk.Symbol() != "," && e.tk.Symbol() != ";") {
		log.Fatal("expect ',' or ';'")
	}

	// has more than one variable
	for e.tk.Symbol() != ";" {
		if e.tk.Symbol() == "," {
			e.writeSymbol()
		} else if e.tk.Identifier() != "" {
			e.writeIdentifier()
		} else {
			log.Fatal("expect ',' or identifier or ';'")
		}
		e.tk.Advance()
	}

	// write ;
	e.writeSymbol()
}

/* ('constructor' | 'function' | 'method') ('void' | type) subroutineName '(' parameterList ')' subroutineBody  */
func (e *Engine) CompileSubroutineDec() {

	for subroutineDec[e.tk.Keyword()] {
		e.subroutineST.Reset()

		e.WriteString("<subroutineDec>\n")
		e.writeKeyword()

		e.tk.Advance()
		// subroutine return type
		if e.tk.TokenType() == jackTokenizer.KEYWORD && (jackType[e.tk.Keyword()] || e.tk.Keyword() == "void") {
			e.writeKeyword()
		} else if e.tk.TokenType() == jackTokenizer.IDENTIFIER {
			e.writeIdentifier()
		} else {
			log.Fatal("CompileSubroutineDec, expect to be one of 'void' | 'int' | 'char' | 'boolean' | className(identifier)", " got: ", e.tk.Token(), " type: ", e.tk.TokenType())
		}

		e.tk.Advance()
		if e.tk.Identifier() == "" {
			log.Fatal("expect an identifier (subRoutineName)")
		}
		e.writeIdentifier()

		e.tk.Advance()
		if e.tk.Symbol() != "(" {
			log.Fatal("expect a '('")
		}
		e.writeSymbol()

		// parameterList
		e.tk.Advance()
		e.CompileParameterList()
		// end parameterList

		if e.tk.Symbol() != ")" {
			log.Fatal("expect a ')'")
		}
		e.writeSymbol()

		e.tk.Advance()
		// subroutineBody
		e.CompileSubroutineBody()
		// end subroutineBody

		e.WriteString("</subroutineDec>\n")
		e.tk.Advance()
	}
}

func (e *Engine) CompileParameterList() {
	fmt.Println("--- CompileParameterList ---")
	e.WriteString("<parameterList>\n")

	for e.tk.Symbol() != ")" {
		// type
		if e.tk.TokenType() == jackTokenizer.KEYWORD && jackType[e.tk.Keyword()] {
			e.writeKeyword()
		} else if e.tk.TokenType() == jackTokenizer.IDENTIFIER {
			e.writeIdentifier()
		} else {
			log.Fatal("expect 'int' | 'char' | 'boolean' | className(identifier)")
		}

		e.tk.Advance()
		// varName
		if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
			log.Fatal("parameterList varName: expect an identifier")
		}
		e.writeIdentifier()

		e.tk.Advance()
		// optional ","
		if e.tk.Symbol() == "," {
			e.writeSymbol()
			e.tk.Advance()
		}
	}

	e.WriteString("</parameterList>\n")
}

/* '{' varDec* statements '}' */
func (e *Engine) CompileSubroutineBody() {
	fmt.Println("--- CompileSubroutineBody ---")

	e.WriteString("<subroutineBody>\n")
	if e.tk.Symbol() != "{" {
		log.Fatal("CompileSubroutineBody expect a '{'")
	}
	e.writeSymbol()

	// varDec*
	e.tk.Advance()
	e.CompileVarDec()
	// end varDec*

	// statements
	e.CompileStatements()
	// end statements

	if e.tk.Symbol() != "}" {
		log.Fatal("CompileSubroutineBody, expect a '}', got: ", e.tk.Token())
	}
	e.writeSymbol()

	e.WriteString("</subroutineBody>\n")

}

/* varDec: 'var' type varName (',' varName)* ';' */
func (e *Engine) CompileVarDec() {
	fmt.Println("--- CompileVarDec ---")
	i := 0
	for e.tk.Keyword() == "var" {
		e.WriteString("<varDec>\n")
		// 'var'
		e.writeKeyword()
		e.tk.Advance()

		// type
		if e.tk.TokenType() == jackTokenizer.KEYWORD && jackType[e.tk.Keyword()] {
			e.writeKeyword()
		} else if e.tk.TokenType() == jackTokenizer.IDENTIFIER {
			e.writeIdentifier()
		} else {
			log.Fatal("CompileVarDec, expect 'int' | 'char' | 'boolean' | className(identifier), got: ", e.tk.Token(), " i: ", i)
		}
		e.tk.Advance()

		// varName
		if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
			log.Fatal("CompileVarDec, varName: expect an identifier")
		}
		e.writeIdentifier()

		e.tk.Advance()
		for e.tk.Symbol() != ";" {
			// optional ","
			if e.tk.Symbol() != "," {
				log.Fatal("CompileVarDec, expect a ','")
			}
			e.writeSymbol()

			e.tk.Advance()
			// varName
			if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
				log.Fatal("CompileVarDec, varName: expect an identifier after ,")
			}
			e.writeIdentifier()
			e.tk.Advance()
		}
		// write ';'
		e.writeSymbol()
		e.WriteString("</varDec>\n")
		e.tk.Advance()

		i++
		if i >= 4 {
			fmt.Println("current: ", e.tk.Token())
			fmt.Println("exceed 4")
			return
		}
	}

}
