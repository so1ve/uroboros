package scope

import (
	"github.com/so1ve/uroboros/internal/utils"
)

type ScopeType uint16

const (
	ScopeTop ScopeType = 1 << iota
	ScopeFunction
	ScopeAsync
	ScopeGenerator
	ScopeArrow
	ScopeSimpleCatch
	ScopeSuper
	ScopeDirectSuper
	ScopeClassStaticBlock
	ScopeVar = ScopeTop | ScopeFunction | ScopeClassStaticBlock
)

type BindType uint8

const (
	BindNone        BindType = iota // Not a binding
	BindVar                         // Var-style binding
	BindLexical                     // Let- or const-style binding
	BindFunction                    // Function declaration
	BindSimpleCatch                 // Simple (identifier pattern) catch binding
	BindOutside                     // Special case for function names as bound inside the function
)

func FunctionFlags(isAsync, isGenerator bool) uint16 {
	return uint16(ScopeFunction) | uint16(utils.If(isAsync, ScopeAsync, 0)) | uint16(utils.If(isGenerator, ScopeGenerator, 0))
}
