package references

import (
	"github.com/zouzhihao-994/gvm/instructions/base"
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/runtime"
)

// GetStatic 获取类的静态字段值
// index指向当前类的运行时常量池，指向对象应该是一个字段类型的符号引用
type GetStatic struct {
	base.InstructionIndex16
}

func (i *GetStatic) Execute(frame *runtime.Frame) {
	fieldRef := frame.GetConstantFieldsInfo(i.Index)

	className := fieldRef.ClassName()
	fieldName, fieldDesc := fieldRef.NameAndDescriptor()
	k := klass.Perm.Get(className)

	// 判断是否需要进行加载
	if k == nil {
		k = klass.ParseByClassName(className)
		frame.RevertPC()
		base.InitClass(k, frame.Thread)
		return
	} else if !k.IsInit { //判断是否需要进行初始化
		frame.RevertPC()
		base.InitClass(k, frame.Thread)
		return
	}

	_, _, _, slots := k.GetStaticField(fieldName, fieldDesc).Fields()
	if fieldDesc == "D" || fieldDesc == "J" {
		frame.PushSlot(slots[0])
		slots = slots[1:]
	}
	frame.PushSlot(slots[0])
}
