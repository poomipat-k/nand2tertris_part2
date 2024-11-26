package codeWriter

import "fmt"

func (c *CodeWriter) WriteArithmetic(cmd string) {
	// write command in comment
	_, err := c.File.WriteString(fmt.Sprintf("// %s\n", cmd))
	check(err)
	if cmd == "add" {
		c.writeAdd()
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
