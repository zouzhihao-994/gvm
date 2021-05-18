package attribute

import "github.com/zouzhihao-994/gvm/classloader"

type Attr_Synthetic struct {
	nameIdx uint16
	name    string
	attrLen uint32
}

func (attr Attr_Synthetic) parse(reader *classloader.ClassReader) {
	// nothing
}

func (attr Attr_Synthetic) Name() string {
	return ""
}
