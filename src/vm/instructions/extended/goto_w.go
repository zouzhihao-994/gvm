package extended

import (
	"../base"
)
import "../../rtda" // Branch always (wide index)

type GOTO_W struct{ offset int }

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	//fmt.Printf("[gvm][goto] goto nextPC offset : %v \n", self.offset)
	// 跳到对应的索引地址
	base.Branch(frame, self.offset)
}