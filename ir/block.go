package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// A BasicBlock represents an LLVM IR basic block, which consists of a sequence
// of non-branching instructions, terminated by a control flow instruction (e.g.
// br or ret).
//
// Basic blocks may be referenced from terminator instructions (e.g. br), and
// are thus considered LLVM IR values of label type.
type BasicBlock struct {
	// Parent function of the basic block.
	parent *Function
	// Label name of the basic block.
	name string
	// Non-branching instructions of the basic block.
	insts []Instruction
	// Terminator instruction of the basic block.
	term Terminator
}

// Type returns the type of the basic block.
func (block *BasicBlock) Type() types.Type {
	return types.Label
}

// Ident returns the identifier associated with the basic block.
func (block *BasicBlock) Ident() string {
	// TODO: Encode name if containing special characters.
	return "%" + block.name
}

// LLVMString returns the LLVM syntax representation of the basic block.
func (block *BasicBlock) LLVMString() string {
	// TODO: Encode name if containing special characters.
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s:\n", block.name)
	for _, inst := range block.insts {
		fmt.Fprintf(buf, "\t%v\n", inst.LLVMString())
	}
	fmt.Fprintf(buf, "\t%v\n", block.term.LLVMString())
	return buf.String()
}

// NewAdd appends a new add instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewAdd(x, y value.Value) *Add {
	inst := NewAdd(x, y)
	inst.SetParent(block)
	block.insts = append(block.insts, inst)
	return inst
}

// NewCall appends a new call instruction to the basic block based on the given
// callee and function arguments.
func (block *BasicBlock) NewCall(callee *Function, args ...value.Value) *Call {
	inst := NewCall(callee, args...)
	inst.SetParent(block)
	block.insts = append(block.insts, inst)
	return inst
}

// NewMul appends a new mul instruction to the basic block based on the given
// operands.
func (block *BasicBlock) NewMul(x, y value.Value) *Mul {
	inst := NewMul(x, y)
	inst.SetParent(block)
	block.insts = append(block.insts, inst)
	return inst
}

// NewLoad appends a new load instruction to the basic block based on the given
// source address.
func (block *BasicBlock) NewLoad(src value.Value) *Load {
	inst := NewLoad(src)
	inst.SetParent(block)
	block.insts = append(block.insts, inst)
	return inst
}

// NewRet appends a new ret instruction to the basic block based on the given
// return value. A nil return value indicates a "void" return instruction.
func (block *BasicBlock) NewRet(x value.Value) *Ret {
	term := NewRet(x)
	term.SetParent(block)
	block.term = term
	return term
}

// Parent returns the parent function of the basic block.
func (block *BasicBlock) Parent() *Function {
	return block.parent
}

// SetParent sets the parent function of the basic block.
func (block *BasicBlock) SetParent(parent *Function) {
	block.parent = parent
}
