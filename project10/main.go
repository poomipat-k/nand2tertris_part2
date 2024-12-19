package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	jackTokenizer "github.com/poomipat-k/nand2tetris/project10/pkg"
)

func main() {
	inputPath := os.Args[1]
	if len(inputPath) == 0 {
		log.Fatal("filePath or dirPath is required")
	}

	paths, outPaths := processInputPath(inputPath)
	for i := 0; i < len(paths); i++ {
		fmt.Println(paths[i])
		fmt.Println(outPaths[i])
		tokenAnalyzer(paths[i])
		fmt.Println("============")
	}
}

// V.0 for unit testing tokenAnalyzer function
func tokenAnalyzer(filePath string) {
	tokenizer, err := jackTokenizer.NewTokenizer(filePath)
	check(err)
	err = tokenizer.Advance() // get the first token
	check(err)

	for tokenizer.HasMoreTokens() {
		/*
			tokenType = type of the current token
			print "<" + tokenType + ">"
			print the current token
			print "</" + tokenType + ">"
			print newLine
			tokenizer.Advance()
		*/
		fmt.Println("line:", tokenizer.GetLine())
		fmt.Println("cursor:", tokenizer.GetLineCursor())
		err = tokenizer.Advance()
		check(err)
	}
}

// return []path, []outfilePaths
func processInputPath(path string) ([]string, []string) {
	splitsSlash := strings.Split(path, "/")
	fileOrDirName := splitsSlash[len(splitsSlash)-1]
	splitFileOrDirName := strings.Split(fileOrDirName, ".")
	// one .jack file
	if len(splitFileOrDirName) == 2 {
		fileName := splitFileOrDirName[0]
		return []string{path}, []string{fmt.Sprintf("%s/%s.xml", strings.Join(splitsSlash[:len(splitsSlash)-1], "/"), fileName)}
	}
	// directory
	fileInfo, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	jackFilePaths := []string{}
	outFilePaths := []string{}
	for _, f := range fileInfo {
		name := f.Name()
		sp := strings.Split(name, ".")
		if len(sp) == 2 && sp[1] == "jack" {
			jackFilePaths = append(jackFilePaths, fmt.Sprintf("%s/%s", path, name))
			outFilePaths = append(outFilePaths, fmt.Sprintf("%s/%s_generated.xml", path, sp[0]))
		}
	}
	return jackFilePaths, outFilePaths
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
