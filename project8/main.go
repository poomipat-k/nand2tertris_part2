package main

import (
	"fmt"
	"log"

	"os"
	"strings"

	codeWriter "github.com/poomipat-k/nand2tetris/project8/pkg/code-writer"
	vmParser "github.com/poomipat-k/nand2tetris/project8/pkg/vm-parser"
)

/*
Proposed design:
 - Parser: parses each VM command into its lexical elements
 - CodeWriter: writes the assembly code that implements the parsed command
 - Main: drives the process(VMTranslator)


Main(VMTranslator)
Input: filename.vm
Output: fileName.asm


Main logic:
 - Input fileName.vm or directoryName
 - Constructs a Parser to handle the input file
 - Constructs a CodeWriter to handle the output file
 - Marches through the input file, parsing each line and generating code from it

*/

func main() {
	fileName := os.Args[1]
	outFileName := os.Args[2]
	fmt.Printf("input: %s, outFileName: %s\n", fileName, outFileName)
	splits := strings.Split(fileName, "/")
	programName := splits[len(splits)-1]
	programName = programName[:len(programName)-3]

	parser, err := vmParser.NewParser(fileName)
	check(err)

	defer parser.File.Close()

	cw, err := codeWriter.NewCodeWriter(outFileName, programName)
	check(err)

	defer cw.File.Close()

	for parser.HasMoreCommands() {
		valid, err := parser.Advance()
		check(err)

		if !valid {
			continue
		}
		cmdType := parser.CommandType()
		if cmdType == "C_POP" || cmdType == "C_PUSH" {
			cw.WritePushPop(parser.Command(), parser.CommandType(), parser.Arg1(), parser.Arg2())
		} else {
			cw.WriteArithmetic(parser.Command())
		}
	}

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
