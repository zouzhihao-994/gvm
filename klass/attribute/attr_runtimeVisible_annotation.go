package attribute

import "github.com/zouzhihao-994/gvm/loader"

// AttrRuntimevisibleannotation located in the classFile, field_info, method_info
type AttrRuntimevisibleannotation struct {
	// uft8 type, the name of annotation
	nameIdx uint16
	name    string
	// the attrLen of AttrRuntimevisibleannotation(Not including the previous six byteso)
	attrLen uint32
	// the num of this annotions
	annotationNum uint16
	annotations   []annotation
}

type annotation struct {
	// uf8 types, name of annotation
	typeIndex            uint16
	elementValuePairsNum uint16
	elementValuePairs    []elementValuePair
}

type elementValuePair struct {
	// utf8 type, name of field on annotation
	elementNameIdx uint16
	// the value of field on annotation
	elementValue elementValue
}

type elementValue struct {
	// an ASCII characters
	tag uint8
	//
	union value
}

type value struct {
	constValueIdx  uint16
	constNameValue emunConstValue
	classInfoIndex uint16
}

type emunConstValue struct {
	typeNameIdx  uint16
	constNameIdx uint16
}

type arrayValue struct {
	valuesNum uint16
}

func (attr *AttrRuntimevisibleannotation) parse(reader *loader.ClassReader) {
	// annotations num
	reader.ReadInt16()
	// annotations array
	reader.ReadBytes(attr.attrLen - 2)
}

func (attr *AttrRuntimevisibleannotation) Name() string {
	return ""
}
