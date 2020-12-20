package instructions

import (
	"fmt"
	"github.com/zouzhihao-994/gvm/src/share/instructions/base"
	"github.com/zouzhihao-994/gvm/src/share/instructions/control"
	"github.com/zouzhihao-994/gvm/src/share/instructions/loads"
	"github.com/zouzhihao-994/gvm/src/share/instructions/references"
	"github.com/zouzhihao-994/gvm/src/share/instructions/stack"
	"github.com/zouzhihao-994/gvm/src/share/instructions/stores"
)

var (
	iload         = &loads.ILOAD{}
	lload         = &loads.LLOAD{}
	fload         = &loads.FLOAD{}
	dload         = &loads.DLOAD{}
	aload         = &loads.ALOAD{}
	iload_0       = &loads.ILOAD_0{}
	iload_1       = &loads.ILOAD_1{}
	iload_2       = &loads.ILOAD_2{}
	iload_3       = &loads.ILOAD_3{}
	lload_0       = &loads.LLOAD_0{}
	lload_1       = &loads.LLOAD_1{}
	lload_2       = &loads.LLOAD_2{}
	lload_3       = &loads.LLOAD_3{}
	fload_0       = &loads.FLOAD_0{}
	fload_1       = &loads.FLOAD_1{}
	fload_2       = &loads.FLOAD_2{}
	fload_3       = &loads.FLOAD_3{}
	dload_0       = &loads.DLOAD_0{}
	dload_1       = &loads.DLOAD_1{}
	dload_2       = &loads.DLOAD_2{}
	dload_3       = &loads.DLOAD_3{}
	aload_0       = &loads.ALOAD_0{}
	aload_1       = &loads.ALOAD_1{}
	aload_2       = &loads.ALOAD_2{}
	aload_3       = &loads.ALOAD_3{}
	istore        = &stores.ISTORE{}
	istore_0      = &stores.ISTORE_0{}
	istore_1      = &stores.ISTORE_1{}
	istore_2      = &stores.ISTORE_2{}
	istore_3      = &stores.ISTORE_3{}
	astore        = &stores.ASTORE{}
	astore_0      = &stores.ASTORE_0{}
	astore_1      = &stores.ASTORE_1{}
	astore_2      = &stores.ASTORE_2{}
	astore_3      = &stores.ASTORE_3{}
	fstore        = &stores.FSTORE{}
	fstore_0      = &stores.FSTORE_0{}
	fstore_1      = &stores.FSTORE_1{}
	fstore_2      = &stores.FSTORE_2{}
	fstore_3      = &stores.FSTORE_3{}
	dstore        = &stores.DSTORE{}
	dstore_0      = &stores.DSTORE_0{}
	dstore_1      = &stores.DSTORE_1{}
	dstore_2      = &stores.DSTORE_2{}
	dstore_3      = &stores.DSTORE_3{}
	lstore        = &stores.LSTORE{}
	lstore_0      = &stores.LSTORE_0{}
	lstore_1      = &stores.LSTORE_1{}
	lstore_2      = &stores.LSTORE_2{}
	lstore_3      = &stores.LSTORE_3{}
	_return       = &control.RETURN{}
	_ireturn      = &control.IRETURN{}
	_areturn      = &control.ARETURN{}
	_dreturn      = &control.DRETURN{}
	_lreturn      = &control.LRETURN{}
	_freturn      = &control.FRETURN{}
	dup           = &stack.Dup{}
	dup_x1        = &stack.Dup_X1{}
	dup_x2        = &stack.Dup_X2{}
	dup2          = &stack.Dup2{}
	dup2_x1       = &stack.Dup2_X1{}
	dup2_x2       = &stack.Dup2_X2{}
	getStatic     = &references.GET_STATIC{}    // 178
	invokeStatic  = &references.INVOKE_STATIC{} // 184
	invokeSpecial = &references.INVOKE_SPECIAL{}
	_new          = &references.NEW{} //187
)

func NewInstruction(opcode byte) base.Base_Instruction {
	switch opcode {
	case 0x15:
		return iload
	case 0x16:
		return lload
	case 0x17:
		return fload
	case 0x18:
		return dload
	case 0x19:
		return aload
	case 0x1a:
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
	case 0x36:
		return istore
	case 0x37:
		return lstore
	case 0x38:
		return fstore
	case 0x39:
		return dstore
	case 0x3a:
		return aload
	case 0x3b:
		return istore_0
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0xac:
		return _ireturn
	case 0xad:
		return _lreturn
	case 0xae:
		return _freturn
	case 0xaf:
		return _dreturn
	case 0xb0:
		return _areturn
	case 0xb1:
		return _return
	case 0xb2:
		return getStatic
	case 0xb7:
		return invokeSpecial
	case 0xb8:
		return invokeStatic
	case 0xbb:
		return _new
	default:
		panic(fmt.Errorf("Unsupported opcode : 0x%x!", opcode))
	}
}
