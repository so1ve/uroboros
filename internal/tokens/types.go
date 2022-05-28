package tokens

type tokenTypeOptions struct {
	beforeExpr,
	startsExpr,
	isLoop,
	isAssign,
	prefix,
	postfix bool
	binop uint8
}

type TokenType struct {
	label string
	beforeExpr,
	startsExpr,
	isLoop,
	isAssign,
	prefix,
	postfix bool
	binop uint8
}

func NewTokenType(label string, options *tokenTypeOptions) *TokenType {
	return &TokenType{
		label:      label,
		beforeExpr: options.beforeExpr,
		startsExpr: options.startsExpr,
		isLoop:     options.isLoop,
		isAssign:   options.isAssign,
		prefix:     options.prefix,
		postfix:    options.postfix,
		binop:      options.binop,
	}
}

func binop(name string, prec uint8) *TokenType {
	return NewTokenType(name, &tokenTypeOptions{beforeExpr: true, binop: prec})
}

var beforeExpr = &tokenTypeOptions{beforeExpr: true}
var startsExpr = &tokenTypeOptions{startsExpr: true}

var Types = map[string]*TokenType{
	"num":       NewTokenType("num", startsExpr),
	"regexp":    NewTokenType("regexp", startsExpr),
	"string":    NewTokenType("string", startsExpr),
	"name":      NewTokenType("name", startsExpr),
	"privateId": NewTokenType("privateId", startsExpr),
	"eof":       NewTokenType("eof", &tokenTypeOptions{}),

	// Punctuation token types.
	"bracketL":        NewTokenType("[", &tokenTypeOptions{beforeExpr: true, startsExpr: true}),
	"bracketR":        NewTokenType("]", &tokenTypeOptions{}),
	"braceL":          NewTokenType("{", &tokenTypeOptions{beforeExpr: true, startsExpr: true}),
	"braceR":          NewTokenType("}", &tokenTypeOptions{}),
	"parenL":          NewTokenType("(", &tokenTypeOptions{beforeExpr: true, startsExpr: true}),
	"parenR":          NewTokenType(")", &tokenTypeOptions{}),
	"comma":           NewTokenType(",", beforeExpr),
	"semi":            NewTokenType(";", beforeExpr),
	"colon":           NewTokenType(":", beforeExpr),
	"dot":             NewTokenType(".", &tokenTypeOptions{}),
	"question":        NewTokenType("?", beforeExpr),
	"questionDot":     NewTokenType("?.", &tokenTypeOptions{}),
	"arrow":           NewTokenType("=>", beforeExpr),
	"template":        NewTokenType("template", &tokenTypeOptions{}),
	"invalidTemplate": NewTokenType("invalidTemplate", &tokenTypeOptions{}),
	"ellipsis":        NewTokenType("...", beforeExpr),
	"backQuote":       NewTokenType("`", startsExpr),
	"dollarBraceL":    NewTokenType("${", &tokenTypeOptions{beforeExpr: true, startsExpr: true}),

	// Operators. These carry several kinds of properties to help the
	// parser use them properly (the presence of these properties is
	// what categorizes them as operators).
	//
	// `binop`, when present, specifies that this operator is a binary
	// operator, and will refer to its precedence.
	//
	// `prefix` and `postfix` mark the operator as a prefix or postfix
	// unary operator.
	//
	// `isAssign` marks all of `=`, `+=`, `-=` etcetera, which act as
	// binary operators with a very low precedence, that should result
	// in AssignmentExpression nodes.

	"eq":         NewTokenType("=", &tokenTypeOptions{beforeExpr: true, isAssign: true}),
	"assign":     NewTokenType("_=", &tokenTypeOptions{beforeExpr: true, isAssign: true}),
	"incDec":     NewTokenType("++/--", &tokenTypeOptions{prefix: true, postfix: true, startsExpr: true}),
	"prefix":     NewTokenType("!/~", &tokenTypeOptions{beforeExpr: true, prefix: true, startsExpr: true}),
	"logicalOR":  binop("||", 1),
	"logicalAND": binop("&&", 2),
	"bitwiseOR":  binop("|", 3),
	"bitwiseXOR": binop("^", 4),
	"bitwiseAND": binop("&", 5),
	"equality":   binop("==/!=/===/!==", 6),
	"relational": binop("</>/<=/>=", 7),
	"bitShift":   binop("<</>>/>>>", 8),
	"plusMin":    NewTokenType("+/-", &tokenTypeOptions{beforeExpr: true, binop: 9, prefix: true, startsExpr: true}),
	"modulo":     binop("%", 10),
	"star":       binop("*", 10),
	"slash":      binop("/", 10),
	"starstar":   NewTokenType("**", &tokenTypeOptions{beforeExpr: true}),
	"coalesce":   binop("??", 1),
}
