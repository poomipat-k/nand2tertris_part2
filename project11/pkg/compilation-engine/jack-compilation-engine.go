package compilationEngine

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	symbolTable "github.com/poomipat-k/nand2tetris/project11/pkg/symbol-table"
	jackTokenizer "github.com/poomipat-k/nand2tetris/project11/pkg/tokenizer"
	vmWriter "github.com/poomipat-k/nand2tetris/project11/pkg/vm-writer"
)

const (
	CONSTRUCTOR = "constructor"
	METHOD      = "method"
	FUNCTION    = "function"
)

type Engine struct {
	tk             *jackTokenizer.Tokenizer
	classST        *symbolTable.SymbolTable
	subroutineST   *symbolTable.SymbolTable
	vmWriter       *vmWriter.VMWriter
	className      string
	subroutineName string
	subroutineType string // constructor, method, function
}

func NewEngine(tokenizer *jackTokenizer.Tokenizer, outputPath string) *Engine {
	cST := symbolTable.NewSymbolTable()
	sST := symbolTable.NewSymbolTable()
	vmWriter := vmWriter.NewVMWriter(outputPath)
	return &Engine{tk: tokenizer, classST: cST, subroutineST: sST, vmWriter: vmWriter}
}

// func (e *Engine) WriteString(s string) {
// 	_, err := e.OutFile.WriteString(s)
// 	check(err)
// }

// func check(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func (e *Engine) Close() {
	e.vmWriter.Close()
}

func (e *Engine) getKindOfIdentifier(name string) string {
	var kind string
	if e.subroutineST.KindOf(name) != "" {
		kind = e.subroutineST.KindOf(name)
	} else if e.classST.KindOf(name) != "" {
		kind = e.classST.KindOf(name)
	} else {
		log.Fatal("getKindOfIdentifier, not found in symbol tables, name: ", name)
	}
	return kind
}

func (e *Engine) getIndexOfIdentifier(name string) int {
	var index int
	if e.subroutineST.IndexOf(name) != -1 {
		index = e.subroutineST.IndexOf(name)
	} else if e.classST.IndexOf(name) != -1 {
		index = e.classST.IndexOf(name)
	} else {
		log.Fatal("getIndexOfIdentifier, not found in symbol tables, name: ", name)
	}
	return index
}
func generateLabel(className string) string {
	uqId := uuid.New()
	idStr := strings.ReplaceAll(uqId.String(), "-", "_")
	return fmt.Sprintf("%s_%s", className, idStr)
}
