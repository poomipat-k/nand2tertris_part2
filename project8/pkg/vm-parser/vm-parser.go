package vmParser

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
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


Commands:
push segment i
pop segment i

add, sub, neg, eq, gt, lt, and, or, not
*/

type VMParser struct {
	scanner      *bufio.Scanner
	File         *os.File
	_commandType string
	_command     string
	_arg1        string
	_arg2        int
}

func NewParser(fileName string) (*VMParser, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	parser := &VMParser{
		scanner: scanner,
		File:    f,
	}
	return parser, nil
}

func (p *VMParser) HasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *VMParser) Advance() (bool, error) {
	line := p.scanner.Text()
	cmd := getCommands(line)
	if len(cmd) == 0 {
		return false, nil
	}
	cl := len(cmd)
	// arithmetic
	first := strings.ToLower(cmd[0])
	if cl == 1 {
		if first == "return" {
			p._commandType = "C_RETURN"
			p._command = cmd[0]
			p._arg1 = ""
			p._arg2 = -1
			return true, nil
		}
		p._commandType = "C_ARITHMETIC"
		p._command = cmd[0]
		p._arg1 = ""
		p._arg2 = -1
		return true, nil
	}

	if cl == 2 {
		if first == "label" {
			p._command = first
			p._commandType = "C_LABEL"
			p._arg1 = cmd[1]
		} else if first == "if-goto" {
			p._command = first
			p._commandType = "C_IF"
			p._arg1 = cmd[1]
		} else if first == "goto" {
			p._command = first
			p._commandType = "C_GOTO"
			p._arg1 = cmd[1]
		}
		p._arg2 = -1
		return true, nil
	}

	if cl == 3 {
		if first == "push" {
			p._commandType = "C_PUSH"
		} else if first == "pop" {
			p._commandType = "C_POP"
		} else if first == "function" {
			p._commandType = "C_FUNCTION"
		}
		p._command = first
		p._arg1 = cmd[1]
		v, err := strconv.Atoi(cmd[2])
		if err != nil {
			log.Fatal(err)
		}
		p._arg2 = v

	} else {
		log.Fatal("invalid VM instruction")
	}
	return true, nil
}

func (p *VMParser) CommandType() string {
	return p._commandType
}
func (p *VMParser) Command() string {
	return p._command
}
func (p *VMParser) Arg1() string {
	return p._arg1
}
func (p *VMParser) Arg2() int {
	return p._arg2
}

func getCommands(line string) []string {
	splits := strings.Split(line, "//")
	if splits[0] == "" {
		return []string{}
	}

	str := strings.TrimSpace(splits[0])
	return strings.Split(str, " ")
}
