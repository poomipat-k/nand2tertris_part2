package vmWriter

type VMWriter struct {
}

func NewVMWriter() *VMWriter {
	return &VMWriter{}
}

func (v *VMWriter) WritePush(segment string, index int) {

}

func (v *VMWriter) WritePop(segment string, index int) {

}

/*
ADD, SUB, NEG, EQ, GT, LT, AND, OR, NOT
*/
func (v *VMWriter) WriteArithmetic(cmd string) {

}

func (v *VMWriter) WriteLabel(label string) {

}

func (v *VMWriter) WriteGoto(label string) {

}

func (v *VMWriter) WriteIf(label string) {

}

func (v *VMWriter) WriteCall(name string, nArgs int) {

}

func (v *VMWriter) WriteFunction(name string, nVars int) {

}

func (v *VMWriter) WriteReturn() {

}

func (v *VMWriter) Close() {

}
