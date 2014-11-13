// TODO(u): Read up about the use of nuw and nsw and update releveant
// instructions.

// TODO(u): Read up about the use of fast-math flags and update releveant
// instructions.

// TODO(u): Read up about the usage of exact and update releveant instructions.

package ir

// An Instruction belongs to one of the following groups:
//
//    * terminator instructions
//    * binary instructions
//    * bitwise binary instructions
//    * memory instructions
//    * other instructions
//
// References:
//    http://llvm.org/docs/LangRef.html#instruction-reference
type Instruction interface {
	// isInst ensures that only instructions can be assigned to the Instruction
	// interface.
	isInst()
}

// =============================================================================
// Terminator Instructions
//
//    ref: http://llvm.org/docs/LangRef.html#terminators
// =============================================================================

// The ReturnInst returns control flow (and optionally a value) from a function
// back to the caller.
//
// Syntax:
//    ret <Type> <Val>
//    ret void
//
// Reference:
//    http://llvm.org/docs/LangRef.html#i-ret
type ReturnInst struct {
	// Return type.
	Type Type
	// Return value; or nil in case of a void return.
	Val Value
}

// The CondBranchInst transfers control flow to one of two basic blocks in the
// current function based on a boolean branching condition.
//
// Syntax:
//    br i1 <Cond>, label <TargetTrue>, label <TargetFalse>
//
// References:
//    http://llvm.org/docs/LangRef.html#i-br
type CondBranchInst struct {
	// Boolean branching condition.
	Cond Value
	// Target branch when the condition evaluates to true.
	True *BasicBlock
	// Target branch when the condition evaluates to false.
	False *BasicBlock
}

// The BranchInst transfers control flow to a basic block in the current
// function.
//
// Syntax:
//    br label <Target>
//
// References:
//    http://llvm.org/docs/LangRef.html#i-br
type BranchInst struct {
	// Target branch.
	Target *BasicBlock
}

// The SwitchInst transfers control flow to one of several basic blocks in the
// current function.
//
// Syntax:
//    switch <IntType> <Val>, label <TargetDefault> [ <IntType> <Const>, label <Target> ... ]
//
// References:
//    http://llvm.org/docs/LangRef.html#i-switch
type SwitchInst struct {
	// TODO(u): Restrict Type to IntType.
	// Comparasion type.
	Type Type
	// Comparasion value.
	Val Value
	// Default target.
	Default *BasicBlock
	// Switch cases.
	Cases []struct {
		// Case value.
		Val Constant
		// Case target.
		Target *BasicBlock
	}
}

// TODO(u): Add the following terminator instructions:
//    - indirectbr
//    - invoke
//    - resume
//    - unreachable

// =============================================================================
// Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#binaryops
// =============================================================================

// The AddInst returns the sum of its two operands, which may be integers or
// vectors of integer values.
//
// Syntax:
//    <Result> = add <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#i-add
type AddInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The FaddInst returns the sum of its two operands, which may be floating point
// values or vectors of floating point values.
//
// Syntax:
//    <Result> = fadd <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#i-fadd
type FaddInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The SubInst returns the difference of its two operands, which may be integers
// or vectors of integer values.
//
// Syntax:
//    <Result> = sub <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#sub-instruction
type SubInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The FsubInst returns the difference of its two operands, which may be
// floating point values or vectors of floating point values.
//
// Syntax:
//    <Result> = fsub <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#i-fsub
type FsubInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The MulInst returns the product of its two operands, which may be integers or
// vectors of integer values.
//
// Syntax:
//    <Result> = mul <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#mul-instruction
type MulInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The FmulInst returns the product of its two operands, which may be floating
// point values or vectors of floating point values.
//
// Syntax:
//    <Result> = fmul <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#fmul-instruction
type FmulInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The UdivInst returns the unsigned integer quotient of its two operands, which
// may be integers or vectors of integer values.
//
// Syntax:
//    <Result> = udiv <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#udiv-instruction
type UdivInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The SdivInst returns the signed integer quotient of its two operands, which
// may be integers or vectors of integer values.
//
// Syntax:
//    <Result> = sdiv <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#sdiv-instruction
type SdivInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The FdivInst returns the quotient of its two operands, which may be floating
// point values or vectors of floating point values.
//
// Syntax:
//    <Result> = fdiv <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#fdiv-instruction
type FdivInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The UremInst returns the unsigned integer remainder of a division between its
// two operands, which may be integers or vectors of integers.
//
// Syntax:
//    <Result> = urem <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#urem-instruction
type UremInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The SremInst returns the signed integer remainder of a division between its
// two operands, which may be integers or vectors of integers.
//
// Syntax:
//    <Result> = srem <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#srem-instruction
type SremInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The FremInst returns the remainder of a division between its two operands,
// which may be floating point values or vectors of floating point values.
//
// Syntax:
//    <Result> = frem <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#frem-instruction
type FremInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// =============================================================================
// Bitwise Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#bitwiseops
// =============================================================================

// The ShlInst returns the first operand shifted to the left a specified number
// of bits. The arguments may be integers or vectors of integer values.
//
// Syntax:
//    <Result> = shl <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#shl-instruction
type ShlInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The LshrInst (logical shift right) returns the first operand shifted to the
// right a specified number of bits with zero fill. The arguments may be
// integers or vectors of integer values.
//
// Syntax:
//    <Result> = lshr <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#lshr-instruction
type LshrInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The AshrInst (arithmetic shift right) returns the first operand shifted to
// the right a specified number of bits with sign extension. The arguments may
// be integers or vectors of integer values.
//
// Syntax:
//    <Result> = ashr <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#ashr-instruction
type AshrInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The AndInst returns the bitwise logical and of its two operands, which may be
// integers or vectors of integer values.
//
// Syntax:
//    <Result> = and <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#and-instruction
type AndInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The OrInst returns the bitwise logical inclusive or of its two operands,
// which may be integers or vectors of integer values.
//
// Syntax:
//    <Result> = or <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#or-instruction
type OrInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// The XorInst returns the bitwise logical exclusive or of its two operands,
// which may be integers or vectors of integer values.
//
// Syntax:
//    <Result> = xor <Type> <Op1>, <Op2>
//
// References:
//    http://llvm.org/docs/LangRef.html#xor-instruction
type XorInst struct {
	// Operand type.
	Type Type
	// Operands.
	Op1, Op2 Value
}

// =============================================================================
// Memory Access and Addressing Operations
//
//    ref: http://llvm.org/docs/LangRef.html#memoryops
// =============================================================================

// TODO(u): Add memory access and addressing operations.

// =============================================================================
// Other Operations
//
//    ref: http://llvm.org/docs/LangRef.html#otherops
// =============================================================================

// TODO(u): Add other operations.

// isInst ensures that only instructions can be assigned to the Instruction
// interface.
func (ReturnInst) isInst()     {}
func (CondBranchInst) isInst() {}
func (BranchInst) isInst()     {}
func (SwitchInst) isInst()     {}
func (AddInst) isInst()        {}
func (FaddInst) isInst()       {}
func (SubInst) isInst()        {}
func (FsubInst) isInst()       {}
func (MulInst) isInst()        {}
func (FmulInst) isInst()       {}
func (UdivInst) isInst()       {}
func (SdivInst) isInst()       {}
func (FdivInst) isInst()       {}
func (UremInst) isInst()       {}
func (SremInst) isInst()       {}
func (FremInst) isInst()       {}
