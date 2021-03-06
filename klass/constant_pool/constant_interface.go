package constant_pool

import (
	"github.com/zouzhihao-994/gvm/loader"
)

type ConstantInterfaceMethodInfo struct {
	Tag uint8
	*ConstantPool
	ClassIdx       uint16
	NameAndTypeIdx uint16
}

// ReadInfo /*
func (im *ConstantInterfaceMethodInfo) ReadInfo(reader *loader.ClassReader) {
	im.ClassIdx = reader.ReadUint16()
	im.NameAndTypeIdx = reader.ReadUint16()
}

// ClassName /*
func (im *ConstantInterfaceMethodInfo) ClassName() string {
	return im.GetClassName(im.ClassIdx)
}

// NameAndDescriptor /*
func (im *ConstantInterfaceMethodInfo) NameAndDescriptor() (string, string) {
	return im.GetNameAndType(im.NameAndTypeIdx)
}
