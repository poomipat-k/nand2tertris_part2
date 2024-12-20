package jackTokenizer

import (
	"bufio"
	"fmt"
	"math"
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
	scanner     *bufio.Scanner
	File        *os.File
	endOfFile   bool
	token       string
	currentLine string
	lineCursor  int
	tokenType   string
	keyWord     string
	symbol      string
	identifier  string
	intVal      int
	stringVal   string
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

func (t *Tokenizer) Advance() {
	/*
		ignore the following
		- whitespace
		- inline comment
		- multiple lines comment
	*/
	var line string
	multipleLnCmtOn := false
	insideDoubleQuote := false
	t.token = ""
	t.tokenType = ""

	for t.token == "" {
		hasMoreLine := t.scanner.Scan()
		if !hasMoreLine {
			t.endOfFile = true
			return
		}

		line = strings.TrimSpace(t.scanner.Text()) // trim left and right white space
		// skip empty line
		if line == "" {
			continue
		}
		t.currentLine = line
		t.lineCursor = 0
		lineLn := len(line)
		i := t.lineCursor
		startTokenIndex := -1

		for i < lineLn {
			upperBoundInd := int(math.Min(float64(i+2), float64(lineLn)))
			slidingTwoChars := line[i:upperBoundInd]

			if multipleLnCmtOn {
				if slidingTwoChars == "*/" {
					multipleLnCmtOn = false
					i += 2 // set cursor to the next char to process
				} else {
					i++
				}
				t.lineCursor = i
				continue
			}
			if insideDoubleQuote {
				if line[i] == '"' {
					t.token = line[startTokenIndex:i]
					t.tokenType = "STRING_CONST"
					insideDoubleQuote = false
					i++
					t.lineCursor = i
					return
				}
				i++
				t.lineCursor = i
				continue
			}
			if slidingTwoChars == "//" { // discard the remaining of this line
				break
			}
			if slidingTwoChars == "/*" {
				multipleLnCmtOn = true
				i += 2 // set cursor to the next char to process
				continue
			}
			if isSpace(rune(line[i])) {
				i++
				continue
			}

			fmt.Println("===Start processing token")

			t.lineCursor = i
			// find token
			fmt.Println(line)
			fmt.Println("i: ", i, ", line[i]:", string(line[i]))
			if startTokenIndex == -1 {
				startTokenIndex = i
			}

			nextChar := line[i+1]
			if line[i] == '"' {
				startTokenIndex = i + 1
				insideDoubleQuote = true
			} else if isSymbol(rune(line[i])) {
				t.token = string(line[i])
				t.tokenType = "SYMBOL"
				t.symbol = t.token
				return
			} else if isSpace(rune(nextChar)) || isSymbol(rune(nextChar)) {
				t.token = "TBD"
				return
			}
			i++
		}
	}
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
