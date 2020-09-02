package classfile

import "fmt"

/*
字段表
*/
type MemberInfo struct {
	// 常量池指针
	cp ConstantPool
	// 访问类型
	accessFlags uint16
	// 字段名索引 -> 常量池
	nameIndex uint16
	// 描述符索引 -> 常量池
	descriptorIndex uint16
	// 属性表
	attributes []AttributeInfo
}

/*
获取方法的Code属性
*/
func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	// 遍历属性表
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}

	return nil
}

/*
读取字段表
*/
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	cpCount := reader.readUint16()
	members := make([]*MemberInfo, cpCount)

	// 遍历数组
	for i := range members {
		// 解析每一个字段和方法
		members[i] = readMember(reader, cp)
	}

	return members
}

/*
解析字段数据
*/
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	member := &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
	fmt.Printf("[gvm][readMember] parse a member_info,"+
		"name: %v ,describe:%v\n", member.Name(), member.Descriptor())
	return member
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

/*
获取方法或字段名称
*/
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)

}

/*
获取字段或方法的描述符
*/
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
