package attribute

import (
	"github.com/zouzhihao-994/gvm/klass/constant_pool"
	"github.com/zouzhihao-994/gvm/loader"
)

// AttrConstantvalue ConstantValue 存在 FieldInfo 的属性表中,是一个常量表达式
// 用法：如果字段被static修饰，对应为 FieldInfo 的 accessFlags 设置了 ACC_STATIC 标识
// 那么这个值将设置成类的 ConstantValue 属性，在类或者接口初始化的时候会也会对该字段进行初始化
// 注意: 这个初始化是要早于调用的类初始化步骤
type AttrConstantvalue struct {
	nameIdx uint16
	name    string
	attrLen uint32
	*constant_pool.ConstantPool
	// 常量值的索引
	ValeIdx uint16
}

func (cv *AttrConstantvalue) parse(reader *loader.ClassReader) {
	cv.ValeIdx = reader.ReadUint16()
}

func (cv AttrConstantvalue) Name() string {
	return cv.name
}
