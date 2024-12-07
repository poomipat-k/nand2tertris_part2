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

	cw.WriteInit()

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
			cmd := parser.Command()
			if cmdType == "C_POP" || cmdType == "C_PUSH" {
				cw.WritePushPop(cmd, cmdType, parser.Arg1(), parser.Arg2())
			} else if cmdType == "C_LABEL" {
				label := fmt.Sprintf("%s$%s", cw.CurFuncName, parser.Arg1())
				cw.WriteLabel(cmd, label)
			} else if cmdType == "C_IF" {
				label := fmt.Sprintf("%s$%s", cw.CurFuncName, parser.Arg1())
				cw.WriteIf(cmd, label)
			} else if cmdType == "C_GOTO" {
				label := fmt.Sprintf("%s$%s", cw.CurFuncName, parser.Arg1())
				cw.WriteGoto(cmd, label)
			} else if cmdType == "C_FUNCTION" {
				cw.WriteFunction(cmd, parser.Arg1(), parser.Arg2())
			} else if cmdType == "C_RETURN" {
				cw.WriteReturn()
			} else if cmdType == "C_CALL" {
				cw.WriteCall(cmd, parser.Arg1(), parser.Arg2())
			} else {
				cw.WriteArithmetic(cmd)
			}
		}
		cw.WriteComment(fmt.Sprintf("// END file: %s\n", fileName))

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
	splitFileOrDirName := strings.Split(fileOrDirName, ".")
	// .vm file
	if len(splitFileOrDirName) == 2 {
		fileName := splitFileOrDirName[0]
		return []string{path}, fmt.Sprintf("%s/%s.asm", strings.Join(splits[:len(splits)-1], "/"), fileName)
	}
	// directory
	dirName := fileOrDirName
	fileInfo, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	vmFilePaths := []string{}
	for _, f := range fileInfo {
		name := f.Name()
		sp := strings.Split(name, ".")
		if len(sp) == 2 && sp[1] == "vm" {
			vmFilePaths = append(vmFilePaths, fmt.Sprintf("%s/%s", path, name))
		}
	}
	return vmFilePaths, fmt.Sprintf("%s/%s.asm", path, dirName)
}

func getFileName(filepath string) string {
	splits := strings.Split(filepath, "/")
	s := splits[len(splits)-1]
	splits = strings.Split(s, ".")
	return splits[0]
}
