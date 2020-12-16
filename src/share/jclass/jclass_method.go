package jclass

import (
	"github.com/zouzhihao-994/gvm/src/share/classfile"
	"github.com/zouzhihao-994/gvm/src/share/jclass/attribute"
	"github.com/zouzhihao-994/gvm/src/share/jclass/constant_pool"
)

type Methods []MethodInfo

type MethodInfo struct {
	accessFlag    uint16
	nameIdx       uint16
	descriptorIdx uint16
	attrCount     uint16
	attribute     attribute.AttributeInfos
	cp            constant_pool.ConstantPool
	argSlotCount  uint
}

// injected a code attribute for method
func (m MethodInfo) InjectCodeAttr() {
	//attributes := make(attribute.AttributeInfos, 1)
	ParseMethodDescriptor(m)
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

func (ms Methods) Clinit() MethodInfo {
	for idx := range ms {
		i := ms[idx].nameIdx
		nameStr := ms[idx].cp.GetUtf8(i)
		if nameStr == "<clinit>" {
			return ms[idx]
		}
	}
	panic("[gvm] the <clinit> is not exist")
}

func (m *MethodInfo) CP() constant_pool.ConstantPool {
	return m.cp
}

func (m MethodInfo) Attributes() attribute.AttributeInfos {
	return m.attribute
}

// 解析方法表
func parseMethod(count uint16, reader *classfile.ClassReader, pool constant_pool.ConstantPool) Methods {
	methods := make([]MethodInfo, count)
	for i := range methods {
		method := MethodInfo{}
		method.cp = pool
		method.accessFlag = reader.ReadUint16()
		method.nameIdx = reader.ReadUint16()
		method.descriptorIdx = reader.ReadUint16()
		method.attrCount = reader.ReadUint16()
		// 解析方法表中的属性表字段
		method.attribute = attribute.ParseAttributes(method.attrCount, reader, pool)
		methods[i] = method
	}
	return methods
}
