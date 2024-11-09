package main

import (
	"fmt"
	"log"
	"os"

	codeWriter "github.com/poomipat-k/nand2tetris2/project7/pkg/code-writer"
	vmParser "github.com/poomipat-k/nand2tetris2/project7/pkg/vm-parser"
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
 - Constructs a Parser to handle the input file
 - Constructs a CodeWriter to handle the output file
 - Marches through the input file, parsing each line and generating code from it


Parser:
 - Handles the parsing of a single .vm file
 - Reads a VM command, parses the command into its lexical components, and provides convenient access to these components
 - Ignores all white space and comments

 1. Constructor args: input file/stream, return: -, function: Opens the input file/stream and get ready to parse it
 2. hasMoreCommands args: -, return: boolean, function: Are there more commands in the input?
 3. advance args: -, return: -, function: Reads the next command from the input and makes it the current command.
	Should be called only if hasMoreCommands() is true. Initially there is no current command.
 4. commandType: args: -, returns: {C_ARITHMETIC, C_PUSH, C_POP, C_LABEL, C_GOTO, C_IF, C_FUNCTION, C_RETURN, C_CALL}
	function: returns a constant representing the type of the current command. C_ARITHMETIC is returned for all the arithmetic/logical commands
 5. arg1 args: -, return: string function: Returns the first argument of the current command. In the case of C_ARITHMETIC, the command itself(add, sub, etc.) is returned.
	Should not be called if the current command is C_RETURN
 6. arg2 args: -, return: int function: Returns the second argument of the current command. Should be called only if the current command is C_PUSH, C_POP, C_FUNCTION or C_CALL


CodeWriter:
 - Generates assembly code from the parsed VM command:
1. Constructor args: Output file/stream, return: -, function: Opens the output file/stream and gets ready to write into it.
2. WriteArithmetic, args: command(string), return: -, function: Writes to the output file the assembly code that implements the given arithmetic command.
3. WritePushPop, args: command(C_PUSH or C_POP), segment(string), index(int), return: -, function: Writes to the output file the assembly code that implements the given command,
	where command is either C_PUSH or C_POP.
4. Close, args: -, return: -, function: Closes the output file.
*/

func main() {
	fileName := os.Args[1]
	outFileName := os.Args[2]
	fmt.Printf("filename %s, outFileName: %s\n", fileName, outFileName)

	parser, err := vmParser.NewParser(fileName)
	check(err)
	defer parser.File.Close()

	cw, err := codeWriter.NewCodeWriter(outFileName)
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
			cw.WritePushPop(parser.Command(), parser.Arg1(), parser.Arg2())
		} else {
			cw.WriteArithmetic(parser.Command())
		}
		// fmt.Println("=====")
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
