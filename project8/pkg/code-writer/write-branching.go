package codeWriter

import "fmt"

func (c *CodeWriter) WriteLabel(cmd string, label string) {
	c.WriteNonCmd(fmt.Sprintf("// %s\n", cmd))
	c.WriteNonCmd(fmt.Sprintf("(%s)\n", label))
}

func (c *CodeWriter) WriteIf(cmd string, target string) {
	c.WriteNonCmd(fmt.Sprintf("// %s\n", cmd))
	c.WriteCmd("@SP\n")
	c.WriteCmd("AM=M-1\n")
	c.WriteCmd("D=M\n")
	c.WriteCmd(fmt.Sprintf("@%s\n", target))
	c.WriteCmd("D;JNE\n")
}

func (c *CodeWriter) WriteGoto(cmd string, target string) {
	c.WriteNonCmd(fmt.Sprintf("// %s\n", cmd))
	c.WriteCmd(fmt.Sprintf("@%s\n", target))
	c.WriteCmd("0;JMP\n")
}
