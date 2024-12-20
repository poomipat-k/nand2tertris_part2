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
	startTokenIndex        int
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

func (t *Tokenizer) resetTokenStates() {
	t.token = ""
	t.tokenType = ""
	t.symbol = ""
	t.identifier = ""
	t.intVal = 0
	t.stringVal = ""
	t.startTokenIndex = -1
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
			// fmt.Println("line:", line, "len:", len(line))
			t.currentLine = line
			t.lineCursor = 0
			// skip empty line
			if t.currentLine == "" {
				continue
			}
		}
		// proceed char by char
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
					t.token = t.currentLine[t.startTokenIndex:t.lineCursor]
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

			// fmt.Println("===Start processing token", string(t.currentLine[t.lineCursor]))
			// find token
			if t.startTokenIndex == -1 {
				t.startTokenIndex = t.lineCursor
			}

			if t.currentLine[t.lineCursor] == '"' {
				t.startTokenIndex = t.lineCursor + 1
				t.insideDoubleQuote = true
				t.lineCursor++
				continue
			}
			if isSymbol(rune(t.currentLine[t.lineCursor])) {
				t.token = string(t.currentLine[t.lineCursor])
				t.tokenType = SYMBOL
				t.symbol = t.token
				t.lineCursor++
				fmt.Println("	symbol:", t.token)
				return
			}
			nextChar := t.currentLine[t.lineCursor+1]
			if isSpace(rune(nextChar)) || isSymbol(rune(nextChar)) {
				// find tokenType between [keyword, identifier, intVal]
				word := t.currentLine[t.startTokenIndex : t.lineCursor+1]
				t.token = word
				fmt.Println("	word:", t.token)
				t.lineCursor++
				return
			}
			t.lineCursor++
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
