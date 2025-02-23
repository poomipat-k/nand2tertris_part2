package compilationEngine

import (
	"fmt"
	"log"

	symbolTable "github.com/poomipat-k/nand2tetris/project11/pkg/symbol-table"
	jackTokenizer "github.com/poomipat-k/nand2tetris/project11/pkg/tokenizer"
	vmWriter "github.com/poomipat-k/nand2tetris/project11/pkg/vm-writer"
)

/* class: 'class' className '{' classVarDec* subroutineDec* '}' */
func (e *Engine) CompileClass() {
	// fmt.Println("---- compile class ----")
	if e.tk.Keyword() != "class" {
		log.Fatal("current token is not a 'class'")
	}

	e.classST.Reset()

	e.tk.Advance()
	if e.tk.Identifier() == "" {
		log.Fatal("expect an identifier (className)")
	}
	e.className = e.tk.Identifier()

	e.tk.Advance()
	if e.tk.Symbol() != "{" {
		log.Fatal("expect an '{'")
	}

	e.tk.Advance()
	e.CompileClassVarDec()   // classVarDec*
	e.CompileSubroutineDec() // subRoutine*

	if e.tk.Symbol() != "}" {
		log.Fatal("expect a '}' at the end of a class, got: ", e.tk.Token())
	}
}

/*
(('static' | 'field') type varName (',' varName)* ';')*
*/
func (e *Engine) CompileClassVarDec() {
	// fmt.Println("---- compile classVarDec ----")
	if !classVarScope[e.tk.Keyword()] {
		return
	}

	// at least one classVarDec exist
	for {
		if e.tk.Symbol() == "}" || subroutineDec[e.tk.Keyword()] {
			break
		}
		e.compileOneClassVarDec()
		e.tk.Advance()
	}
}

// ('static' | 'field') type varName (',' varName)* ';'
func (e *Engine) compileOneClassVarDec() {
	if !classVarScope[e.tk.Keyword()] {
		log.Fatal("expect static or field")
	}
	// static | field
	isStatic := false
	if e.tk.Keyword() == "static" {
		isStatic = true
	}

	e.tk.Advance()
	// type
	var dataType string
	if e.tk.TokenType() == jackTokenizer.KEYWORD && jackType[e.tk.Keyword()] {
		dataType = e.tk.Keyword()
	} else if e.tk.TokenType() == jackTokenizer.IDENTIFIER {
		dataType = e.tk.Identifier()
	} else {
		log.Fatal("expect 'int' | 'char' | 'boolean' | className(identifier)")
	}

	e.tk.Advance()
	// varName
	if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
		log.Fatal("expect identifier")
	}

	varNameKind := symbolTable.FIELD
	if isStatic {
		varNameKind = symbolTable.STATIC
	}
	// add variable to the class symbolTable
	e.classST.Define(e.tk.Identifier(), dataType, varNameKind)

	e.tk.Advance()
	if e.tk.TokenType() != jackTokenizer.SYMBOL || (e.tk.Symbol() != "," && e.tk.Symbol() != ";") {
		log.Fatal("expect ',' or ';'")
	}

	// has more than one variable
	for e.tk.Symbol() != ";" {
		if e.tk.Symbol() == "," {
		} else if e.tk.Identifier() != "" {
			e.classST.Define(e.tk.Identifier(), dataType, varNameKind)
		} else {
			log.Fatal("expect ',' or identifier or ';'")
		}
		e.tk.Advance()
	}

	// write ;
}

/* ('constructor' | 'function' | 'method') ('void' | type) subroutineName '(' parameterList ')' subroutineBody  */
func (e *Engine) CompileSubroutineDec() {

	for subroutineDec[e.tk.Keyword()] {
		e.subroutineST.Reset()

		e.subroutineType = e.tk.Keyword()

		e.tk.Advance()
		// subroutine return type
		if e.tk.TokenType() == jackTokenizer.KEYWORD && (jackType[e.tk.Keyword()] || e.tk.Keyword() == "void") {

		} else if e.tk.TokenType() == jackTokenizer.IDENTIFIER {

		} else {
			log.Fatal("CompileSubroutineDec, expect to be one of 'void' | 'int' | 'char' | 'boolean' | className(identifier)", " got: ", e.tk.Token(), " type: ", e.tk.TokenType())
		}

		e.tk.Advance()
		// subroutine name
		if e.tk.Identifier() == "" {
			log.Fatal("expect an identifier (subRoutineName)")
		}
		subroutineName := e.tk.Identifier()
		e.subroutineName = subroutineName

		e.tk.Advance()
		if e.tk.Symbol() != "(" {
			log.Fatal("expect a '('")
		}

		// parameterList
		e.tk.Advance()
		e.CompileParameterList()
		// end parameterList

		if e.tk.Symbol() != ")" {
			log.Fatal("expect a ')'")
		}

		e.tk.Advance()

		e.CompileSubroutineBody()

		e.tk.Advance()
	}
}

