package compilationEngine

import (
	"log"
	"os"

	symbolTable "github.com/poomipat-k/nand2tetris/project11/pkg/symbol-table"
	jackTokenizer "github.com/poomipat-k/nand2tetris/project11/pkg/tokenizer"
)

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
