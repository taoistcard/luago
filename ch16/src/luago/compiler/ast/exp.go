package ast

type Exp interface{}

type NilExp struct{ Line int }
type TrueExp struct{ Line int }
type FalseExp struct{ Line int }
type VarargExp struct{ Line int }
type IntegerExp struct {
	Line int
	Val  int64
}
type FloatExp struct {
	Line int
	Val  float64
}
type StringExp struct {
	Line int
	Str  string
}
type NameExp struct {
	Line int
	Name string
}

//一元运算符表达式
type UnopExp struct {
	Line int //操作符的行号
	Op   int
	Exp
}

//二元表达式
type BinopExp struct {
	Line int //操作符的行号
	Op   int
	Exp1 Exp
	Exp2 Exp
}

//拼接表达式
type ConcatExp struct {
	Line int //末位行号
	Exps []Exp
}

//表构造表达式
type TableConstructorExp struct {
	Line     int //line "{""
	LastLine int //line of "}"
	KeyExps  []Exp
	ValExps  []Exp
}

//函数定义表达式
type FuncDefExp struct {
	Line     int
	LastLine int
	ParList  []string
	IsVararg bool
	Block    *Block
}

//圆括号表达式
type ParensExp struct {
	Exp Exp
}

//表访问表达式
type TableAccessExp struct {
	LastLine  int // line of "]"
	PrefixExp Exp
	KeyExps   Exp
}
type FuncCallExp struct {
	Line      int //line of "("
	LastLine  int //line of ")"
	PrefixExp Exp
	NameExp   *StringExp
	Args      []Exp
}
