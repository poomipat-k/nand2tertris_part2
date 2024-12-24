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

var keywordConstant = map[string]bool{
	"true":  true,
	"false": true,
	"null":  true,
	"this":  true,
}

var opSymbol = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
	"&": true,
	"|": true,
	"<": true,
	">": true,
	"=": true,
}

var unaryOp = map[string]bool{
	"-": true,
	"~": true,
}
