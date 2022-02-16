package main

import "fmt"

type RuntimeConstantPool struct {
	infos            []RuntimeConstantPoolInfo
	bootstrapMethods []RBootstrapMethod
}

type RBootstrapMethod struct {
	bootstrapMethodRef RMethodHandle
	arguments          []RuntimeConstantPoolInfo
}

type RuntimeConstantPoolInfo interface {
	id() byte
}

const (
	RClassID        = 1
	RFieldRefID     = 2
	RMethodRefID    = 3
	RStringID       = 4
	RIntegerID      = 5
	RFloatID        = 6
	RLongID         = 7
	RDoubleID       = 8
	RNameAndTypeMID = 9
	RNameAndTypeFID = 10
	RUtf8ID         = 11
	RMethodHandleID = 12
	RMethodTypeID   = 13
	RDynamicID      = 14
	RModuleID       = 15
	RPackageID      = 16
)

func (_ RClass) id() byte {
	return RClassID
}

func (_ RFieldRef) id() byte {
	return RFieldRefID
}

func (_ RMethodRef) id() byte {
	return RMethodRefID
}

func (_ RString) id() byte {
	return RStringID
}

func (_ RInteger) id() byte {
	return RIntegerID
}

func (_ RFloat) id() byte {
	return RFloatID
}

func (_ RLong) id() byte {
	return RLongID
}

func (_ RDouble) id() byte {
	return RDoubleID
}

func (_ RNameAndTypeM) id() byte {
	return RNameAndTypeMID
}

func (_ RNameAndTypeF) id() byte {
	return RNameAndTypeFID
}

func (_ RUtf8) id() byte {
	return RUtf8ID
}

func (_ RMethodHandle) id() byte {
	return RMethodHandleID
}

func (_ RMethodType) id() byte {
	return RMethodTypeID
}

func (_ RDynamic) id() byte {
	return RDynamicID
}

func (_ RModule) id() byte {
	return RModuleID
}

func (_ RPackage) id() byte {
	return RPackageID
}

type RClass struct {
	RuntimeConstantPoolInfo

	name string
}

type RFieldRef struct {
	RuntimeConstantPoolInfo

	className  string
	name       string
	descriptor Type
}

type RMethodRef struct {
	RuntimeConstantPoolInfo

	className  string
	name       string
	descriptor MethodDescriptor
}

type RString struct {
	RuntimeConstantPoolInfo

	value string
}

type RInteger struct {
	RuntimeConstantPoolInfo

	value int32
}

type RFloat struct {
	RuntimeConstantPoolInfo

	value float32
}

type RLong struct {
	RuntimeConstantPoolInfo

	value int64
}

type RDouble struct {
	RuntimeConstantPoolInfo

	value float64
}

type RNameAndTypeM struct {
	RuntimeConstantPoolInfo

	name       string
	descriptor MethodDescriptor
}

type RNameAndTypeF struct {
	RuntimeConstantPoolInfo

	name       string
	descriptor Type
}

type RUtf8 struct {
	RuntimeConstantPoolInfo

	value string
}

type RMethodHandle struct {
	RuntimeConstantPoolInfo

	referenceKind  uint8
	referenceIndex uint16
}

type RMethodType struct {
	RuntimeConstantPoolInfo

	descriptor Type
}

type RDynamic struct {
	RuntimeConstantPoolInfo

	bootstrapMethodAttrIndex uint16
	name                     string
	descriptor               Type
}

type RModule struct {
	RuntimeConstantPoolInfo

	name string
}

type RPackage struct {
	RuntimeConstantPoolInfo

	name string
}

func Transform(class ClassFile) *RuntimeConstantPool {
	ocp := class.constantPool
	info := make([]RuntimeConstantPoolInfo, len(ocp))

	for i := 1; i < len(info); i++ {
		info[i] = transformInfo(ocp, ocp[i])
	}

	return &RuntimeConstantPool{
		infos:            info,
		bootstrapMethods: transformBootstrapMethods(class, info),
	}
}

