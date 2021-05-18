package jclass

import (
	"github.com/zouzhihao-994/gvm/classfile"
	"github.com/zouzhihao-994/gvm/jclass/attribute"
	"github.com/zouzhihao-994/gvm/jclass/constant_pool"
)

type Methods []*MethodInfo

type MethodInfo struct {
	accessFlag    uint16
	nameIdx       uint16
	descriptorIdx uint16
	attrCount     uint16
	attribute     attribute.AttributesInfo
	cp            constant_pool.ConstantPool
	argSlotCount  uint
	jclass        *JClassInstance
}

// InjectCodeAttr injected a code attribute for method
func (m *MethodInfo) InjectCodeAttr() {
	if !IsNative(m.accessFlag) {
		panic("[gvm] Inject CodeAttr error , not is native")
	}
	tmpMaxStack := uint16(4)
	tmpMaxLocal := uint16(4)
	attributes := make(attribute.AttributesInfo, 1)
	methodDescriptor := ParseMethodDescriptor(m.Descriptor())
	var codeAttr *attribute.Attr_Code
	switch methodDescriptor.returnTypt {
	case "V":
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xb1}, m.cp) // return
	case "D":
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xaf}, m.cp) // dreturn
	case "F":
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xae}, m.cp) // freturn
	case "J":
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xad}, m.cp) // lreturn
	case "L", "[":
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xb0}, m.cp) // areturn
	default:
		codeAttr = attribute.CreateCodeAttr(tmpMaxStack, tmpMaxLocal, []byte{0xfe, 0xbc}, m.cp) // ireturn
	}
	attributes[0] = codeAttr
	m.attribute = attributes
}

func (m MethodInfo) JClass() *JClassInstance {
	return m.jclass
}

func (m *MethodInfo) SetJClass(jci *JClassInstance) {
	m.jclass = jci
}

func (m MethodInfo) Descriptor() string {
	return m.cp.GetUtf8(m.descriptorIdx)
}

func (m MethodInfo) Name() string {
	return m.cp.GetUtf8(m.nameIdx)
}

func (m MethodInfo) AccessFlag() uint16 {
	return m.accessFlag
}

func (m MethodInfo) ArgSlotCount() uint {
	return m.argSlotCount
}

func (ms Methods) Clinit() (*MethodInfo, bool) {
	for idx := range ms {
		i := ms[idx].nameIdx
		nameStr := ms[idx].cp.GetUtf8(i)
		if nameStr == "<clinit>" {
			return ms[idx], true
		}
	}
	return nil, false
}

func (ms Methods) FindMethod(name, desc string) (*MethodInfo, bool) {
	for idx := range ms {
		nameStr := ms[idx].cp.GetUtf8(ms[idx].nameIdx)
		descStr := ms[idx].cp.GetUtf8(ms[idx].descriptorIdx)
		if nameStr == name && descStr == desc {
			return ms[idx], true
		}
	}
	return nil, false
}

func (m *MethodInfo) CP() constant_pool.ConstantPool {
	return m.cp
}

func (m MethodInfo) Attributes() attribute.AttributesInfo {
	return m.attribute
}

// 解析方法表
func parseMethod(count uint16, reader *classfile.ClassReader, pool constant_pool.ConstantPool) Methods {
	methods := make([]*MethodInfo, count)
	for i := range methods {
		method := &MethodInfo{}
		method.cp = pool
		method.accessFlag = reader.ReadUint16()
		method.nameIdx = reader.ReadUint16()
		method.descriptorIdx = reader.ReadUint16()
		method.attrCount = reader.ReadUint16()
		// 解析方法表中的属性表字段
		method.attribute = attribute.ParseAttributes(method.attrCount, reader, pool)
		methods[i] = method
		method.argSlotCount = ParseMethodDescriptor(method.Descriptor()).ParamsCount()
	}
	return methods
}

func (m *MethodInfo) IsRegisterNatives() bool {
	return IsStatic(m.accessFlag) && m.Name() == "registerNatives" && m.Descriptor() == "()V"
}

func (m *MethodInfo) IsInitIDs() bool {
	return IsStatic(m.accessFlag) && m.Name() == "initIDs" && m.Descriptor() == "()V"
}
