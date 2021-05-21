package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/oops"
	"github.com/zouzhihao-994/gvm/runtime"
	"github.com/zouzhihao-994/gvm/utils"
)

// InvokeVirtual 调用实例方法
type InvokeVirtual struct {
	base.InstructionIndex16
}

func (i *InvokeVirtual) Execute(frame *runtime.Frame) {
	constantMethod := frame.Method().CP().GetConstantInfo(i.Index).(*constant_pool.ConstantMethodInfo)
	methodNameStr, methodDescStr := constantMethod.NameAndDescriptor()
	objectRef := frame.OperandStack().GetByIdx(0)
	k := objectRef.Ref.(*oops.OopInstance).Klass()
	method, err, _ := k.FindMethod(methodNameStr, methodDescStr)
	utils.AssertError(err, "klass to find method err")

	base.InvokeMethod(frame, method, false)
}