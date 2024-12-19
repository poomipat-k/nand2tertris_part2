package jackTokenizer

import (
	"bufio"
	"math"
	"os"
	"strings"
	"unicode"
)

type Tokenizer struct {
	scanner     *bufio.Scanner
	File        *os.File
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

func (t *Tokenizer) Advance() error {
	return nil
}

func (t *Tokenizer) GetLineCursor() int {
	return t.lineCursor
}

func (t *Tokenizer) GetLine() string {
	return t.currentLine
}

func (t *Tokenizer) HasMoreTokens() bool {
	/*
		ignore the following
		- whitespace
		- inline comment
		- multiple lines comment
	*/
	var line string
	multipleLnCmtOn := false

	for t.scanner.Scan() {
		line = strings.TrimSpace(t.scanner.Text()) // trim left and right white space
		// skip empty line
		if line == "" {
			continue
		}
		t.currentLine = line
		t.lineCursor = 0
		lineLn := len(line)
		i := 0
		for i < lineLn {
			upperBoundInd := int(math.Min(float64(i+2), float64(lineLn)))
			slidingTwoChars := line[i:upperBoundInd]
			// fmt.Println("==sliding: ", slidingTwoChars)
			if multipleLnCmtOn {
				if slidingTwoChars == "*/" {
					multipleLnCmtOn = false
					i += 2 // set cursor to the next char to process
					t.lineCursor = i
				} else {
					i++
				}
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
			if unicode.IsSpace(rune(line[i])) {
				i++
				continue
			}
			// fmt.Println(line[i:])
			t.lineCursor = i
			return true
		}
		// fmt.Println(line)
	}
	return false
}
