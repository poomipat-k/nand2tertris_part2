package vmWriter

import (
	"fmt"
	"log"
	"os"
)

const (
	SEG_CONSTANT = "constant"
	SEG_ARG      = "argument"
)

type VMWriter struct {
	outFile *os.File
}

func NewVMWriter(outputPath string) *VMWriter {
	writeFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal("NewVMWriter, err: ", err)
	}
	return &VMWriter{outFile: writeFile}
}

func (v *VMWriter) WritePush(segment string, index int) {
	v.writeString(fmt.Sprintf("push %s %d\n", segment, index))
}

func (v *VMWriter) WritePop(segment string, index int) {
	v.writeString(fmt.Sprintf("pop %s %d\n", segment, index))
}

/*
ADD, SUB, NEG, EQ, GT, LT, AND, OR, NOT
*/
func (v *VMWriter) WriteArithmetic(cmd string) {
	v.writeString(fmt.Sprintf("%s\n", cmd))
}

func (v *VMWriter) WriteLabel(label string) {
	v.writeString(fmt.Sprintf("label %s\n", label))
}

func (v *VMWriter) WriteGoto(label string) {
	v.writeString(fmt.Sprintf("goto %s\n", label))
}

func (v *VMWriter) WriteIf(label string) {
	v.writeString(fmt.Sprintf("if-goto %s\n", label))
}

func (v *VMWriter) WriteCall(name string, nArgs int) {
	v.writeString(fmt.Sprintf("call %s %d\n", name, nArgs))
}

func (v *VMWriter) WriteFunction(name string, nVars int) {
	v.writeString(fmt.Sprintf("function %s %d\n", name, nVars))
}

func (v *VMWriter) WriteReturn() {
	v.writeString("return\n")
}

func (v *VMWriter) Close() {
	v.outFile.Close()
}

func (v *VMWriter) writeString(s string) {
	_, err := v.outFile.WriteString(s)
	if err != nil {
		log.Fatal("WriteString, err: ", err)
	}
}
