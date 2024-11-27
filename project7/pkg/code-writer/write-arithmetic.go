package codeWriter

import "fmt"

func (c *CodeWriter) WriteArithmetic(cmd string) {
	// write command in comment
	_, err := c.File.WriteString(fmt.Sprintf("// %s\n", cmd))
	check(err)
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
	}
}

func (c *CodeWriter) writeAdd() {
	_, err := c.File.WriteString("@SP\n")
	check(err)
	_, err = c.File.WriteString("AM=M-1\n")
	check(err)
	_, err = c.File.WriteString("D=M\n")
	check(err)
	_, err = c.File.WriteString("A=A-1\n")
	check(err)
	_, err = c.File.WriteString("M=D+M\n")
	check(err)
}

func (c *CodeWriter) writeSub() {
	_, err := c.File.WriteString("@SP\n")
	check(err)
	_, err = c.File.WriteString("AM=M-1\n")
	check(err)
	_, err = c.File.WriteString("D=M\n")
	check(err)
	_, err = c.File.WriteString("A=A-1\n")
	check(err)
	_, err = c.File.WriteString("M=M-D\n")
	check(err)
}

func (c *CodeWriter) writeNeg() {
	_, err := c.File.WriteString("@SP\n")
	check(err)
	_, err = c.File.WriteString("A=M-1\n")
	check(err)
	_, err = c.File.WriteString("M=-M\n")
	check(err)
}

func (c *CodeWriter) writeAnd() {
	_, err := c.File.WriteString("@SP\n")
	check(err)
	_, err = c.File.WriteString("AM=M-1\n")
	check(err)
	_, err = c.File.WriteString("D=M\n")
	check(err)
	_, err = c.File.WriteString("A=A-1\n")
	check(err)
	_, err = c.File.WriteString("M=D&M\n")
	check(err)
}

func (c *CodeWriter) writeOr() {
	_, err := c.File.WriteString("@SP\n")
	check(err)
	_, err = c.File.WriteString("AM=M-1\n")
	check(err)
	_, err = c.File.WriteString("D=M\n")
	check(err)
	_, err = c.File.WriteString("A=A-1\n")
	check(err)
	_, err = c.File.WriteString("M=D|M\n")
	check(err)
}

func (c *CodeWriter) writeNot() {
	_, err := c.File.WriteString("@SP\n")
	check(err)
	_, err = c.File.WriteString("A=M-1\n")
	check(err)
	_, err = c.File.WriteString("M=!M\n")
	check(err)
}
