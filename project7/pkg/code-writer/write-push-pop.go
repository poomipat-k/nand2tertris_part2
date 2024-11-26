package codeWriter

import "fmt"

func (c *CodeWriter) WritePushPop(cmd string, cmdType string, segment string, index int) {
	// write command in comment
	_, err := c.File.WriteString(fmt.Sprintf("// %s %s %d\n", cmd, segment, index))
	check(err)
	if cmdType == "C_PUSH" {
		c.writePush(segment, index)
	} else {
		c.writePop(segment, index)
	}
}

func (c *CodeWriter) writePush(segment string, index int) {
	var err error
	if segment == "constant" {
		_, err = c.File.WriteString(fmt.Sprintf("@%d\n", index))
		check(err)

		_, err = c.File.WriteString("D=A\n")
		check(err)
	} else {
		/*
			eg. push local 5
			@5
			D=A
			@local
			A=D+A
			D=M
		*/
		_, err = c.File.WriteString(fmt.Sprintf("@%d\n", index))
		check(err)

		_, err = c.File.WriteString("D=A\n")
		check(err)

		smSym := MEMORY_SEGMENT_DICT[segment]
		_, err = c.File.WriteString(fmt.Sprintf("@%s\n", smSym))
		check(err)

		_, err = c.File.WriteString("A=D+A\n")
		check(err)

		_, err = c.File.WriteString("D=M\n")
		check(err)
	}

	// increment SP
	_, err = c.File.WriteString("@SP\n")
	check(err)

	_, err = c.File.WriteString("AM=M+1\n")
	check(err)

	// Get back to the SP that want to push value to
	_, err = c.File.WriteString("A=A-1\n")
	check(err)

	_, err = c.File.WriteString("M=D\n")
	check(err)
}

func (c *CodeWriter) writePop(segment string, index int) {
	// invalid segment
	if segment == "constant" {
		return
	}
	var err error
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

	_, err = c.File.WriteString(fmt.Sprintf("@%s\n", smSym))
	check(err)
	_, err = c.File.WriteString("D=M\n")
	check(err)
	if index > 0 {
		_, err = c.File.WriteString(fmt.Sprintf("@%d\n", index))
		check(err)
		_, err = c.File.WriteString("D=D+A\n")
		check(err)
	}
	// save LCL+index address to R13 (general purpose register)
	_, err = c.File.WriteString("@R13\n")
	check(err)
	_, err = c.File.WriteString("M=D\n")
	check(err)

	// decrement SP
	_, err = c.File.WriteString("@SP\n")
	check(err)
	_, err = c.File.WriteString("AM=M-1\n")
	check(err)
	_, err = c.File.WriteString("D=M\n")
	check(err)

	// save D to desired ram position (Get address from M of R13 register)
	_, err = c.File.WriteString("@R13\n")
	check(err)
	_, err = c.File.WriteString("A=M\n")
	check(err)
	_, err = c.File.WriteString("M=D\n")
	check(err)
}
