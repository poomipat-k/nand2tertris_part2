package codeWriter

import (
	"fmt"
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
	File *os.File
}

func NewCodeWriter(fileName string) (*CodeWriter, error) {
	writeFile, err := os.Create(fileName)
	return &CodeWriter{File: writeFile}, err
}

func (c *CodeWriter) WriteArithmetic(cmd string) {
	_, err := c.File.WriteString(fmt.Sprintf("// %s\n", cmd))
	check(err)
}

func (c *CodeWriter) WritePushPop(cmd string, cmdType string, segment string, index int) {
	_, err := c.File.WriteString(fmt.Sprintf("// %s %s %d\n", cmd, segment, index))
	check(err)
	if cmdType == "C_PUSH" {
		c.writePush(segment, index)
	} else {
		c.writePop(segment, index)
	}
}

func (c *CodeWriter) writePush(segment string, index int) {
	var err error
	if segment == "constant" {
		_, err = c.File.WriteString(fmt.Sprintf("@%d\n", index))
		check(err)

		_, err = c.File.WriteString("D=A\n") // D=10
		check(err)
	} else {
		/*
			eg. push local 5
			@5
			D=A
			@local
			A=D+A
			D=M
		*/
		_, err = c.File.WriteString(fmt.Sprintf("@%d\n", index))
		check(err)

		_, err = c.File.WriteString("D=A\n")
		check(err)

		smSym := MEMORY_SEGMENT_DICT[segment]
		_, err = c.File.WriteString(fmt.Sprintf("@%s\n", smSym))
		check(err)

		_, err = c.File.WriteString("A=D+A\n")
		check(err)

		_, err = c.File.WriteString("D=M\n")
		check(err)
	}

	// increment SP
	_, err = c.File.WriteString("@SP\n")
	check(err)

	_, err = c.File.WriteString("AM=M+1\n")
	check(err)

	// Get back to the SP that want to push value to
	_, err = c.File.WriteString("A=A-1\n")
	check(err)

	_, err = c.File.WriteString("M=D\n")
	check(err)
}

func (c *CodeWriter) writePop(segment string, index int) {
	// invalid segment
	if segment == "constant" {
		return
	}
	var err error
	smSym := MEMORY_SEGMENT_DICT[segment]

	if index == 0 {
		// decrement SP
		_, err = c.File.WriteString("@SP\n")
		check(err)
		_, err = c.File.WriteString("AM=M-1\n")
		check(err)
		_, err = c.File.WriteString("D=M\n")
		check(err)

		// save D to desired ram position (segment + index)
		_, err = c.File.WriteString(fmt.Sprintf("@%s\n", smSym))
		check(err)
		_, err = c.File.WriteString("A=M\n")
		check(err)
		_, err = c.File.WriteString("M=D\n")
		check(err)
	} else {
		// segment(M) = segment + index
		_, err = c.File.WriteString(fmt.Sprintf("@%d\n", index))
		check(err)
		_, err = c.File.WriteString("D=A\n")
		check(err)
		_, err = c.File.WriteString(fmt.Sprintf("@%s\n", smSym))
		check(err)
		_, err = c.File.WriteString("M=D+M\n")
		check(err)

		// decrement SP
		_, err = c.File.WriteString("@SP\n")
		check(err)
		_, err = c.File.WriteString("AM=M-1\n")
		check(err)
		_, err = c.File.WriteString("D=M\n")
		check(err)

		// save D to desired ram position (segment + index)
		_, err = c.File.WriteString(fmt.Sprintf("@%s\n", smSym))
		check(err)
		_, err = c.File.WriteString("A=M\n")
		check(err)
		_, err = c.File.WriteString("M=D\n")
		check(err)

		// revert segment back segment + index -> segment
		_, err = c.File.WriteString(fmt.Sprintf("@%d\n", index))
		check(err)
		_, err = c.File.WriteString("D=A\n")
		check(err)
		_, err = c.File.WriteString(fmt.Sprintf("@%s\n", smSym))
		check(err)
		_, err = c.File.WriteString("M=M-D\n")
		check(err)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
