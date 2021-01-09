package attribute

import "github.com/zouzhihao-994/gvm/src/share/classfile"

// 保存invokedynamic指令引用的引导方法限定符号
// classfile 中最多一个 BootstrapmethodsAttribute 属性
type BootstrapmethodsAttribute struct {
	// utf8 格式
	nameIdx uint16
	name    string
	// 属性长度，不包括nameIdx和attrLen的6字节
	attrLen uint32
	// 方法数量
	methodsNum uint16
	// 方法数组
	methods []BootstrapMethod
}

func (attr *BootstrapmethodsAttribute) parse(reader *classfile.ClassReader) {
	attr.methodsNum = reader.ReadUint16()
	attr.methods = make([]BootstrapMethod, attr.methodsNum)
	for i := 0; i < int(attr.methodsNum); i++ {
		bsm := BootstrapMethod{}
		bsm.parse(reader)
	}
}

func (attr *BootstrapmethodsAttribute) Name() string {
	return attr.name
}

// 该结构指明了一个引导方法，并指明了一个由索引组成的序列
// 序列里的索引指明引导方法的静态参数
type BootstrapMethod struct {
	// 常量池索引  constant_pool.ConstantMethodHandle 结构
	// 方法的具体形式由 references.INVOKE_DYNAMIC 指令持续解析决定
	methodRef uint16
	// 决定 arguments 数组大小
	argumentsNum uint16
	// 每一个元素都是运行时的一个索引，结构必须是下列之一：
	// constant_pool.ConstantStringInfo, constant_pool.ConstantClassInfo, constant_pool.ConstantIntegerInfo
	// constant_pool.ConstantFloatInfo, constant_pool.ConstantDoubleInfo, constant_pool.ConstantMethodHandleInfo
	// constant_pool.ConstantLongInfo, constant_pool.ConstantMethodTypeInfo, constant_pool.ConstantMethodInfo
	arguments []uint16
}

func (attr *BootstrapMethod) parse(reader *classfile.ClassReader) {
	attr.methodRef = reader.ReadUint16()
	attr.argumentsNum = reader.ReadUint16()
	attr.arguments = make([]uint16, attr.argumentsNum)
	for i := 0; i < int(attr.argumentsNum); i++ {
		attr.arguments[i] = reader.ReadUint16()
	}
}