func (e *Engine) CompileParameterList() {
	// fmt.Println("--- CompileParameterList ---")
	if e.subroutineType == METHOD {
		e.subroutineST.Define("this", e.className, symbolTable.ARG)
	}

	for e.tk.Symbol() != ")" {
		// type
		var dataType string
		if e.tk.TokenType() == jackTokenizer.KEYWORD && jackType[e.tk.Keyword()] {
			dataType = e.tk.Keyword()
		} else if e.tk.TokenType() == jackTokenizer.IDENTIFIER {
			// className type
			dataType = e.tk.Identifier()
		} else {
			log.Fatal("expect 'int' | 'char' | 'boolean' | className(identifier)")
		}

		e.tk.Advance()
		// varName
		if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
			log.Fatal("parameterList varName: expect an identifier")
		}
		e.subroutineST.Define(e.tk.Identifier(), dataType, symbolTable.ARG)

		e.tk.Advance()
		// optional ","
		if e.tk.Symbol() == "," {
			e.tk.Advance()
		}
	}
}

/* '{' varDec* statements '}' */
func (e *Engine) CompileSubroutineBody() {
	// fmt.Println("--- CompileSubroutineBody ---")

	if e.tk.Symbol() != "{" {
		log.Fatal("CompileSubroutineBody expect a '{'")
	}

	// varDec*
	e.tk.Advance()
	e.CompileVarDec()
	// end varDec*

	if e.subroutineType == CONSTRUCTOR {
		e.vmWriter.WriteFunction(fmt.Sprintf("%s.%s", e.className, e.subroutineName), e.subroutineST.VarCount(symbolTable.VAR))

		fieldCount := e.classST.VarCount(symbolTable.FIELD)
		e.vmWriter.WritePush(vmWriter.SEG_CONSTANT, fieldCount)
		e.vmWriter.WriteCall("Memory.alloc", 1)
		e.vmWriter.WritePop(vmWriter.SEG_POINTER, 0)
	} else if e.subroutineType == METHOD {
		localVarCount := e.subroutineST.VarCount(symbolTable.VAR)
		// function Class.subroutine nLocals
		e.vmWriter.WriteFunction(fmt.Sprintf("%s.%s", e.className, e.subroutineName), localVarCount)
		e.vmWriter.WritePush(vmWriter.SEG_ARG, 0) // first argument of a method is always "this"
		e.vmWriter.WritePop(vmWriter.SEG_POINTER, 0)
	} else if e.subroutineType == FUNCTION {
		// function
		e.vmWriter.WriteFunction(fmt.Sprintf("%s.%s", e.className, e.subroutineName), e.subroutineST.VarCount(symbolTable.VAR))
	} else {
		log.Fatal("Should be one of [constructor, method, function]")
	}

	e.CompileStatements()

	if e.tk.Symbol() != "}" {
		log.Fatal("CompileSubroutineBody, expect a '}', got: ", e.tk.Token())
	}
}

/* varDec: 'var' type varName (',' varName)* ';' */
func (e *Engine) CompileVarDec() {
	// fmt.Println("--- CompileVarDec ---")

	for e.tk.Keyword() == "var" {
		// 'var'
		e.tk.Advance()

		// type
		var dataType string
		if e.tk.TokenType() == jackTokenizer.KEYWORD && jackType[e.tk.Keyword()] {
			dataType = e.tk.Keyword()
		} else if e.tk.TokenType() == jackTokenizer.IDENTIFIER {
			dataType = e.tk.Identifier()
		} else {
			log.Fatal("CompileVarDec, expect 'int' | 'char' | 'boolean' | className(identifier), got: ", e.tk.Token())
		}
		e.tk.Advance()

		// varName
		if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
			log.Fatal("CompileVarDec, varName: expect an identifier")
		}
		e.subroutineST.Define(e.tk.Identifier(), dataType, symbolTable.VAR)

		e.tk.Advance()
		for e.tk.Symbol() != ";" {
			// optional ","
			if e.tk.Symbol() != "," {
				log.Fatal("CompileVarDec, expect a ','")
			}

			e.tk.Advance()
			// varName
			if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
				log.Fatal("CompileVarDec, varName: expect an identifier after ,")
			}
			e.subroutineST.Define(e.tk.Identifier(), dataType, symbolTable.VAR)

			e.tk.Advance()
		}
		// write ';'
		e.tk.Advance()
	}

}
