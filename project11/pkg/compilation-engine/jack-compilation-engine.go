package compilationEngine

import (
	"log"
	"os"

	symbolTable "github.com/poomipat-k/nand2tetris/project11/pkg/symbol-table"
	jackTokenizer "github.com/poomipat-k/nand2tetris/project11/pkg/tokenizer"
)

// const (
// 	VAR        = "VAR"
// 	ARG        = "ARG"
// 	STATIC     = "STATIC"
// 	FIELD      = "FIELD"
// 	CLASS      = "CLASS"
// 	SUBROUTINE = "SUBROUTINE"
// )

type Engine struct {
	OutFile      *os.File
	tk           *jackTokenizer.Tokenizer
	classST      *symbolTable.SymbolTable
	subroutineST *symbolTable.SymbolTable
}

func NewEngine(tokenizer *jackTokenizer.Tokenizer, outputPath string) (*Engine, error) {
	writeFile, err := os.Create(outputPath)
	cST := symbolTable.NewSymbolTable()
	sST := symbolTable.NewSymbolTable()
	return &Engine{tk: tokenizer, OutFile: writeFile, classST: cST, subroutineST: sST}, err
}

func (e *Engine) WriteString(s string) {
	_, err := e.OutFile.WriteString(s)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) getKindOfIdentifier(name string) string {
	var kind string
	if e.subroutineST.KindOf(e.tk.Identifier()) != "" {
		kind = e.subroutineST.KindOf(e.tk.Identifier())
	} else if e.classST.KindOf(e.tk.Identifier()) != "" {
		kind = e.classST.KindOf(e.tk.Identifier())
	} else {
		log.Fatal("getKindOfIdentifier, not found in symbol tables, name: ", e.tk.Identifier())
	}
	return kind
}
