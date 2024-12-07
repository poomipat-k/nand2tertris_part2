package codeWriter

import (
	"fmt"
)

func (c *CodeWriter) WriteFunction(cmd string, functionName string, nLocalVars int) {
	c.WriteComment(fmt.Sprintf("// %s %s %d\n", cmd, functionName, nLocalVars))
	c.WriteNonCmd(fmt.Sprintf("(%s)\n", functionName))
	for i := 0; i < nLocalVars; i++ {
		c.writePush("constant", 0)
	}
	c.CurFuncName = functionName
	c.curFuncCallCounter = 1
}

func (c *CodeWriter) WriteReturn() {
	/*
		R13 for write push/pop arithmetic to use
		R14 is endFrame
		R15 is retAddr
	*/
	c.WriteComment("// return\n")
	// endFrame = LCL
	c.WriteCmd("@LCL\n")
	c.WriteCmd("D=M\n")
	c.WriteCmd("@R14\n")
	c.WriteCmd("M=D\n")
	// retAddr = *(endFrame - 5)
	c.WriteCmd("@5\n")
	c.WriteCmd("D=A\n")
	c.WriteCmd("@R14\n")
	c.WriteCmd("A=M-D\n")
	c.WriteCmd("D=M\n")
	c.WriteCmd("@R15\n")
	c.WriteCmd("M=D\n")
	// *ARG = pop()
	c.writePop("argument", 0)
	// SP = ARG + 1
	c.WriteCmd("@ARG\n")
	c.WriteCmd("D=M+1\n")
	c.WriteCmd("@SP\n")
	c.WriteCmd("M=D\n")
	// THAT = *(endFrame - 1)
	c.getEndFrameMinusXToDRegister(1)
	c.WriteCmd("@THAT\n")
	c.WriteCmd("M=D\n")
	// THIS = *(endFrame - 2)
	c.getEndFrameMinusXToDRegister(2)
	c.WriteCmd("@THIS\n")
	c.WriteCmd("M=D\n")
	// ARG = *(endFrame - 3)
	c.getEndFrameMinusXToDRegister(3)
	c.WriteCmd("@ARG\n")
	c.WriteCmd("M=D\n")
	// LCL = *(endFrame - 4)
	c.getEndFrameMinusXToDRegister(4)
	c.WriteCmd("@LCL\n")
	c.WriteCmd("M=D\n")
	// goto retAddr
	c.WriteCmd("@R15\n")
	c.WriteCmd("A=M\n")
	c.WriteCmd("0;JMP\n")
}

func (c *CodeWriter) WriteCall(cmd string, functionName string, nArgs int) {
	c.WriteComment(fmt.Sprintf("// %s %s %d\n", cmd, functionName, nArgs))
	// generate returnAddress label

	returnLabel := fmt.Sprintf("%s$ret.%d", c.CurFuncName, c.curFuncCallCounter)
	c.curFuncCallCounter++

	// push returnAddress (using label create below)
	c.WriteCmd(fmt.Sprintf("@%s\n", returnLabel))
	c.WriteCmd("D=A\n")
	c.pushDToStack()

	// push LCL
	c.WriteCmd("@LCL\n")
	c.WriteCmd("D=M\n")
	c.pushDToStack()
	// push ARG
	c.WriteCmd("@ARG\n")
	c.WriteCmd("D=M\n")
	c.pushDToStack()
	// push THIS
	c.WriteCmd("@THIS\n")
	c.WriteCmd("D=M\n")
	c.pushDToStack()
	// push THAT
	c.WriteCmd("@THAT\n")
	c.WriteCmd("D=M\n")
	c.pushDToStack()
	// ARG = SP - 5 - nArgs
	c.WriteCmd("@SP\n")
	c.WriteCmd("D=M\n")
	c.WriteCmd("@5\n")
	c.WriteCmd("D=D-A\n")
	c.WriteCmd(fmt.Sprintf("@%d\n", nArgs))
	c.WriteCmd("D=D-A\n")
	c.WriteCmd("@ARG\n")
	c.WriteCmd("M=D\n")
	// LCL = SP
	c.WriteCmd("@SP\n")
	c.WriteCmd("D=M\n")
	c.WriteCmd("@LCL\n")
	c.WriteCmd("M=D\n")
	// goto functionName
	c.WriteGoto("goto", functionName)
	// (returnAddress) declares a label for the return address eg. Sys$ret.1
	c.WriteLabel("label", returnLabel)
}

func (c *CodeWriter) getEndFrameMinusXToDRegister(val int) {
	c.WriteCmd(fmt.Sprintf("@%d\n", val))
	c.WriteCmd("D=A\n")
	// R14 is endFrame
	c.WriteCmd("@R14\n")
	c.WriteCmd("A=M-D\n")
	c.WriteCmd("D=M\n")
}
