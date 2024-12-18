package jackTokenizer

import "os"

type Tokenizer struct {
	File *os.File
}

func NewTokenizer(fileName string) (*Tokenizer, error) {
	writeFile, err := os.Create(fileName)
	return &Tokenizer{File: writeFile}, err
}
