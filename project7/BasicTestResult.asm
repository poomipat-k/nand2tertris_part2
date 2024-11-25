// push constant 10
@10
D=A
@SP
AM=M+1
A=A-1
M=D
// pop local 0
@SP
AM=M-1
D=M
@LCL
A=M
M=D
// push constant 21
@21
D=A
@SP
AM=M+1
A=A-1
M=D
// push constant 22
@22
D=A
@SP
AM=M+1
A=A-1
M=D
// pop argument 2
@2
D=A
@ARG
M=D+M
@SP
AM=M-1
D=M
@ARG
A=M
M=D
@2
D=A
@ARG
M=M-D
// pop argument 1
@1
D=A
@ARG
M=D+M
@SP
AM=M-1
D=M
@ARG
A=M
M=D
@1
D=A
@ARG
M=M-D
// push constant 36
@36
D=A
@SP
AM=M+1
A=A-1
M=D
// pop this 6
@6
D=A
@THIS
M=D+M
@SP
AM=M-1
D=M
@THIS
A=M
M=D
@6
D=A
@THIS
M=M-D
// push constant 42
@42
D=A
@SP
AM=M+1
A=A-1
M=D
// push constant 45
@45
D=A
@SP
AM=M+1
A=A-1
M=D
// pop that 5
@5
D=A
@THAT
M=D+M
@SP
AM=M-1
D=M
@THAT
A=M
M=D
@5
D=A
@THAT
M=M-D
// pop that 2
@2
D=A
@THAT
M=D+M
@SP
AM=M-1
D=M
@THAT
A=M
M=D
@2
D=A
@THAT
M=M-D
// push constant 510
@510
D=A
@SP
AM=M+1
A=A-1
M=D
// pop temp 6
@6
D=A
@
M=D+M
@SP
AM=M-1
D=M
@
A=M
M=D
@6
D=A
@
M=M-D
// push local 0
@0
D=A
@LCL
A=D+A
D=M
@SP
AM=M+1
A=A-1
M=D
// push that 5
@5
D=A
@THAT
A=D+A
D=M
@SP
AM=M+1
A=A-1
M=D
// add
// push argument 1
@1
D=A
@ARG
A=D+A
D=M
@SP
AM=M+1
A=A-1
M=D
// sub
// push this 6
@6
D=A
@THIS
A=D+A
D=M
@SP
AM=M+1
A=A-1
M=D
// push this 6
@6
D=A
@THIS
A=D+A
D=M
@SP
AM=M+1
A=A-1
M=D
// add
// sub
// push temp 6
@6
D=A
@
A=D+A
D=M
@SP
AM=M+1
A=A-1
M=D
// add
