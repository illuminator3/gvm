package main

import "fmt"

type Instruction interface {
	Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame)
}

type Instructionaaload struct { // 0x32
	Instruction
}
type Instructionaastore struct { // 0x53
	Instruction
}
type InstructionaconstNull struct { // 0x01
	Instruction
}
type Instructionaload struct { // 0x19
	Instruction

	Index byte
}
type Instructionaload0 struct { // 0x2a
	Instruction
}
type Instructionaload1 struct { // 0x2b
	Instruction
}
type Instructionaload2 struct { // 0x2c
	Instruction
}
type Instructionaload3 struct { // 0x2d
	Instruction
}
type Instructionanewarray struct { // 0xbd
	Instruction

	Index uint16
}
type Instructionareturn struct { // 0xb0
	Instruction
}
type Instructionarraylength struct { // 0xbe
	Instruction
}
type Instructionastore struct { // 0x3a
	Instruction

	Index byte
}
type Instructionastore0 struct { // 0x4b
	Instruction
}
type Instructionastore1 struct { // 0x4c
	Instruction
}
type Instructionastore2 struct { // 0x4d
	Instruction
}
type Instructionastore3 struct { // 0x4e
	Instruction
}
type Instructionathrow struct { // 0xbf
	Instruction
}
type Instructionbaload struct { // 0x33
	Instruction
}
type Instructionbastore struct { // 0x54
	Instruction
}
type Instructionbipush struct { // 0x10
	Instruction

	Value byte
}
type Instructionbreakpoint struct { // 0xca
	Instruction
}
type Instructioncaload struct { // 0x34
	Instruction
}
type Instructioncastore struct { // 0x55
	Instruction
}
type Instructioncheckcast struct { // 0xc0
	Instruction
	Index uint16
}
type Instructiond2f struct { // 0x90
	Instruction
}
type Instructiond2i struct { // 0x8e
	Instruction
}
type Instructiond2l struct { // 0x8f
	Instruction
}
type Instructiondadd struct { // 0x63
	Instruction
}
type Instructiondaload struct { // 0x31
	Instruction
}
type Instructiondastore struct { // 0x52
	Instruction
}
type Instructiondcmpg struct { // 0x98
	Instruction
}
type Instructiondcmpl struct { // 0x97
	Instruction
}
type Instructiondconst0 struct { // 0x0e
	Instruction
}
type Instructiondconst1 struct { // 0x0f
	Instruction
}
type Instructionddiv struct { // 0x6f
	Instruction
}
type Instructiondload struct { // 0x18
	Instruction
	Index byte
}
type Instructiondload0 struct { // 0x26
	Instruction
}
type Instructiondload1 struct { // 0x27
	Instruction
}
type Instructiondload2 struct { // 0x28
	Instruction
}
type Instructiondload3 struct { // 0x29
	Instruction
}
type Instructiondmul struct { // 0x6b
	Instruction
}
type Instructiondneg struct { // 0x77
	Instruction
}
type Instructiondrem struct { // 0x73
	Instruction
}
type Instructiondreturn struct { // 0xaf
	Instruction
}
type Instructiondstore struct { // 0x39
	Instruction

	Index byte
}
type Instructiondstore0 struct { // 0x47
	Instruction
}
type Instructiondstore1 struct { // 0x48
	Instruction
}
type Instructiondstore2 struct { // 0x49
	Instruction
}
type Instructiondstore3 struct { // 0x4a
	Instruction
}
type Instructiondsub struct { // 0x67
	Instruction
}
type Instructiondup struct { // 0x59
	Instruction
}
type InstructiondupX1 struct { // 0x5a
	Instruction
}
type InstructiondupX2 struct { // 0x5b
	Instruction
}
type Instructiondup2 struct { // 0x5c
	Instruction
}
type Instructiondup2X1 struct { // 0x5d
	Instruction
}
type Instructiondup2X2 struct { // 0x5e
	Instruction
}
type Instructionf2d struct { // 0x8d
	Instruction
}
type Instructionf2i struct { // 0x8b
	Instruction
}
type Instructionf2l struct { // 0x8c
	Instruction
}
type Instructionfadd struct { // 0x62
	Instruction
}
type Instructionfaload struct { // 0x30
	Instruction
}
type Instructionfastore struct { // 0x51
	Instruction
}
type Instructionfcmpg struct { // 0x96
	Instruction
}
type Instructionfcmpl struct { // 0x95
	Instruction
}
type Instructionfconst0 struct { // 0x0b
	Instruction
}
type Instructionfconst1 struct { // 0x0c
	Instruction
}
type Instructionfconst2 struct { // 0x0d
	Instruction
}
type Instructionfdiv struct { // 0x6e
	Instruction
}
type Instructionfload struct { // 0x17
	Instruction

	Index byte
}
type Instructionfload0 struct { // 0x22
	Instruction
}
type Instructionfload1 struct { // 0x23
	Instruction
}
type Instructionfload2 struct { // 0x24
	Instruction
}
type Instructionfload3 struct { // 0x25
	Instruction
}
type Instructionfmul struct { // 0x6a
	Instruction
}
type Instructionfneg struct { // 0x76
	Instruction
}
type Instructionfrem struct { // 0x72
	Instruction
}
type Instructionfreturn struct { // 0xae
	Instruction
}
type Instructionfstore struct { // 0x38
	Instruction

	Index byte
}
type Instructionfstore0 struct { // 0x43
	Instruction
}
type Instructionfstore1 struct { // 0x44
	Instruction
}
type Instructionfstore2 struct { // 0x45
	Instruction
}
type Instructionfstore3 struct { // 0x46
	Instruction
}
type Instructionfsub struct { // 0x66
	Instruction
}
type Instructiongetfield struct { // 0xb4
	Instruction

	Index uint16
}
type Instructiongetstatic struct { // 0xb2
	Instruction

	Index uint16
}
type Instructiongoto struct { // 0xa7
	Instruction

	OffsetShort int16
}
type InstructiongotoW struct { // 0xc8
	Instruction

	OffsetInt int32
}
type Instructioni2b struct { // 0x91
	Instruction
}
type Instructioni2c struct { // 0x92
	Instruction
}
type Instructioni2d struct { // 0x87
	Instruction
}
type Instructioni2f struct { // 0x86
	Instruction
}
type Instructioni2l struct { // 0x85
	Instruction
}
type Instructioni2s struct { // 0x93
	Instruction
}
type Instructioniadd struct { // 0x60
	Instruction
}
type Instructioniaload struct { // 0x2e
	Instruction
}
type Instructioniand struct { // 0x7e
	Instruction
}
type Instructioniastore struct { // 0x4f
	Instruction
}
type InstructioniconstM1 struct { // 0x02
	Instruction
}
type Instructioniconst0 struct { // 0x03
	Instruction
}
type Instructioniconst1 struct { // 0x04
	Instruction
}
type Instructioniconst2 struct { // 0x05
	Instruction
}
type Instructioniconst3 struct { // 0x06
	Instruction
}
type Instructioniconst4 struct { // 0x07
	Instruction
}
type Instructioniconst5 struct { // 0x08
	Instruction
}
type Instructionidiv struct { // 0x6c
	Instruction
}
type InstructionifAcmpeq struct { // 0xa5
	Instruction

	OffsetShort int16
}
type InstructionifAcmpne struct { // 0xa6
	Instruction

	OffsetShort int16
}
type InstructionifIcmpeq struct { // 0x9f
	Instruction

	OffsetShort int16
}
type InstructionifIcmpge struct { // 0xa2
	Instruction

	OffsetShort int16
}
type InstructionifIcmpgt struct { // 0xa3
	Instruction

	OffsetShort int16
}
type InstructionifIcmple struct { // 0xa4
	Instruction

	OffsetShort int16
}
type InstructionifIcmplt struct { // 0xa1
	Instruction

	OffsetShort int16
}
type InstructionifIcmpne struct { // 0xa0
	Instruction

	OffsetShort int16
}
type Instructionifeq struct { // 0x99
	Instruction

	OffsetShort int16
}
type Instructionifge struct { // 0x9c
	Instruction

	OffsetShort int16
}
type Instructionifgt struct { // 0x9d
	Instruction

	OffsetShort int16
}
type Instructionifle struct { // 0x9e
	Instruction

	OffsetShort int16
}
type Instructioniflt struct { // 0x9b
	Instruction

	OffsetShort int16
}
type Instructionifne struct { // 0x9a
	Instruction

	OffsetShort int16
}
type Instructionifnonnull struct { // 0xc7
	Instruction

	OffsetShort int16
}
type Instructionifnull struct { // 0xc6
	Instruction

	OffsetShort int16
}
type Instructioniinc struct { // 0x84
	Instruction

	Index byte
	Const byte
}
type Instructioniload struct { // 0x15
	Instruction

	Index byte
}
type Instructioniload0 struct { // 0x1a
	Instruction
}
type Instructioniload1 struct { // 0x1b
	Instruction
}
type Instructioniload2 struct { // 0x1c
	Instruction
}
type Instructioniload3 struct { // 0x1d
	Instruction
}
type Instructionimpdep1 struct { // 0xfe
	Instruction
}
type Instructionimpdep2 struct { // 0xff
	Instruction
}
type Instructionimul struct { // 0x68
	Instruction
}
type Instructionineg struct { // 0x74
	Instruction
}
type Instructioninstanceof struct { // 0xc1
	Instruction

	Index uint16
}
type Instructioninvokedynamic struct { // 0xba
	Instruction

	Index uint16
}
type Instructioninvokeinterface struct { // 0xb9
	Instruction

	Index uint16
	Count byte
}
type Instructioninvokespecial struct { // 0xb7
	Instruction

	Index uint16
}
type Instructioninvokestatic struct { // 0xb8
	Instruction

	Index uint16
}
type Instructioninvokevirtual struct { // 0xb6
	Instruction

	Index uint16
}
type Instructionior struct { // 0x80
	Instruction
}
type Instructionirem struct { // 0x70
	Instruction
}
type Instructionireturn struct { // 0xac
	Instruction
}
type Instructionishl struct { // 0x78
	Instruction
}
type Instructionishr struct { // 0x7a
	Instruction
}
type Instructionistore struct { // 0x36
	Instruction

	Index byte
}
type Instructionistore0 struct { // 0x3b
	Instruction
}
type Instructionistore1 struct { // 0x3c
	Instruction
}
type Instructionistore2 struct { // 0x3d
	Instruction
}
type Instructionistore3 struct { // 0x3e
	Instruction
}
type Instructionisub struct { // 0x64
	Instruction
}
type Instructioniushr struct { // 0x7c
	Instruction
}
type Instructionixor struct { // 0x82
	Instruction
}
type Instructionjsr struct { // 0xa8
	Instruction

	OffsetShort int16
}
type InstructionjsrW struct { // 0xc9
	Instruction

	OffsetInt int32
}
type Instructionl2d struct { // 0x8a
	Instruction
}
type Instructionl2f struct { // 0x89
	Instruction
}
type Instructionl2i struct { // 0x88
	Instruction
}
type Instructionladd struct { // 0x61
	Instruction
}
type Instructionlaload struct { // 0x2f
	Instruction
}
type Instructionland struct { // 0x7f
	Instruction
}
type Instructionlastore struct { // 0x50
	Instruction
}
type Instructionlcmp struct { // 0x94
	Instruction
}
type Instructionlconst0 struct { // 0x09
	Instruction
}
type Instructionlconst1 struct { // 0x0a
	Instruction
}
type Instructionldc struct { // 0x12
	Instruction

	Index byte
}
type InstructionldcW struct { // 0x13
	Instruction

	Index uint16
}
type Instructionldc2W struct { // 0x14
	Instruction

	Index uint16
}
type Instructionldiv struct { // 0x6d
	Instruction
}
type Instructionlload struct { // 0x16
	Instruction

	Index byte
}
type Instructionlload0 struct { // 0x1e
	Instruction
}
type Instructionlload1 struct { // 0x1f
	Instruction
}
type Instructionlload2 struct { // 0x20
	Instruction
}
type Instructionlload3 struct { // 0x21
	Instruction
}
type Instructionlmul struct { // 0x69
	Instruction
}
type Instructionlneg struct { // 0x75
	Instruction
}
type Instructionlookupswitch struct { // 0xab
	Instruction

	Default int32
	Pairs   []MatchOffsetPair
}

