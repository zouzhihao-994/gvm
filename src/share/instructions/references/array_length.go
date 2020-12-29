package references

import (
	"github.com/zouzhihao-994/gvm/src/share/exception"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/runtime"
)

type ARRAY_LENGTH struct {
	base.InstructionIndex8
}

func (i *ARRAY_LENGTH) Execute(frame *runtime.Frame) {
	arrayRef := frame.OperandStack().PopRef()
	exception.AssertFalse(arrayRef == nil, "NullPointerException")

}