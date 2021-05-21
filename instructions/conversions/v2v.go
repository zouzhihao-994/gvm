package conversions

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type I2f struct {
	base.InstructionIndex0
}
type F2i struct {
	base.InstructionIndex0
}

func (i I2f) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushFloat(float32(frame.OperandStack().PopInt()))
}

func (i F2i) Execute(frame *runtime.Frame) {
	frame.OperandStack().PushInt(int32(frame.OperandStack().PopFloat()))
}