func transformBootstrapMethods(class ClassFile, cp []RuntimeConstantPoolInfo) (bootstrapMethods []RBootstrapMethod) {
	bootstrapMethods = make([]RBootstrapMethod, 0)

	for _, attr := range class.attributes {
		if bm, ok := attr.attrType.(*BootstrapMethods); ok {
			for _, bm := range bm.bootstrapMethods {
				bootstrapMethods = append(bootstrapMethods, RBootstrapMethod{
					bootstrapMethodRef: cp[bm.bootstrapMethodRef].(RMethodHandle),
					arguments:          transformBootstrapArguments(cp, bm.bootstrapArguments),
				})
			}

			return
		}
	}

	return
}

func transformBootstrapArguments(cp []RuntimeConstantPoolInfo, bs []uint16) (arguments []RuntimeConstantPoolInfo) {
	arguments = make([]RuntimeConstantPoolInfo, len(bs))

	for i, b := range bs {
		arguments[i] = cp[b]
	}

	return
}

func transformInfo(cp []ClassPoolInfo, info ClassPoolInfo) RuntimeConstantPoolInfo {
	switch info := info.(type) {
	case *Class:
		return &RClass{
			name: AsString(cp[info.nameIndex]),
		}
	case *FieldRef:
		return &RFieldRef{
			className:  AsString(cp[info.classIndex]),
			name:       AsString(cp[(cp[info.nameAndTypeIndex]).(*NameAndType).nameIndex]),
			descriptor: ReadType(CreateStringConsumer(AsString(cp[(cp[info.nameAndTypeIndex]).(*NameAndType).descriptorIndex]))),
		}
	case *MethodRef:
		return &RMethodRef{
			className:  AsString(cp[info.classIndex]),
			name:       AsString(cp[(cp[info.nameAndTypeIndex]).(*NameAndType).nameIndex]),
			descriptor: MakeMethodDescriptor(CreateStringConsumer(AsString(cp[(cp[info.nameAndTypeIndex]).(*NameAndType).descriptorIndex]))),
		}
	case *InterfaceMethodRef:
		return &RMethodRef{
			className:  AsString(cp[info.classIndex]),
			name:       AsString(cp[(cp[info.nameAndTypeIndex]).(*NameAndType).nameIndex]),
			descriptor: MakeMethodDescriptor(CreateStringConsumer(AsString(cp[(cp[info.nameAndTypeIndex]).(*NameAndType).descriptorIndex]))),
		}
	case *String:
		return &RString{
			value: AsString(cp[info.stringIndex]),
		}
	case *Integer:
		return &RInteger{
			value: int32(info.bytes),
		}
	case *Float:
		return &RFloat{
			value: float32(info.bytes), // TODO rework
		}
	case *Long:
		return &RLong{
			value: (int64(info.highBytes) << 32) | int64(info.lowBytes),
		}
	case *Double:
		return &RDouble{
			value: float64((int64(info.highBytes) << 32) | int64(info.lowBytes)), // TODO rework
		}
	case *NameAndType:
		if (AsString(cp[info.descriptorIndex])[0]) == '(' {
			return &RNameAndTypeM{
				name:       AsString(cp[info.nameIndex]),
				descriptor: MakeMethodDescriptor(CreateStringConsumer(AsString(cp[info.descriptorIndex]))),
			}
		}

		return &RNameAndTypeF{
			name:       AsString(cp[info.nameIndex]),
			descriptor: ReadType(CreateStringConsumer(AsString(cp[info.descriptorIndex]))),
		}
	case *Utf8:
		return &RUtf8{
			value: string(info.bytes),
		}
	case *MethodHandle:
		return &RMethodHandle{
			referenceKind:  info.referenceKind,
			referenceIndex: info.referenceIndex,
		}
	case *MethodType:
		return &RMethodType{
			descriptor: ReadType(CreateStringConsumer(AsString(cp[info.descriptorIndex]))),
		}
	case *Dynamic:
		return &RDynamic{
			bootstrapMethodAttrIndex: info.bootstrapMethodAttrIndex,
			name:                     AsString(cp[(cp[info.nameAndTypeIndex]).(*NameAndType).nameIndex]),
			descriptor:               ReadType(CreateStringConsumer(AsString(cp[(cp[info.nameAndTypeIndex]).(*NameAndType).descriptorIndex]))),
		}
	case *ModuleClassPoolInfo:
		return &RModule{
			name: AsString(cp[info.nameIndex]),
		}
	case *Package:
		return &RPackage{
			name: AsString(cp[info.nameIndex]),
		}
	default:
		panic(fmt.Sprintf("Unknown constant pool info: %T", info))
	}
}
