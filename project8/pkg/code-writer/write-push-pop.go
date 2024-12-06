package codeWriter

import "fmt"

func (c *CodeWriter) WritePushPop(cmd string, cmdType string, segment string, index int) {
	// write command in comment
	c.WriteComment(fmt.Sprintf("// %s %s %d\n", cmd, segment, index))

	if cmdType == "C_PUSH" {
		c.writePush(segment, index)
	} else {
		c.writePop(segment, index)
	}
}

func (c *CodeWriter) writePush(segment string, index int) {
	// get data to D register
	c._savePushDataToDRegister(segment, index)

	// // increment SP
	// c.WriteCmd("@SP\n")
	// c.WriteCmd("AM=M+1\n")

	// // Get back to the SP that want to push value to
	// c.WriteCmd("A=A-1\n")

	// c.WriteCmd("M=D\n")

	c.pushDToStack()
}

func (c *CodeWriter) pushDToStack() {
	// increment SP
	c.WriteCmd("@SP\n")
	c.WriteCmd("AM=M+1\n")

	// Get back to the SP that want to push value to
	c.WriteCmd("A=A-1\n")

	c.WriteCmd("M=D\n")
}

func (c *CodeWriter) writePop(segment string, index int) {
	// invalid segment
	if segment == "constant" {
		return
	}
	if segment == "temp" {
		c.writePopTemp(index)
		return
	}
	if segment == "static" {
		c.writePopStatic(index)
		return
	}
	if segment == "pointer" {
		c.writePopPointer(index)
		return
	}
	smSym := MEMORY_SEGMENT_DICT[segment]
	/*
		@LCL
		D=M
		@0
		D=D+A
		@R13
		M=D
		@SP
		AM=M-1
		D=M
		@R13
		A=M
		M=D
	*/

	c.WriteCmd(fmt.Sprintf("@%s\n", smSym))

	c.WriteCmd("D=M\n")

	c.WriteCmd(fmt.Sprintf("@%d\n", index))

	c.WriteCmd("D=D+A\n")

	// save LCL+index address to R13 (general purpose register)
	c.WriteCmd("@R13\n")

	c.WriteCmd("M=D\n")

	// decrement SP
	c.WriteCmd("@SP\n")

	c.WriteCmd("AM=M-1\n")

	c.WriteCmd("D=M\n")

	// save D to the desired ram position (Get address from M of R13 register)
	c.WriteCmd("@R13\n")

	c.WriteCmd("A=M\n")

	c.WriteCmd("M=D\n")

}

func (c *CodeWriter) writePopTemp(index int) {

	// decrement SP
	c.WriteCmd("@SP\n")

	c.WriteCmd("AM=M-1\n")

	c.WriteCmd("D=M\n")

	// temp is store at RAM[5] to RAM[12]
	offset := 5 + index
	c.WriteCmd(fmt.Sprintf("@%d\n", offset))

	c.WriteCmd("M=D\n")

}

func (c *CodeWriter) writePopStatic(index int) {

	// decrement SP
	c.WriteCmd("@SP\n")

	c.WriteCmd("AM=M-1\n")

	c.WriteCmd("D=M\n")

	c.WriteCmd(fmt.Sprintf("@%s.%d\n", c.filename, index))

	c.WriteCmd("M=D\n")

}

func (c *CodeWriter) writePopPointer(val int) {

	// decrement SP
	c.WriteCmd("@SP\n")

	c.WriteCmd("AM=M-1\n")

	c.WriteCmd("D=M\n")

	targetSegment := "THIS"
	if val == 1 {
		targetSegment = "THAT"
	}
	c.WriteCmd(fmt.Sprintf("@%s\n", targetSegment))

	c.WriteCmd("M=D\n")

}

func (c *CodeWriter) _savePushDataToDRegister(segment string, index int) {

	if segment == "constant" {
		c.WriteCmd(fmt.Sprintf("@%d\n", index))

		c.WriteCmd("D=A\n")

		return
	}
	if segment == "temp" {
		// temp is store at RAM[5] to RAM[12]
		offset := 5 + index
		c.WriteCmd(fmt.Sprintf("@%d\n", offset))

		c.WriteCmd("D=M\n")

		return
	}
	if segment == "static" {
		c.WriteCmd(fmt.Sprintf("@%s.%d\n", c.filename, index))

		c.WriteCmd("D=M\n")

		return
	}
	if segment == "pointer" {
		targetSegment := "THIS"
		if index == 1 {
			targetSegment = "THAT"
		}
		c.WriteCmd(fmt.Sprintf("@%s\n", targetSegment))

		c.WriteCmd("D=M\n")

		return
	}
	/*
			eg. push local 3
			@LCL
		 	D=M
			@3
			A=D+A
			D=M
			@SP
			AM=M+1
			A=A-1
			M=D
	*/
	smSym := MEMORY_SEGMENT_DICT[segment]
	c.WriteCmd(fmt.Sprintf("@%s\n", smSym))

	c.WriteCmd("D=M\n")

	c.WriteCmd(fmt.Sprintf("@%d\n", index))

	c.WriteCmd("A=D+A\n")

	c.WriteCmd("D=M\n")

}
