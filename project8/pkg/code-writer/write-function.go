package codeWriter

import "fmt"

func (c *CodeWriter) WriteFunction(cmd string, functionName string, nLocalVars int) {
	c.WriteNonCmd(fmt.Sprintf("// %s\n", cmd))
	c.WriteNonCmd(fmt.Sprintf("(%s)\n", functionName))
	for i := 0; i < nLocalVars; i++ {
		c.writePush("constant", 0)
	}
}

func (c *CodeWriter) WriteReturn() {
	/*
		R13 for write push/pop arithmetic to use
		R14 is endFrame
		R15 is retAddr
	*/
	c.WriteNonCmd("// return\n")
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

func (c *CodeWriter) getEndFrameMinusXToDRegister(val int) {
	c.WriteCmd(fmt.Sprintf("@%d\n", val))
	c.WriteCmd("D=A\n")
	// R14 is endFrame
	c.WriteCmd("@R14\n")
	c.WriteCmd("A=M-D\n")
	c.WriteCmd("D=M\n")
}
