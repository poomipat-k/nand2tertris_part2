package codeWriter

import "fmt"

func (c *CodeWriter) WriteArithmetic(cmd string) {
	// write command in comment
	c.WriteComment(fmt.Sprintf("// %s\n", cmd))

	if cmd == "add" {
		c.writeAdd()
	} else if cmd == "sub" {
		c.writeSub()
	} else if cmd == "neg" {
		c.writeNeg()
	} else if cmd == "and" {
		c.writeAnd()
	} else if cmd == "or" {
		c.writeOr()
	} else if cmd == "not" {
		c.writeNot()
	} else if cmd == "eq" {
		c.writeEq()
	} else if cmd == "gt" {
		c.writeGt()
	} else if cmd == "lt" {
		c.writeLt()
	}
}

func (c *CodeWriter) writeAdd() {
	c.WriteCmd("@SP\n")

	c.WriteCmd("AM=M-1\n")

	c.WriteCmd("D=M\n")

	c.WriteCmd("A=A-1\n")

	c.WriteCmd("M=D+M\n")

}

func (c *CodeWriter) writeSub() {
	c.WriteCmd("@SP\n")

	c.WriteCmd("AM=M-1\n")

	c.WriteCmd("D=M\n")

	c.WriteCmd("A=A-1\n")

	c.WriteCmd("M=M-D\n")

}

func (c *CodeWriter) writeNeg() {
	c.WriteCmd("@SP\n")

	c.WriteCmd("A=M-1\n")

	c.WriteCmd("M=-M\n")

}

func (c *CodeWriter) writeAnd() {
	c.WriteCmd("@SP\n")

	c.WriteCmd("AM=M-1\n")

	c.WriteCmd("D=M\n")

	c.WriteCmd("A=A-1\n")

	c.WriteCmd("M=D&M\n")

}

func (c *CodeWriter) writeOr() {
	c.WriteCmd("@SP\n")

	c.WriteCmd("AM=M-1\n")

	c.WriteCmd("D=M\n")

	c.WriteCmd("A=A-1\n")

	c.WriteCmd("M=D|M\n")

}

func (c *CodeWriter) writeNot() {
	c.WriteCmd("@SP\n")

	c.WriteCmd("A=M-1\n")

	c.WriteCmd("M=!M\n")

}

func (c *CodeWriter) writeEq() {
	c.WriteCmd("@SP\n")
	c.WriteCmd("AM=M-1\n")
	c.WriteCmd("D=M\n")
	c.WriteCmd("@SP\n")
	c.WriteCmd("AM=M-1\n")
	c.WriteCmd("D=D-M\n")
	// D == 0 then -1 else 0
	c.WriteCmd("M=-1\n") // set true
	current := c.lineCounter
	c.WriteCmd(fmt.Sprintf("@%d\n", current+5))
	c.WriteCmd("D;JEQ\n")
	c.WriteCmd("@SP\n")
	c.WriteCmd("A=M\n")
	c.WriteCmd("M=0\n") // set false
	c.WriteCmd("@SP\n")
	c.WriteCmd("M=M+1\n")
}

func (c *CodeWriter) writeGt() {
	c.WriteCmd("@SP\n")
	c.WriteCmd("AM=M-1\n")
	c.WriteCmd("D=M\n")
	c.WriteCmd("@SP\n")
	c.WriteCmd("AM=M-1\n")
	c.WriteCmd("D=M-D\n")
	// D > 0 then -1 else 0
	c.WriteCmd("M=-1\n") // set true
	current := c.lineCounter
	c.WriteCmd(fmt.Sprintf("@%d\n", current+5))
	c.WriteCmd("D;JGT\n")
	c.WriteCmd("@SP\n")
	c.WriteCmd("A=M\n")
	c.WriteCmd("M=0\n") // set false
	c.WriteCmd("@SP\n")
	c.WriteCmd("M=M+1\n")
}

func (c *CodeWriter) writeLt() {
	c.WriteCmd("@SP\n")
	c.WriteCmd("AM=M-1\n")
	c.WriteCmd("D=M\n")
	c.WriteCmd("@SP\n")
	c.WriteCmd("AM=M-1\n")
	c.WriteCmd("D=M-D\n")
	// D < 0 then -1 else 0
	c.WriteCmd("M=-1\n") // set true
	current := c.lineCounter
	c.WriteCmd(fmt.Sprintf("@%d\n", current+5))
	c.WriteCmd("D;JLT\n")
	c.WriteCmd("@SP\n")
	c.WriteCmd("A=M\n")
	c.WriteCmd("M=0\n") // set false
	c.WriteCmd("@SP\n")
	c.WriteCmd("M=M+1\n")
}
