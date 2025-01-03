package symbolTable

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

}

func (s *SymbolTable) VarCount(name string) int {
	kind := s.KindOf(name)
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
	return s.data[name].kind
}

func (s *SymbolTable) TypeOf(name string) string {
	return s.data[name].dataType
}

func (s *SymbolTable) IndexOf(name string) int {
	return s.data[name].number
}
