package codeWriter

import (
	"log"
	"os"
)

/*
CodeWriter:
 - Generates assembly code from the parsed VM command:
1. Constructor args: Output file/stream, return: -, function: Opens the output file/stream and gets ready to write into it.
2. WriteArithmetic, args: command(string), return: -, function: Writes to the output file the assembly code that implements the given arithmetic command.
3. WritePushPop, args: command(C_PUSH or C_POP), segment(string), index(int), return: -, function: Writes to the output file the assembly code that implements the given command,
	where command is either C_PUSH or C_POP.
4. Close, args: -, return: -, function: Closes the output file.
*/

var COMMAND_DICT = map[string]string{
	"C_PUSH": "push",
	"C_POP":  "pop",
}

var MEMORY_SEGMENT_DICT = map[string]string{
	"local":    "LCL",
	"argument": "ARG",
	"this":     "THIS",
	"that":     "THAT",
}

type CodeWriter struct {
	File        *os.File
	programName string
	lineCounter int
}

func NewCodeWriter(fileName, programName string) (*CodeWriter, error) {
	writeFile, err := os.Create(fileName)
	return &CodeWriter{File: writeFile, programName: programName}, err
}

func (c *CodeWriter) WriteCmd(cmd string) {
	_, err := c.File.WriteString(cmd)
	check(err)
	c.lineCounter++
	// fmt.Println("now line:", c.lineCounter)
}

func (c *CodeWriter) WriteComment(cmd string) {
	_, err := c.File.WriteString(cmd)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
