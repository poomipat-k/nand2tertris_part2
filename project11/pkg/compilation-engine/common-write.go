package compilationEngine

// func (e *Engine) writeSymbol() {
// 	e.WriteString(fmt.Sprintf("<symbol> %s </symbol>\n", e.tk.Symbol()))
// }

// func (e *Engine) writeKeyword() {
// 	e.WriteString(fmt.Sprintf("<keyword> %s </keyword>\n", e.tk.Keyword()))
// }

// func (e *Engine) writeIntegerConst() {
// 	e.WriteString(fmt.Sprintf("<integerConstant> %d </integerConstant>\n", e.tk.IntVal()))
// }

// func (e *Engine) writeStringConst() {
// 	e.WriteString(fmt.Sprintf("<stringConstant> %s </stringConstant>\n", e.tk.StringVal()))
// }

/*
role: [dec, used]
kind: [VAR, ARGUMENT, STATIC, FIELD, CLASS, SUBROUTINE]
*/
// func (e *Engine) writeIdentifier(identifier string, role string, kind string) {
// 	if identifier == "" {
// 		log.Fatal("writeIdentifier, identifier should not empty")
// 	}
// 	tag := fmt.Sprintf("identifier_%s_%s", role, strings.ToLower(kind))
// 	// append _runningNumber if kind is one of [VAR, ARGUMENT, STATIC, FIELD]
// 	if kind == symbolTable.VAR ||
// 		kind == symbolTable.ARG ||
// 		kind == symbolTable.STATIC ||
// 		kind == symbolTable.FIELD {
// 		name := identifier
// 		index := e.subroutineST.IndexOf(name)
// 		if index == -1 {
// 			index = e.classST.IndexOf(name)
// 		}
// 		if index == -1 {
// 			log.Fatal("writeIdentifier, not found in symbol table, name: ", name)
// 		}
// 		tag += fmt.Sprintf("_%d", index)
// 	}

// 	// e.WriteString(fmt.Sprintf("<%s> %s </%s>\n", tag, identifier, tag))
// }
