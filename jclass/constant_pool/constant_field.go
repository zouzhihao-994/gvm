package constant_pool

import (
	"github.com/zouzhihao-994/gvm/classloader"
)

type ConstantFieldInfo struct {
	Tag              uint8
	Cp               ConstantPool
	ClassIndex       uint16
	NameAndTypeIndex uint16
	NameAndType      ConstantNameAndTypeInfo
}

func (field *ConstantFieldInfo) ReadInfo(reader *classloader.ClassReader) {
	field.ClassIndex = reader.ReadUint16()
	field.NameAndTypeIndex = reader.ReadUint16()
}

func (field *ConstantFieldInfo) ClassName() string {
	return field.Cp.GetClassName(field.ClassIndex)
}
func (field *ConstantFieldInfo) NameAndDescriptor() (string, string) {
	return field.Cp.GetNameAndType(field.NameAndTypeIndex)
}