package compilationEngine

import (
	"fmt"
	"log"

	jackTokenizer "github.com/poomipat-k/nand2tetris/project11/pkg/tokenizer"
	vmWriter "github.com/poomipat-k/nand2tetris/project11/pkg/vm-writer"
)

/*
term:

	integerConstant | stringConstant | keywordConstant | varName |
	varName'['expression']' | subroutineCall | '('expression')' | unaryOp term
*/
func (e *Engine) CompileTerm() {
	fmt.Println("--- CompileTerm ---, token: ", e.tk.Token())

	// e.WriteString("<term>\n")

	tokenType := e.tk.TokenType()

	if tokenType == jackTokenizer.INT_CONST {
		// e.writeIntegerConst()
		e.vmWriter.WritePush(vmWriter.SEG_CONSTANT, e.tk.IntVal())

	} else if tokenType == jackTokenizer.STRING_CONST {
		// e.writeStringConst()
		fmt.Println("	TERM STRING_CONST")
	} else if constVal, isKeywordConst := keywordConstant[e.tk.Keyword()]; isKeywordConst {
		// e.writeKeyword()
		e.vmWriter.WritePush(vmWriter.SEG_CONSTANT, constVal)
		if constVal == 1 {
			e.vmWriter.WriteArithmetic("neg")
		}
	} else if e.tk.Symbol() == "(" {
		// e.writeSymbol()

		e.tk.Advance()
		e.CompileExpression()

		if e.tk.Symbol() != ")" {
			log.Fatal("CompileTerm, (expression) expect a closing ), got", e.tk.Token())
		}
		// e.writeSymbol()

		// eg. let i = i * (-j); we need to reset skipAdvance to false after j
		e.tk.SetSkipAdvance(false)

	} else if _, isUnaryOp := unaryOp[e.tk.Symbol()]; isUnaryOp {
		// e.writeSymbol()
		op := e.tk.Symbol()
		e.tk.Advance()
		e.CompileTerm()

		// ~ or -
		if op == "~" {
			e.vmWriter.WriteArithmetic("not")
		} else if op == "-" {
			e.vmWriter.WriteArithmetic("neg")
		} else {
			log.Fatal("CompileTerm expect unaryOp, got: ", op)
		}

	} else if tokenType == jackTokenizer.IDENTIFIER {
		fmt.Println("==== CompileTerm, identifier: ", e.tk.Identifier())
		prevId := e.tk.Identifier()

		e.tk.Advance()
		if e.tk.Symbol() == "[" {
			// "prevId" is either in subroutineST or classST
			if e.subroutineST.KindOf(prevId) != "" {
				// e.writeIdentifier(prevId, "used", e.subroutineST.KindOf(prevId))
			} else if e.classST.KindOf(prevId) != "" {
				// e.writeIdentifier(prevId, "used", e.classST.KindOf(prevId))
			} else {
				log.Fatal("CompileTerm, identifier before [ should be in one of symbol tables, token: ", prevId)
			}

			// e.writeIdentifier(e.tk.Identifier(), "used", "")

			// e.writeSymbol()

			e.tk.Advance()
			e.CompileExpression()

			if e.tk.Symbol() != "]" {
				log.Fatal("CompileTerm, expect a ']', got: ", e.tk.Token())
			}
			// e.writeSymbol()

			// let sum = sum + a[i];
			e.tk.SetSkipAdvance(false)
		} else if e.tk.Symbol() == "(" {
			// e.writeIdentifier(prevId, "used", symbolTable.SUBROUTINE)

			// e.writeSymbol()

			e.tk.Advance()
			nArgs := e.CompileExpressionList()

			if e.tk.Symbol() != ")" {
				log.Fatal("CompileTerm, expect a ')'")
			}
			// e.writeSymbol()
			e.vmWriter.WriteCall(prevId, nArgs)

		} else if e.tk.Symbol() == "." {
			// "prevId" is either a className or a varName
			prevIsVarName := true
			var classTypeOfVar string
			if e.subroutineST.KindOf(prevId) != "" {
				// e.writeIdentifier(prevId, "used", e.subroutineST.KindOf(prevId))
				classTypeOfVar = e.subroutineST.TypeOf(prevId)
			} else if e.classST.KindOf(prevId) != "" {
				// e.writeIdentifier(prevId, "used", e.classST.KindOf(prevId))
				classTypeOfVar = e.classST.TypeOf(prevId)
			} else {
				// e.writeIdentifier(prevId, "used", symbolTable.CLASS)
				prevIsVarName = false
			}

			// .
			// e.writeSymbol()

			e.tk.Advance()
			if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
				log.Fatal("CompileTerm className|varName (identifier) (expect identifier), got:", e.tk.Token())
			}
			subroutineName := e.tk.Identifier()
			// e.writeIdentifier(e.tk.Identifier(), "used", symbolTable.SUBROUTINE)

			e.tk.Advance()
			if e.tk.Symbol() != "(" {
				log.Fatal("CompileTerm expect '('")
			}
			// e.writeSymbol()

			e.tk.Advance()
			nArgs := e.CompileExpressionList()

			if e.tk.Symbol() != ")" {
				log.Fatal("CompileTerm, expect a ')'")
			}

			if prevIsVarName {
				e.vmWriter.WriteCall(fmt.Sprintf("%s.%s", classTypeOfVar, subroutineName), nArgs)
			} else {
				e.vmWriter.WriteCall(fmt.Sprintf("%s.%s", prevId, subroutineName), nArgs)
			}

			// e.writeSymbol()
		} else {
			// skip advance if the current is varName
			var segment string
			var index int
			if e.subroutineST.KindOf(prevId) != "" {
				// e.writeIdentifier(prevId, "used", e.subroutineST.KindOf(prevId))
				segment = e.vmWriter.KindToSegment(e.subroutineST.KindOf(prevId))
				index = e.subroutineST.IndexOf(prevId)
			} else if e.classST.KindOf(prevId) != "" {
				// e.writeIdentifier(prevId, "used", e.classST.KindOf(prevId))
				segment = e.vmWriter.KindToSegment(e.classST.KindOf(prevId))
				index = e.classST.IndexOf(prevId)
			} else {
				log.Fatal("CompileTerm, this should be a var name")
			}
			e.vmWriter.WritePush(segment, index)
			e.tk.SetSkipAdvance(true)
		}
	} else {
		log.Fatal("CompileTerm, unsupported term, got:", e.tk.Token())
	}

	// e.WriteString("</term>\n")
	fmt.Println("		END CompileTerm")
}

func (e *Engine) isTerm() bool {
	tokenType := e.tk.TokenType()
	_, isKeywordConst := keywordConstant[e.tk.Keyword()]
	_, isUnaryOp := unaryOp[e.tk.Symbol()]
	return tokenType == jackTokenizer.INT_CONST ||
		tokenType == jackTokenizer.STRING_CONST ||
		isKeywordConst ||
		e.tk.Symbol() == "(" ||
		isUnaryOp ||
		tokenType == jackTokenizer.IDENTIFIER
}
