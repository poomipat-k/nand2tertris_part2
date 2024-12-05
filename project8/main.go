package main

import (
	"fmt"
	"log"

	"os"
	"strings"

	codeWriter "github.com/poomipat-k/nand2tetris/project8/pkg/code-writer"
	vmParser "github.com/poomipat-k/nand2tetris/project8/pkg/vm-parser"
)

func main() {
	inputPath := os.Args[1]
	if len(inputPath) == 0 {
		log.Fatal("filePath is required")
	}
	filePaths, outPath := processInput(inputPath)

	fmt.Printf("input: %s, outPath: %s, len: %d\n", inputPath, outPath, len(filePaths))

	cw, err := codeWriter.NewCodeWriter(outPath)
	check(err)
	defer cw.File.Close()

	for _, fPath := range filePaths {
		fmt.Println("==fPath:", fPath)
		parser, err := vmParser.NewParser(fPath)
		check(err)

		fileName := getFileName(fPath)
		cw.SetFileName(fileName)
		for parser.HasMoreCommands() {
			valid, err := parser.Advance()
			check(err)

			if !valid {
				continue
			}
			cmdType := parser.CommandType()
			if cmdType == "C_POP" || cmdType == "C_PUSH" {
				cw.WritePushPop(parser.Command(), parser.CommandType(), parser.Arg1(), parser.Arg2())
			} else if cmdType == "C_LABEL" {
				cw.WriteLabel(parser.Command(), parser.Arg1())
			} else if cmdType == "C_IF" {
				cw.WriteIf(parser.Command(), parser.Arg1())
			} else {
				cw.WriteArithmetic(parser.Command())
			}
		}

		parser.File.Close()
	}

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// return []path, outfilePath
func processInput(path string) ([]string, string) {
	splits := strings.Split(path, "/")
	fileOrDirName := splits[len(splits)-1]
	fmt.Println("fileOrDirName:", fileOrDirName)
	splitFileOrDirName := strings.Split(fileOrDirName, ".")
	// .vm file
	if len(splitFileOrDirName) == 2 {
		fileName := splitFileOrDirName[0]
		return []string{path}, fmt.Sprintf("%s/%s.asm", strings.Join(splits[:len(splits)-1], "/"), fileName)
	}
	// directory

	return []string{}, ""
}

func getFileName(filepath string) string {
	splits := strings.Split(filepath, "/")
	s := splits[len(splits)-1]
	splits = strings.Split(s, ".")
	return splits[0]
}
