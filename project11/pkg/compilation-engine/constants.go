package compilationEngine

var classVarScope = map[string]bool{
	"static": true,
	"field":  true,
}

var subroutineDec = map[string]bool{
	"constructor": true,
	"function":    true,
	"method":      true,
}

var jackType = map[string]bool{
	"int":     true,
	"char":    true,
	"boolean": true,
}

var statementKeywords = map[string]bool{
	"let":    true,
	"if":     true,
	"while":  true,
	"do":     true,
	"return": true,
}

var keywordConstant = map[string]int{
	"true":  1,
	"false": 0,
	"null":  0,
	"this":  0,
}

var opSymbol = map[string]string{
	"+":     "add",
	"-":     "sub",
	"*":     "call Math.multiply 2",
	"/":     "call Math.divide 2",
	"&amp;": "and",
	"|":     "or",
	"&lt;":  "lt",
	"&gt;":  "gt",
	"=":     "eq",
}

var unaryOp = map[string]string{
	"-": "neg",
	"~": "not",
}
