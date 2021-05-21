package oops

import (
	"github.com/zouzhihao-994/gvm/klass"
	"github.com/zouzhihao-994/gvm/utils"
)

type OopInstance struct {
	markWords     *MarkWords
	fields        *OopFields
	isArray       bool
	isString      bool
	jString       string
	jArray        *JArray
	klassInstance *klass.Klass
}

func (o *OopInstance) MarkWord() *MarkWords {
	return o.markWords
}

func (o *OopInstance) Klass() *klass.Klass {
	return o.klassInstance
}

func (o *OopInstance) ArrayLength() uint32 {
	utils.AssertTrue(o.isArray, "class is not array")
	return o.jArray.length
}

func (o *OopInstance) ArrayData() *JArray {
	return o.jArray
}

func (o *OopInstance) Fields() *OopFields {
	return o.fields
}

func (o *OopInstance) JString() string {
	return o.jString
}

// FindField find the field of oopInstance by field name
// n: field name
func (o *OopInstance) FindField(n string) (OopField, bool) {
	targetOop := o
	isSuper := false
	var f OopField
	for f, isSuper = targetOop.fields.GetField(n, isSuper); true != isSuper; {
		// todo: find from super
	}
	return f, true
}

// NewOopInstance create non-array oops
func NewOopInstance(k *klass.Klass) *OopInstance {
	return &OopInstance{
		markWords:     NewMarkWords(),
		fields:        InitOopFields(k),
		isArray:       false,
		klassInstance: k,
	}
}

// NewArrayOopInstance create array oops
func NewArrayOopInstance(arrayData *JArray) *OopInstance {
	return &OopInstance{
		markWords: NewMarkWords(),
		isArray:   true,
		jArray:    arrayData,
	}
}

func NewStringOopInstance(str string) *OopInstance {
	return &OopInstance{
		markWords: NewMarkWords(),
		isString:  true,
		jString:   str,
	}
}