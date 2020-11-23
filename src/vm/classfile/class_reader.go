package classfile

import "encoding/binary"

// 读取class文件的二进制数据到ClassReader中
type ClassReader struct {
	classFile []byte
}

/*
从ClassReader中读取1字节数据,对应jvm中的u1
*/
func (classReader *ClassReader) readUint8() uint8 {
	// 获取第1个字节的数据
	val := classReader.classFile[0]
	// 删除第一个字节的数
	classReader.classFile = classReader.classFile[1:]
	return val
}

/*
读取2字节长度,对应jvm中的u2
*/
func (classReader *ClassReader) readUint16() uint16 {
	// 从self.data中读取16位的数据
	val := binary.BigEndian.Uint16(classReader.classFile)
	// 新的data中省略前两位
	classReader.classFile = classReader.classFile[2:]
	return val
}

/**
读取4字节长度,对应jvm中的u4
*/
func (classReader *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(classReader.classFile)
	classReader.classFile = classReader.classFile[4:]
	return val
}

/*
读取8字节
*/
func (classReader *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(classReader.classFile)
	classReader.classFile = classReader.classFile[8:]
	return val
}

/*
读取2字节的表头
*/
func (classReader *ClassReader) readUint16s() []uint16 {
	// 调用unit16()的方法
	n := classReader.readUint16()
	// n长度的uint16[]数组
	s := make([]uint16, n)
	// 填充数据
	for i := range s {
		s[i] = classReader.readUint16()
	}
	return s
}

/*
读取length字节长度的数据
*/
func (classReader *ClassReader) readBytes(length uint32) []byte {
	val := classReader.classFile[:length]
	classReader.classFile = classReader.classFile[length:]
	return val
}
