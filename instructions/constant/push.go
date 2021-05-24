package constants

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/runtime"
)

type BIPUSH struct{ base.InstructionIndex8 } // Push byte

type SIPUSH struct{ base.InstructionIndex16 } // Push short

/*
将读取到的byte转化成int后推入到栈顶
*/
func (self BIPUSH) Execute(frame *runtime.Frame) {
	i := int32(self.Index)
	frame.PushInt(i)
}

/*
将读取到的short转化成int后推入到栈顶
*/
func (self SIPUSH) Execute(frame *runtime.Frame) {
	i := int32(self.Index)
	frame.PushInt(i)
}
