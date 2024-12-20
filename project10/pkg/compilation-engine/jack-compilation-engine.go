package compilationEngine

import (
	"log"
	"os"
)

type Engine struct {
	File *os.File
}

func NewEngine(filePath string) (*Engine, error) {
	writeFile, err := os.Create(filePath)
	return &Engine{File: writeFile}, err
}

func (e *Engine) WriteString(s string) {
	_, err := e.File.WriteString(s)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