type MatchOffsetPair struct {
	Match  int32
	Offset int32
}

type Instructionlor struct { // 0x81
	Instruction
}
type Instructionlrem struct { // 0x71
	Instruction
}
type Instructionlreturn struct { // 0xad
	Instruction
}
type Instructionlshl struct { // 0x79
	Instruction
}
type Instructionlshr struct { // 0x7b
	Instruction
}
type Instructionlstore struct { // 0x37
	Instruction

	Index byte
}
type Instructionlstore0 struct { // 0x3f
	Instruction
}
type Instructionlstore1 struct { // 0x40
	Instruction
}
type Instructionlstore2 struct { // 0x41
	Instruction
}
type Instructionlstore3 struct { // 0x42
	Instruction
}
type Instructionlsub struct { // 0x65
	Instruction
}
type Instructionlushr struct { // 0x7d
	Instruction
}
type Instructionlxor struct { // 0x83
	Instruction
}
type Instructionmonitorenter struct { // 0xc2
	Instruction
}
type Instructionmonitorexit struct { // 0xc3
	Instruction
}
type Instructionmultianewarray struct { // 0xc5
	Instruction

	Index      uint16
	Dimensions byte
}
type Instructionnew struct { // 0xbb
	Instruction

	Index uint16
}
type Instructionnewarray struct { // 0xbc
	Instruction

	Atype byte
}
type Instructionnop struct { // 0x00
	Instruction
}
type Instructionpop struct { // 0x57
	Instruction
}
type Instructionpop2 struct { // 0x58
	Instruction
}
type Instructionputfield struct { // 0xb5
	Instruction

	Index uint16
}
type Instructionputstatic struct { // 0xb3
	Instruction

	Index uint16
}
type Instructionret struct { // 0xa9
	Instruction

	Index byte
}
type Instructionreturn struct { // 0xb1
	Instruction
}
type Instructionsaload struct { // 0x35
	Instruction
}
type Instructionsastore struct { // 0x53
	Instruction
}
type Instructionsipush struct { // 0x15
	Instruction

	Value int16
}
type Instructionswap struct { // 0x5f
	Instruction
}
type Instructiontableswitch struct { // 0xaa
	Instruction

	Default     int32
	Low         int32
	High        int32
	JumpOffsets []int32
}
type Instructionwide struct { // 0xc4
	Instruction

	Opcode byte
}

