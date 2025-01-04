package symbolTable

import "log"

type SymbolTable struct {
	data            map[string]fields
	varCounter      int
	argumentCounter int
	staticCounter   int
	fieldCounter    int
}

type fields struct {
	dataType string
	kind     string
	number   int
}

func NewSymbolTable() *SymbolTable {
	d := map[string]fields{}
	return &SymbolTable{
		data: d,
	}
}

func (s *SymbolTable) Reset() {
	s.data = make(map[string]fields)
	s.varCounter = 0
	s.argumentCounter = 0
	s.staticCounter = 0
	s.fieldCounter = 0
}

func (s *SymbolTable) Define(name string, dataType string, kind string) {
	if row, found := s.data[name]; found && row.kind == kind {
		log.Fatal("SymbolTable: duplicate declaration, name: ", name)
	}
	vc := s.VarCount(kind)
	s.data[name] = fields{dataType: dataType, kind: kind, number: vc}
	if kind == "STATIC" {
		s.staticCounter++
	} else if kind == "FIELD" {
		s.fieldCounter++
	} else if kind == "ARG" {
		s.argumentCounter++
	} else if kind == "VAR" {
		s.varCounter++
	} else {
		log.Fatal("SymbolTable, kind is not valid, got: ", kind)
	}
}

func (s *SymbolTable) VarCount(kind string) int {
	if kind == "STATIC" {
		return s.staticCounter
	}
	if kind == "FIELD" {
		return s.fieldCounter
	}
	if kind == "ARG" {
		return s.argumentCounter
	}
	if kind == "VAR" {
		return s.varCounter
	}
	return -1
}

func (s *SymbolTable) KindOf(name string) string {
	if _, found := s.data[name]; !found {
		return ""
	}
	return s.data[name].kind
}

func (s *SymbolTable) TypeOf(name string) string {
	if _, found := s.data[name]; !found {
		return ""
	}
	return s.data[name].dataType
}

func (s *SymbolTable) IndexOf(name string) int {
	if _, found := s.data[name]; !found {
		return -1
	}
	return s.data[name].number
}
