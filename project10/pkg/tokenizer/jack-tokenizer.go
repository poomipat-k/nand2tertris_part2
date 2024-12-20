package jackTokenizer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// tokenType = [KEYWORD, SYMBOL, IDENTIFIER, INT_CONST, STRING_CONST]
const KEYWORD = "KEYWORD"
const SYMBOL = "SYMBOL"
const IDENTIFIER = "IDENTIFIER"
const INT_CONST = "INT_CONST"
const STRING_CONST = "STRING_CONST"

type Tokenizer struct {
	scanner                *bufio.Scanner
	File                   *os.File
	endOfFile              bool
	currentLine            string
	lineCursor             int
	insideMultilineComment bool
	insideDoubleQuote      bool
	// essential fields
	token      string
	tokenType  string
	keyword    string
	symbol     string
	identifier string
	intVal     int
	stringVal  string
}

var symbolSet = map[rune]bool{
	'{': true,
	'}': true,
	'(': true,
	')': true,
	'[': true,
	']': true,
	'.': true,
	',': true,
	';': true,
	'+': true,
	'-': true,
	'*': true,
	'/': true,
	'&': true,
	'|': true,
	'<': true,
	'>': true,
	'=': true,
	'~': true,
}

func NewTokenizer(filePath string) (*Tokenizer, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	tk := &Tokenizer{
		scanner: scanner,
		File:    f,
	}
	return tk, nil
}

func (t *Tokenizer) Advance2() {
	// scan until found non-comment/whitespace

	// reset previous token states
	t.resetTokenStates()
	// scan new line and skip empty lines
	for {
		if !t.scanner.Scan() {
			t.endOfFile = true
			return
		}
		line := strings.TrimSpace(t.scanner.Text())
		// skip empty line
		if line == "" {
			continue
		}
		t.currentLine = line
		t.lineCursor = 0
	}
	// scan new line until reach the beginning of the new token
}

func (t *Tokenizer) resetTokenStates() {
	t.token = ""
	t.tokenType = ""
	t.symbol = ""
	t.identifier = ""
	t.intVal = 0
	t.stringVal = ""
}

func (t *Tokenizer) Advance() {
	/*
		ignore the following
		- whitespace
		- inline comment
		- multiple lines comment
	*/
	t.resetTokenStates()
	for {
		if t.currentLine == "" || t.lineCursor >= len(t.currentLine) {
			hasMoreLine := t.scanner.Scan()
			if !hasMoreLine {
				t.endOfFile = true
				return
			}
			line := strings.TrimSpace(t.scanner.Text()) // trim left and right white space
			fmt.Println("line:", line)
			t.currentLine = line
			t.lineCursor = 0
			// skip empty line
			if t.currentLine == "" {
				continue
			}
		}
		startTokenIndex := -1

		for t.lineCursor < len(t.currentLine) {
			upperBoundInd := min(t.lineCursor+2, len(t.currentLine))

			slidingTwoChars := t.currentLine[t.lineCursor:upperBoundInd]

			if t.insideMultilineComment {
				if slidingTwoChars == "*/" {
					t.insideMultilineComment = false
					t.lineCursor += 2 // set cursor to the next char to process
				} else {
					t.lineCursor++
				}
				continue
			}
			if t.insideDoubleQuote {
				if t.currentLine[t.lineCursor] == '"' {
					t.token = t.currentLine[startTokenIndex:t.lineCursor]
					t.tokenType = STRING_CONST
					t.insideDoubleQuote = false
					t.lineCursor++
					return
				}
				t.lineCursor++
				continue
			}
			if slidingTwoChars == "//" { // discard the remaining of this line and go scan a new line
				t.currentLine = ""
				t.lineCursor = 0
				break
			}
			if slidingTwoChars == "/*" {
				t.insideMultilineComment = true
				t.lineCursor += 2 // set cursor to the next char to process
				continue
			}
			if isSpace(rune(t.currentLine[t.lineCursor])) {
				t.lineCursor++
				continue
			}

			fmt.Println("===Start processing token", string(t.currentLine[t.lineCursor]))
			// // find token
			// if startTokenIndex == -1 {
			// 	startTokenIndex = i
			// }

			// nextChar := line[t.lineCursor+1]
			// if line[t.lineCursor] == '"' {
			// 	startTokenIndex = i + 1
			// 	t.insideDoubleQuote = true
			// } else if isSymbol(rune(line[i])) {
			// 	t.token = string(line[i])
			// 	t.tokenType = "SYMBOL"
			// 	t.symbol = t.token
			// 	return
			// } else if isSpace(rune(nextChar)) || isSymbol(rune(nextChar)) {
			// 	t.token = "TBD"
			// 	return
			// }
			t.lineCursor++
			return
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (t *Tokenizer) GetLineCursor() int {
	return t.lineCursor
}

func (t *Tokenizer) GetLine() string {
	return t.currentLine
}

func (t *Tokenizer) HasMoreTokens() bool {
	return t.endOfFile
}

func isSymbol(c rune) bool {
	return symbolSet[c]
}

func isSpace(char rune) bool {
	return unicode.IsSpace(char)
}
