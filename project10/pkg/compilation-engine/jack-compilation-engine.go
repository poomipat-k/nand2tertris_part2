package compilationEngine

import (
	"log"
	"os"

	jackTokenizer "github.com/poomipat-k/nand2tetris/project10/pkg/tokenizer"
)

type Engine struct {
	tk          *jackTokenizer.Tokenizer
	OutFile     *os.File
	skipTermTag bool
}

func NewEngine(tokenizer *jackTokenizer.Tokenizer, outputPath string) (*Engine, error) {
	writeFile, err := os.Create(outputPath)
	return &Engine{tk: tokenizer, OutFile: writeFile}, err
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
