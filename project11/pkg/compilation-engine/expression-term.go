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
	tokenType := e.tk.TokenType()

	if tokenType == jackTokenizer.INT_CONST {
		e.vmWriter.WritePush(vmWriter.SEG_CONSTANT, e.tk.IntVal())

	} else if tokenType == jackTokenizer.STRING_CONST {
		fmt.Println("		@@@@str:", e.tk.StringVal())
		str := e.tk.StringVal()
		sLen := len(str)

		e.vmWriter.WritePush(vmWriter.SEG_CONSTANT, sLen)
		e.vmWriter.WriteCall("String.new", 1)
		for i := 0; i < sLen; i++ {
			e.vmWriter.WritePush(vmWriter.SEG_CONSTANT, int(str[i]))
			e.vmWriter.WriteCall("String.appendChar", 2)
		}

	} else if e.tk.Keyword() == "this" {
		e.vmWriter.WritePush(vmWriter.SEG_POINTER, 0)
	} else if constVal, isKeywordConst := keywordConstant[e.tk.Keyword()]; isKeywordConst {
		e.vmWriter.WritePush(vmWriter.SEG_CONSTANT, constVal)
		if constVal == 1 {
			e.vmWriter.WriteArithmetic("neg")
		}
	} else if e.tk.Symbol() == "(" {
		e.tk.Advance()
		e.CompileExpression()

		if e.tk.Symbol() != ")" {
			log.Fatal("CompileTerm, (expression) expect a closing ), got", e.tk.Token())
		}

		// eg. let i = i * (-j); we need to reset skipAdvance to false after j
		e.tk.SetSkipAdvance(false)

	} else if _, isUnaryOp := unaryOp[e.tk.Symbol()]; isUnaryOp {
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
		prevId := e.tk.Identifier()

		e.tk.Advance()
		if e.tk.Symbol() == "[" {
			// "prevId" is either in subroutineST or classST
			prevKind := e.getKindOfIdentifier(prevId)
			if prevKind == "" {
				log.Fatal("CompileTerm, identifier before [ should be in one of symbol tables, token: ", prevId)
			}
			prevSegment := e.vmWriter.KindToSegment(prevKind)
			prevIndex := e.getIndexOfIdentifier(prevId)

			// push the Array var
			e.vmWriter.WritePush(prevSegment, prevIndex)

			e.tk.Advance()
			e.CompileExpression()

			e.vmWriter.WriteArithmetic("add")

			if e.tk.Symbol() != "]" {
				log.Fatal("CompileTerm, expect a ']', got: ", e.tk.Token())
			}

			e.vmWriter.WritePop(vmWriter.SEG_POINTER, 1)
			e.vmWriter.WritePush(vmWriter.SEG_THAT, 0) // push b[j] to top of stack
			// e.vmWriter.WritePop(vmWriter.SEG_TEMP, 0)

			// let sum = sum + a[i];
			e.tk.SetSkipAdvance(false)
		} else if e.tk.Symbol() == "(" {

			e.tk.Advance()
			nArgs := e.CompileExpressionList()

			if e.tk.Symbol() != ")" {
				log.Fatal("CompileTerm 1, expect a ')'")
			}

			e.vmWriter.WritePush(vmWriter.SEG_POINTER, 0)
			e.vmWriter.WriteCall(fmt.Sprintf("%s.%s", e.className, prevId), nArgs+1)

		} else if e.tk.Symbol() == "." {
			// "prevId" is either a className or a varName
			prevIsClassVarInstance := true
			var classTypeOfVar string
			if e.subroutineST.KindOf(prevId) != "" {
				classTypeOfVar = e.subroutineST.TypeOf(prevId)
			} else if e.classST.KindOf(prevId) != "" {
				classTypeOfVar = e.classST.TypeOf(prevId)
			} else {
				prevIsClassVarInstance = false
			}

			// .

			e.tk.Advance()
			if e.tk.TokenType() != jackTokenizer.IDENTIFIER {
				log.Fatal("CompileTerm className|varName (identifier) (expect identifier), got:", e.tk.Token())
			}
			subroutineName := e.tk.Identifier()

			e.tk.Advance()
			if e.tk.Symbol() != "(" {
				log.Fatal("CompileTerm expect '('")
			}

			e.tk.Advance()

			if prevIsClassVarInstance {
				prevKind := e.getKindOfIdentifier(prevId)
				prevSegment := e.vmWriter.KindToSegment(prevKind)
				prevIndex := e.getIndexOfIdentifier(prevId)
				// push prevId as a first method argument
				e.vmWriter.WritePush(prevSegment, prevIndex)
			}
			nArgs := e.CompileExpressionList()

			if e.tk.Symbol() != ")" {
				log.Fatal("CompileTerm 2, expect a ')'")
			}

			if prevIsClassVarInstance {
				nArgs++
				e.vmWriter.WriteCall(fmt.Sprintf("%s.%s", classTypeOfVar, subroutineName), nArgs)
			} else {
				// prevId is a class
				e.vmWriter.WriteCall(fmt.Sprintf("%s.%s", prevId, subroutineName), nArgs)
			}
		} else {
			// skip advance if the current is varName
			var segment string
			var index int
			if e.subroutineST.KindOf(prevId) != "" {
				segment = e.vmWriter.KindToSegment(e.subroutineST.KindOf(prevId))
				index = e.subroutineST.IndexOf(prevId)
			} else if e.classST.KindOf(prevId) != "" {
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

	fmt.Println("		END CompileTerm")
}

func (e *Engine) isTerm() bool {
	tokenType := e.tk.TokenType()
	_, isKeywordConst := keywordConstant[e.tk.Keyword()]
	_, isUnaryOp := unaryOp[e.tk.Symbol()]
	return tokenType == jackTokenizer.INT_CONST ||
		tokenType == jackTokenizer.STRING_CONST ||
		isKeywordConst ||
		e.tk.Keyword() == "this" ||
		e.tk.Symbol() == "(" ||
		isUnaryOp ||
		tokenType == jackTokenizer.IDENTIFIER
}