func ParseInstructions(bytes []byte) (instructions []Instruction) {
	container := ByteContainer{bytes, 0}

	for container.HasNext() {
		instructions = append(instructions, ParseInstruction(&container))
	}

	return
}

func ParseInstruction(bytes *ByteContainer) Instruction {
	opcode := bytes.NextByte()

	switch opcode {
	case 0x00:
		return Instructionnop{}
	case 0x01:
		return InstructionaconstNull{}
	case 0x02:
		return InstructioniconstM1{}
	case 0x03:
		return Instructioniconst0{}
	case 0x04:
		return Instructioniconst1{}
	case 0x05:
		return Instructioniconst2{}
	case 0x06:
		return Instructioniconst3{}
	case 0x07:
		return Instructioniconst4{}
	case 0x08:
		return Instructioniconst5{}
	case 0x09:
		return Instructionlconst0{}
	case 0x0a:
		return Instructionlconst1{}
	case 0x0b:
		return Instructionfconst0{}
	case 0x0c:
		return Instructionfconst1{}
	case 0x0d:
		return Instructionfconst2{}
	case 0x0e:
		return Instructiondconst0{}
	case 0x0f:
		return Instructiondconst1{}
	case 0x10:
		return Instructionbipush{
			Value: bytes.NextByte(),
		}
	case 0x11:
		return Instructionsipush{
			Value: bytes.Nexti16(),
		}
	case 0x12:
		return Instructionldc{
			Index: bytes.NextByte(),
		}
	case 0x13:
		return InstructionldcW{
			Index: bytes.Nextu16(),
		}
	case 0x14:
		return Instructionldc2W{
			Index: bytes.Nextu16(),
		}
	case 0x15:
		return Instructioniload{
			Index: bytes.NextByte(),
		}
	case 0x16:
		return Instructionlload{
			Index: bytes.NextByte(),
		}
	case 0x17:
		return Instructionfload{
			Index: bytes.NextByte(),
		}
	case 0x18:
		return Instructiondload{
			Index: bytes.NextByte(),
		}
	case 0x19:
		return Instructionaload{
			Index: bytes.NextByte(),
		}
	case 0x1a:
		return Instructioniload0{}
	case 0x1b:
		return Instructioniload1{}
	case 0x1c:
		return Instructioniload2{}
	case 0x1d:
		return Instructioniload3{}
	case 0x1e:
		return Instructionlload0{}
	case 0x1f:
		return Instructionlload1{}
	case 0x20:
		return Instructionlload2{}
	case 0x21:
		return Instructionlload3{}
	case 0x22:
		return Instructionfload0{}
	case 0x23:
		return Instructionfload1{}
	case 0x24:
		return Instructionfload2{}
	case 0x25:
		return Instructionfload3{}
	case 0x26:
		return Instructiondload0{}
	case 0x27:
		return Instructiondload1{}
	case 0x28:
		return Instructiondload2{}
	case 0x29:
		return Instructiondload3{}
	case 0x2a:
		return Instructionaload0{}
	case 0x2b:
		return Instructionaload1{}
	case 0x2c:
		return Instructionaload2{}
	case 0x2d:
		return Instructionaload3{}
	case 0x2e:
		return Instructioniaload{}
	case 0x2f:
		return Instructionlaload{}
	case 0x30:
		return Instructionfaload{}
	case 0x31:
		return Instructiondaload{}
	case 0x32:
		return Instructionaaload{}
	case 0x33:
		return Instructionbaload{}
	case 0x34:
		return Instructioncaload{}
	case 0x35:
		return Instructionsaload{}
	case 0x36:
		return Instructionistore{
			Index: bytes.NextByte(),
		}
	case 0x37:
		return Instructionlstore{
			Index: bytes.NextByte(),
		}
	case 0x38:
		return Instructionfstore{
			Index: bytes.NextByte(),
		}
	case 0x39:
		return Instructiondstore{
			Index: bytes.NextByte(),
		}
	case 0x3a:
		return Instructionastore{
			Index: bytes.NextByte(),
		}
	case 0x3b:
		return Instructionistore0{}
	case 0x3c:
		return Instructionistore1{}
	case 0x3d:
		return Instructionistore2{}
	case 0x3e:
		return Instructionistore3{}
	case 0x3f:
		return Instructionlstore0{}
	case 0x40:
		return Instructionlstore1{}
	case 0x41:
		return Instructionlstore2{}
	case 0x42:
		return Instructionlstore3{}
	case 0x43:
		return Instructionfstore0{}
	case 0x44:
		return Instructionfstore1{}
	case 0x45:
		return Instructionfstore2{}
	case 0x46:
		return Instructionfstore3{}
	case 0x47:
		return Instructiondstore0{}
	case 0x48:
		return Instructiondstore1{}
	case 0x49:
		return Instructiondstore2{}
	case 0x4a:
		return Instructiondstore3{}
	case 0x4b:
		return Instructionastore0{}
	case 0x4c:
		return Instructionastore1{}
	case 0x4d:
		return Instructionastore2{}
	case 0x4e:
		return Instructionastore3{}
	case 0x4f:
		return Instructioniastore{}
	case 0x50:
		return Instructionlastore{}
	case 0x51:
		return Instructionfastore{}
	case 0x52:
		return Instructiondastore{}
	case 0x53:
		return Instructionaastore{}
	case 0x54:
		return Instructionbastore{}
	case 0x55:
		return Instructioncastore{}
	case 0x56:
		return Instructionsastore{}
	case 0x57:
		return Instructionpop{}
	case 0x58:
		return Instructionpop2{}
	case 0x59:
		return Instructiondup{}
	case 0x5a:
		return InstructiondupX1{}
	case 0x5b:
		return InstructiondupX2{}
	case 0x5c:
		return Instructiondup2{}
	case 0x5d:
		return Instructiondup2X1{}
	case 0x5e:
		return Instructiondup2X2{}
	case 0x5f:
		return Instructionswap{}
	case 0x60:
		return Instructioniadd{}
	case 0x61:
		return Instructionladd{}
	case 0x62:
		return Instructionfadd{}
	case 0x63:
		return Instructiondadd{}
	case 0x64:
		return Instructionisub{}
	case 0x65:
		return Instructionlsub{}
	case 0x66:
		return Instructionfsub{}
	case 0x67:
		return Instructiondsub{}
	case 0x68:
		return Instructionimul{}
	case 0x69:
		return Instructionlmul{}
	case 0x6a:
		return Instructionfmul{}
	case 0x6b:
		return Instructiondmul{}
	case 0x6c:
		return Instructionidiv{}
	case 0x6d:
		return Instructionldiv{}
	case 0x6e:
		return Instructionfdiv{}
	case 0x6f:
		return Instructionddiv{}
	case 0x70:
		return Instructionirem{}
	case 0x71:
		return Instructionlrem{}
	case 0x72:
		return Instructionfrem{}
	case 0x73:
		return Instructiondrem{}
	case 0x74:
		return Instructionineg{}
	case 0x75:
		return Instructionlneg{}
	case 0x76:
		return Instructionfneg{}
	case 0x77:
		return Instructiondneg{}
	case 0x78:
		return Instructionishl{}
	case 0x79:
		return Instructionlshl{}
	case 0x7a:
		return Instructionishr{}
	case 0x7b:
		return Instructionlshr{}
	case 0x7c:
		return Instructioniushr{}
	case 0x7d:
		return Instructionlushr{}
	case 0x7e:
		return Instructioniand{}
	case 0x7f:
		return Instructionland{}
	case 0x80:
		return Instructionior{}
	case 0x81:
		return Instructionlor{}
	case 0x82:
		return Instructionixor{}
	case 0x83:
		return Instructionlxor{}
	case 0x84:
		return Instructioniinc{}
	case 0x85:
		return Instructioni2l{}
	case 0x86:
		return Instructioni2f{}
	case 0x87:
		return Instructioni2d{}
	case 0x88:
		return Instructionl2i{}
	case 0x89:
		return Instructionl2f{}
	case 0x8a:
		return Instructionl2d{}
	case 0x8b:
		return Instructionf2i{}
	case 0x8c:
		return Instructionf2l{}
	case 0x8d:
		return Instructionf2d{}
	case 0x8e:
		return Instructiond2i{}
	case 0x8f:
		return Instructiond2l{}
	case 0x90:
		return Instructiond2f{}
	case 0x91:
		return Instructioni2b{}
	case 0x92:
		return Instructioni2c{}
	case 0x93:
		return Instructioni2s{}
	case 0x94:
		return Instructionlcmp{}
	case 0x95:
		return Instructionfcmpl{}
	case 0x96:
		return Instructionfcmpg{}
	case 0x97:
		return Instructiondcmpl{}
	case 0x98:
		return Instructiondcmpg{}
	case 0x99:
		return Instructionifeq{
			OffsetShort: bytes.Nexti16(),
		}
	case 0x9a:
		return Instructionifne{
			OffsetShort: bytes.Nexti16(),
		}
	case 0x9b:
		return Instructioniflt{
			OffsetShort: bytes.Nexti16(),
		}
	case 0x9c:
		return Instructionifge{
			OffsetShort: bytes.Nexti16(),
		}
	case 0x9d:
		return Instructionifgt{
			OffsetShort: bytes.Nexti16(),
		}
	case 0x9e:
		return Instructionifle{
			OffsetShort: bytes.Nexti16(),
		}
	case 0x9f:
		return InstructionifIcmpeq{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xa0:
		return InstructionifIcmpne{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xa1:
		return InstructionifIcmplt{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xa2:
		return InstructionifIcmpge{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xa3:
		return InstructionifIcmpgt{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xa4:
		return InstructionifIcmple{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xa5:
		return InstructionifAcmpeq{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xa6:
		return InstructionifAcmpne{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xa7:
		return Instructiongoto{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xa8:
		return Instructionjsr{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xa9:
		return Instructionret{
			Index: bytes.NextByte(),
		}
	case 0xaa:
		bytes.NextBytes(bytes.BytesRead() % 4) // padding

		defaultOffset := bytes.Nexti32()
		low := bytes.Nexti32()
		high := bytes.Nexti32()
		jumpOffsets := bytes.Nexti32s(int(high - low + 1))

		return Instructiontableswitch{
			Default:     defaultOffset,
			Low:         low,
			High:        high,
			JumpOffsets: jumpOffsets,
		}
	case 0xab:
		bytes.NextBytes(bytes.BytesRead() % 4) // padding

		defaultOffset := bytes.Nexti32()
		npairs := bytes.Nexti32()
		pairs := make([]MatchOffsetPair, npairs)

		for i := 0; i < int(npairs); i++ {
			pairs[i].Match = bytes.Nexti32()
			pairs[i].Offset = bytes.Nexti32()
		}

		return Instructionlookupswitch{
			Default: defaultOffset,
			Pairs:   pairs,
		}
	case 0xac:
		return Instructionireturn{}
	case 0xad:
		return Instructionlreturn{}
	case 0xae:
		return Instructionfreturn{}
	case 0xaf:
		return Instructiondreturn{}
	case 0xb0:
		return Instructionareturn{}
	case 0xb1:
		return Instructionreturn{}
	case 0xb2:
		return Instructiongetstatic{
			Index: bytes.Nextu16(),
		}
	case 0xb3:
		return Instructionputstatic{
			Index: bytes.Nextu16(),
		}
	case 0xb4:
		return Instructiongetfield{
			Index: bytes.Nextu16(),
		}
	case 0xb5:
		return Instructionputfield{
			Index: bytes.Nextu16(),
		}
	case 0xb6:
		return Instructioninvokevirtual{
			Index: bytes.Nextu16(),
		}
	case 0xb7:
		return Instructioninvokespecial{
			Index: bytes.Nextu16(),
		}
	case 0xb8:
		return Instructioninvokestatic{
			Index: bytes.Nextu16(),
		}
	case 0xb9:
		return Instructioninvokeinterface{
			Index: bytes.Nextu16(),
			Count: bytes.NextByte(),
		}
	case 0xba:
		return Instructioninvokedynamic{
			Index: bytes.Nextu16(),
		}
	case 0xbb:
		return Instructionnew{
			Index: bytes.Nextu16(),
		}
	case 0xbc:
		return Instructionnewarray{
			Atype: bytes.NextByte(),
		}
	case 0xbd:
		return Instructionanewarray{
			Index: bytes.Nextu16(),
		}
	case 0xbe:
		return Instructionarraylength{}
	case 0xbf:
		return Instructionathrow{}
	case 0xc0:
		return Instructioncheckcast{
			Index: bytes.Nextu16(),
		}
	case 0xc1:
		return Instructioninstanceof{
			Index: bytes.Nextu16(),
		}
	case 0xc2:
		return Instructionmonitorenter{}
	case 0xc3:
		return Instructionmonitorexit{}
	case 0xc4:
		// TODO read this stupid instruction

		return Instructionwide{
			Opcode: bytes.NextByte(),
		}
	case 0xc5:
		return Instructionmultianewarray{
			Index: bytes.Nextu16(),
		}
	case 0xc6:
		return Instructionifnull{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xc7:
		return Instructionifnonnull{
			OffsetShort: bytes.Nexti16(),
		}
	case 0xc8:
		return InstructiongotoW{
			OffsetInt: bytes.Nexti32(),
		}
	case 0xc9:
		return InstructionjsrW{
			OffsetInt: bytes.Nexti32(),
		}
	case 0xca:
		return Instructionbreakpoint{}
	case 0xfe:
		return Instructionimpdep1{}
	case 0xff:
		return Instructionimpdep2{}
	default:
		panic(fmt.Sprintf("Unknown opcode: %x", opcode))
	}
}

type ByteContainer struct {
	bytes     []byte
	bytesRead int
}

func (r *ByteContainer) BytesRead() int {
	return r.bytesRead
}

func (r *ByteContainer) HasNext() bool {
	return len(r.bytes) > 0
}

func (r *ByteContainer) NextByte() (b byte) {
	b = r.bytes[0]

	r.bytes = r.bytes[1:]
	r.bytesRead++

	return
}

func (r *ByteContainer) NextBytes(n int) (bytes []byte) {
	bytes = r.bytes[:n]

	r.bytes = r.bytes[n:]
	r.bytesRead += n

	return
}

func (r *ByteContainer) Nextu16() (u16 uint16) {
	u16 = uint16(r.bytes[0])<<8 | uint16(r.bytes[1])

	r.bytes = r.bytes[2:]
	r.bytesRead += 2

	return
}

func (r *ByteContainer) Nexti16() (i16 int16) {
	i16 = int16(r.bytes[0])<<8 | int16(r.bytes[1])

	r.bytes = r.bytes[2:]
	r.bytesRead += 2

	return
}

func (r *ByteContainer) Nexti16s(n int) (i16s []int16) {
	i16s = make([]int16, n)

	for i := 0; i < n; i++ {
		i16s[i] = r.Nexti16()
	}

	return
}

func (r *ByteContainer) Nextu32() (u32 uint32) {
	u32 = uint32(r.bytes[0])<<24 | uint32(r.bytes[1])<<16 | uint32(r.bytes[2])<<8 | uint32(r.bytes[3])

	r.bytes = r.bytes[4:]
	r.bytesRead += 4

	return
}

func (r *ByteContainer) Nexti32() (i32 int32) {
	i32 = int32(r.bytes[0])<<24 | int32(r.bytes[1])<<16 | int32(r.bytes[2])<<8 | int32(r.bytes[3])

	r.bytes = r.bytes[4:]
	r.bytesRead += 4

	return
}

func (r *ByteContainer) Nexti32s(n int) (i32s []int32) {
	i32s = make([]int32, n)

	for i := 0; i < n; i++ {
		i32s[i] = r.Nexti32()
	}

	return
}

func (r *ByteContainer) Nextu64() (u64 uint64) {
	u64 = uint64(r.bytes[0])<<56 | uint64(r.bytes[1])<<48 | uint64(r.bytes[2])<<40 | uint64(r.bytes[3])<<32 | uint64(r.bytes[4])<<24 | uint64(r.bytes[5])<<16 | uint64(r.bytes[6])<<8 | uint64(r.bytes[7])

	r.bytes = r.bytes[8:]
	r.bytesRead += 8

	return
}

func (r *ByteContainer) Nexti64() (i64 int64) {
	i64 = int64(r.bytes[0])<<56 | int64(r.bytes[1])<<48 | int64(r.bytes[2])<<40 | int64(r.bytes[3])<<32 | int64(r.bytes[4])<<24 | int64(r.bytes[5])<<16 | int64(r.bytes[6])<<8 | int64(r.bytes[7])

	r.bytes = r.bytes[8:]
	r.bytesRead += 8

	return
}
