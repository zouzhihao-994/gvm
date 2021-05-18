package stores

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// pop ref from stack and store the ref to localvars
type IASTORE struct {
	base.InstructionIndex0
}

func (i *IASTORE) Execute(frame *runtime.Frame) {
	val := frame.OperandStack().PopInt()
	idx := frame.OperandStack().PopInt()
	array := frame.OperandStack().PopRef()
	utils.AssertFalse(array == nil, "NullPointerException")
	array.ArrayData().SetIVal(idx, val)
}