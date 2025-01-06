package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	compilationEngine "github.com/poomipat-k/nand2tetris/project11/pkg/compilation-engine"
	jackTokenizer "github.com/poomipat-k/nand2tetris/project11/pkg/tokenizer"
)

func main() {
	inputPath := os.Args[1]
	if len(inputPath) == 0 {
		log.Fatal("filePath or dirPath is required")
	}

	srcPaths, outPaths := processInputPath(inputPath)

	for i := 0; i < len(srcPaths); i++ {
		fmt.Println(srcPaths[i])
		fmt.Println(outPaths[i])

		compile(srcPaths[i], outPaths[i])
		fmt.Println("============")
	}
}

func compile(srcFilePath string, outputPath string) {
	tknz, err := jackTokenizer.NewTokenizer(srcFilePath)
	check(err)
	defer tknz.File.Close()

	engine := compilationEngine.NewEngine(tknz, outputPath)
	defer engine.Close()

	tknz.Advance()
	if tknz.Token() != "class" {
		log.Fatal("First token should be a 'class' keyword")
	}

	engine.CompileClass()
}

// return []path, []outfilePaths
func processInputPath(path string) ([]string, []string) {
	splitsSlash := strings.Split(path, "/")
	fileOrDirName := splitsSlash[len(splitsSlash)-1]
	splitFileOrDirName := strings.Split(fileOrDirName, ".")
	// one .jack file
	if len(splitFileOrDirName) == 2 {
		fileName := splitFileOrDirName[0]
		return []string{path}, []string{fmt.Sprintf("%s/%s_generated.vm", strings.Join(splitsSlash[:len(splitsSlash)-1], "/"), fileName)}
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
			outFilePaths = append(outFilePaths, fmt.Sprintf("%s/%s_generated.vm", path, sp[0]))
		}
	}
	return jackFilePaths, outFilePaths
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
