package main

import (
	"encoding/binary"
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool []ClassPoolInfo
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []MemberInfo
	methods      []MemberInfo
	attributes   []AttributeInfo
}

type ClassPoolInfo interface{}

type Class struct {
	ClassPoolInfo

	nameIndex uint16
}

type FieldRef struct {
	ClassPoolInfo

	classIndex       uint16
	nameAndTypeIndex uint16
}

type MethodRef struct {
	ClassPoolInfo

	classIndex       uint16
	nameAndTypeIndex uint16
}

type InterfaceMethodRef struct {
	ClassPoolInfo

	classIndex       uint16
	nameAndTypeIndex uint16
}

type String struct {
	ClassPoolInfo

	stringIndex uint16
}

type Integer struct {
	ClassPoolInfo

	bytes uint32
}

type Float struct {
	ClassPoolInfo

	bytes uint32
}

type Long struct {
	ClassPoolInfo

	highBytes uint32
	lowBytes  uint32
}

type Double struct {
	ClassPoolInfo

	highBytes uint32
	lowBytes  uint32
}

type NameAndType struct {
	ClassPoolInfo

	nameIndex       uint16
	descriptorIndex uint16
}

type Utf8 struct {
	ClassPoolInfo

	length uint16
	bytes  []uint8
}

type MethodHandle struct {
	ClassPoolInfo

	referenceKind  uint8
	referenceIndex uint16
}

type MethodType struct {
	ClassPoolInfo

	descriptorIndex uint16
}

type Dynamic struct {
	ClassPoolInfo

	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

type InvokeDynamic struct {
	ClassPoolInfo

	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

type ModuleClassPoolInfo struct {
	ClassPoolInfo

	nameIndex uint16
}

type Package struct {
	ClassPoolInfo

	nameIndex uint16
}

type UnknownClassPoolInfo struct {
	ClassPoolInfo

	tag uint8
}

type ClassReader struct {
	data []uint8
}

func (reader *ClassReader) ReadUint8() uint8 {
	val := reader.data[0]
	reader.data = reader.data[1:]
	return val
}

func (reader *ClassReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(reader.data)
	reader.data = reader.data[2:]
	return val
}

func (reader *ClassReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(reader.data)
	reader.data = reader.data[4:]
	return val
}

func (reader *ClassReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(reader.data)
	reader.data = reader.data[8:]
	return val
}

func (reader *ClassReader) ReadUint8s(n uint64) []uint8 {
	val := reader.data[:n]
	reader.data = reader.data[n:]
	return val
}

func (reader *ClassReader) ReadUint16s(n uint64) []uint16 {
	val := make([]uint16, n)
	for i := uint64(0); i < n; i++ {
		val[i] = reader.ReadUint16()
	}
	return val
}

func (reader *ClassReader) ReadUint32s(n uint64) []uint32 {
	val := make([]uint32, n)
	for i := uint64(0); i < n; i++ {
		val[i] = reader.ReadUint32()
	}
	return val
}

func (reader *ClassReader) ReadUint64s(n uint64) []uint64 {
	val := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		val[i] = reader.ReadUint64()
	}
	return val
}

func (reader *ClassReader) ReadClassPoolInfo() ClassPoolInfo {
	tag := reader.ReadUint8()
	switch tag {
	case 7:
		return &Class{
			nameIndex: reader.ReadUint16(),
		}
	case 9:
		return &FieldRef{
			classIndex:       reader.ReadUint16(),
			nameAndTypeIndex: reader.ReadUint16(),
		}
	case 10:
		return &MethodRef{
			classIndex:       reader.ReadUint16(),
			nameAndTypeIndex: reader.ReadUint16(),
		}
	case 11:
		return &InterfaceMethodRef{
			classIndex:       reader.ReadUint16(),
			nameAndTypeIndex: reader.ReadUint16(),
		}
	case 8:
		return &String{
			stringIndex: reader.ReadUint16(),
		}
	case 3:
		return &Integer{
			bytes: reader.ReadUint32(),
		}
	case 4:
		return &Float{
			bytes: reader.ReadUint32(),
		}
	case 5:
		return &Long{
			highBytes: reader.ReadUint32(),
			lowBytes:  reader.ReadUint32(),
		}
	case 6:
		return &Double{
			highBytes: reader.ReadUint32(),
			lowBytes:  reader.ReadUint32(),
		}
	case 12:
		return &NameAndType{
			nameIndex:       reader.ReadUint16(),
			descriptorIndex: reader.ReadUint16(),
		}
	case 1:
		length := reader.ReadUint16()
		bytes := reader.ReadUint8s(uint64(length))
		return &Utf8{
			length: length,
			bytes:  bytes,
		}
	case 15:
		return &MethodHandle{
			referenceKind:  reader.ReadUint8(),
			referenceIndex: reader.ReadUint16(),
		}
	case 16:
		return &MethodType{
			descriptorIndex: reader.ReadUint16(),
		}
	case 17:
		return &Dynamic{
			bootstrapMethodAttrIndex: reader.ReadUint16(),
			nameAndTypeIndex:         reader.ReadUint16(),
		}
	case 18:
		return &InvokeDynamic{
			bootstrapMethodAttrIndex: reader.ReadUint16(),
			nameAndTypeIndex:         reader.ReadUint16(),
		}
	case 19:
		return &ModuleClassPoolInfo{
			nameIndex: reader.ReadUint16(),
		}
	case 20:
		return &Package{
			nameIndex: reader.ReadUint16(),
		}
	default:
		return &UnknownClassPoolInfo{
			tag: tag,
		}
	}
}

func AsString(info ClassPoolInfo) string {
	if utf8, ok := info.(*Utf8); ok {
		return string(utf8.bytes)
	}

	return ""
}

type MemberInfo struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	attributes      []AttributeInfo
}

type AttributeInfo struct {
	attributeNameIndex uint16
	attributeLength    uint32
	attrType           AttributeType
}

type AttributeType interface{}

type ConstantValue struct {
	AttributeType

	constantValueIndex uint16
}

type Code struct {
	AttributeType

	maxStack       uint16
	maxLocals      uint16
	codeLength     uint32
	code           []uint8
	exceptionTable []ExceptionTableEntry
	attributes     []AttributeInfo
}

type StackMapTable struct {
	AttributeType

	numberOfEntries uint16
	entries         []StackMapFrame
}

type StackMapFrame interface{}

type SameFrame struct {
	StackMapFrame

	frameType uint8
}

type SameLocals1StackItemFrame struct {
	StackMapFrame

	frameType uint8
	stack     VerificationTypeInfo
}

type SameLocals1StackItemFrameExtended struct {
	StackMapFrame

	frameType   uint8
	offsetDelta uint16
	stack       VerificationTypeInfo
}

type ChopFrame struct {
	StackMapFrame

	frameType   uint8
	offsetDelta uint16
}

type SameFrameExtended struct {
	StackMapFrame

	frameType   uint8
	offsetDelta uint16
}

type AppendFrame struct {
	StackMapFrame

	frameType   uint8
	offsetDelta uint16
	locals      []VerificationTypeInfo
}

type FullFrame struct {
	StackMapFrame

	frameType          uint8
	offsetDelta        uint16
	numberOfLocals     uint16
	locals             []VerificationTypeInfo
	numberOfStackItems uint16
	stack              []VerificationTypeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

type VerificationTypeInfo interface{}

type TopVariable struct {
	VerificationTypeInfo
}

type IntegerVariable struct {
	VerificationTypeInfo
}

type FloatVariable struct {
	VerificationTypeInfo
}

type NullVariable struct {
	VerificationTypeInfo
}

type UninitializedThisVariable struct {
	VerificationTypeInfo
}

type ObjectVariable struct {
	VerificationTypeInfo

	cpoolIndex uint16
}

type UninitializedVariable struct {
	VerificationTypeInfo

	offset uint16
}

type LongVariable struct {
	VerificationTypeInfo
}

type DoubleVariable struct {
	VerificationTypeInfo
}

type Exceptions struct {
	AttributeType

	numberOfExceptions uint16
	exceptionIndices   []uint16
}

type InnerClasses struct {
	AttributeType

	numberOfClasses uint16
	classes         []InnerClass
}

type InnerClass struct {
	innerClassInfoIndex   uint16
	outerClassInfoIndex   uint16
	innerNameIndex        uint16
	innerClassAccessFlags uint16
}

type EnclosingMethod struct {
	AttributeType

	classIndex  uint16
	methodIndex uint16
}

type Synthetic struct {
	AttributeType
}

type Signature struct {
	AttributeType

	signatureIndex uint16
}

type SourceFile struct {
	AttributeType

	sourceFileIndex uint16
}

type SourceDebugExtension struct {
	AttributeType

	debugExtension []uint8
}

type LineNumberTable struct {
	AttributeType

	lineNumberTable []LineNumber
}

type LineNumber struct {
	startPc    uint16
	lineNumber uint16
}

type LocalVariableTable struct {
	AttributeType

	localVariableTable []LocalVariable
}

type LocalVariable struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

type LocalVariableTypeTable struct {
	AttributeType

	localVariableTypeTable []LocalVariableType
}

type LocalVariableType struct {
	startPc        uint16
	length         uint16
	nameIndex      uint16
	signatureIndex uint16
	index          uint16
}

type Deprecated struct {
	AttributeType
}

type RuntimeVisibleAnnotations struct {
	AttributeType

	annotations []Annotation
}

type Annotation struct {
	typeIndex    uint16
	elementPairs []ElementPair
}

type ElementPair struct {
	elementNameIndex uint16
	value            ElementValue
}

type ElementValue struct {
	tag   uint8
	value ElementValueInner
}

type ElementValueInner interface{}

type ConstantValueIndex struct {
	ElementValueInner

	constantValueIndex uint16
}

type EnumConstantValue struct {
	ElementValueInner

	typeNameIndex     uint16
	constantNameIndex uint16
}

type ClassInfoIndex struct {
	ElementValueInner

	classInfoIndex uint16
}

type AnnotationValue struct {
	ElementValueInner

	annotationValue []Annotation
}

type ArrayValue struct {
	ElementValueInner

	values []ElementValue
}

type RuntimeInvisibleAnnotations struct {
	AttributeType

	annotations []Annotation
}

type RuntimeInvisibleParameterAnnotations struct {
	AttributeType

	parameterAnnotations []ParameterAnnotation
}

type ParameterAnnotation struct {
	parameterIndex uint8
	annotations    []Annotation
}

type RuntimeVisibleParameterAnnotations struct {
	AttributeType

	parameterAnnotations []ParameterAnnotation
}

type AnnotationDefault struct {
	AttributeType

	defaultValue ElementValue
}

type BootstrapMethods struct {
	AttributeType

	numberOfBootstrapMethods uint16
	bootstrapMethods         []BootstrapMethod
}

type BootstrapMethod struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}

type MethodParameters struct {
	AttributeType

	numberOfParameters uint8
	parameters         []MethodParameter
}

type MethodParameter struct {
	nameIndex   uint16
	accessFlags uint16
}

type Module struct {
	AttributeType

	moduleNameIndex    uint16
	moduleFlags        uint16
	moduleVersionIndex uint16
	moduleInfo         ModuleInfo
}

type ModuleInfo struct {
	requiresCount uint16
	requires      []ModuleRequire
	exportsCount  uint16
	exports       []ModuleExport
	opensCount    uint16
	opens         []ModuleOpen
	usesCount     uint16
	usesIndex     []uint16
	providesCount uint16
	provides      []ModuleProvide
}

type ModuleRequire struct {
	requiresIndex        uint16
	requiresFlags        uint16
	requiresVersionIndex uint16
}

type ModuleExport struct {
	exportsIndex   uint16
	exportsFlags   uint16
	exportsToCount uint16
	exportsToIndex []uint16
}

type ModuleOpen struct {
	opensIndex   uint16
	opensFlags   uint16
	opensToCount uint16
	opensToIndex []uint16
}

type ModuleProvide struct {
	providesIndex     uint16
	providesWithCount uint16
	providesWithIndex []uint16
}

type ModulePackages struct {
	AttributeType

	numberOfPackages uint16
	packages         []uint16
}

type ModuleMainClass struct {
	AttributeType

	mainClassIndex uint16
}

type NestHost struct {
	AttributeType

	classIndex uint16
}

type NestMembers struct {
	AttributeType

	numberOfClasses uint16
	classes         []uint16
}

type Record struct {
	AttributeType

	componentsCount uint16
	components      []RecordComponentInfo
}

type RecordComponentInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
	attributeInfo   []AttributeInfo
}

type UnknownAttributeType struct {
	AttributeType

	data []uint8
}

func (reader *ClassReader) ReadClassFile() *ClassFile {
	magic := reader.ReadUint32()

	if magic != 0xCAFEBABE {
		panic("magic number error")
	}

	minor := reader.ReadUint16()
	major := reader.ReadUint16()

	constantPoolCount := reader.ReadUint16()
	constantPool := make([]ClassPoolInfo, constantPoolCount)

	for i := 1; i < int(constantPoolCount); i++ {
		constantPool[i] = reader.ReadClassPoolInfo()
	}

	accessFlags := reader.ReadUint16()
	thisClass := reader.ReadUint16()
	superClass := reader.ReadUint16()
	interfacesCount := reader.ReadUint16()
	interfaces := reader.ReadUint16s(uint64(interfacesCount))
	fieldsCount := reader.ReadUint16()
	fields := make([]MemberInfo, fieldsCount)

	for i := 0; i < int(fieldsCount); i++ {
		fields[i] = reader.ReadMemberInfo(constantPool)
	}

	methodsCount := reader.ReadUint16()
	methods := make([]MemberInfo, methodsCount)

	for i := 0; i < int(methodsCount); i++ {
		methods[i] = reader.ReadMemberInfo(constantPool)
	}

	attributesCount := reader.ReadUint16()
	attributes := make([]AttributeInfo, attributesCount)

	for i := 0; i < int(attributesCount); i++ {
		attributes[i] = reader.ReadAttributeInfo(constantPool)
	}

	return &ClassFile{
		magic,
		minor,
		major,
		constantPool,
		accessFlags,
		thisClass,
		superClass,
		interfaces,
		fields,
		methods,
		attributes,
	}
}

func (reader *ClassReader) ReadMemberInfo(cp []ClassPoolInfo) MemberInfo {
	accessFlags := reader.ReadUint16()
	nameIndex := reader.ReadUint16()
	descriptorIndex := reader.ReadUint16()
	attributesCount := reader.ReadUint16()
	attributes := make([]AttributeInfo, attributesCount)

	for i := 0; i < int(attributesCount); i++ {
		attributes[i] = reader.ReadAttributeInfo(cp)
	}

	return MemberInfo{
		accessFlags,
		nameIndex,
		descriptorIndex,
		attributesCount,
		attributes,
	}
}

func (reader *ClassReader) ReadAttributeInfo(cp []ClassPoolInfo) AttributeInfo {
	attributeNameIndex := reader.ReadUint16()
	attributeLength := reader.ReadUint32()
	attributeName := AsString(cp[attributeNameIndex])

	return AttributeInfo{
		attributeNameIndex,
		attributeLength,
		reader.ReadAttributeType(attributeLength, attributeName, cp),
	}
}

func (reader *ClassReader) ReadAttributeType(attributeLength uint32, attributeName string, cp []ClassPoolInfo) AttributeType {
	switch attributeName {
	case "ConstantValue":
		return ConstantValue{
			constantValueIndex: reader.ReadUint16(),
		}
	case "Code":
		maxStack := reader.ReadUint16()
		maxLocals := reader.ReadUint16()
		codeLength := reader.ReadUint32()
		code := reader.ReadUint8s(uint64(codeLength))

		return Code{
			maxStack:       maxStack,
			maxLocals:      maxLocals,
			codeLength:     codeLength,
			code:           code,
			exceptionTable: reader.ReadExceptionTable(),
			attributes:     reader.ReadAttributes(reader.ReadUint16(), cp),
		}
	case "StackMapTable":
		numberOfEntries := reader.ReadUint16()

		return StackMapTable{
			numberOfEntries: numberOfEntries,
			entries:         reader.ReadStackMapTableEntries(numberOfEntries),
		}
	case "Exceptions":
		return Exceptions{
			numberOfExceptions: reader.ReadUint16(),
			exceptionIndices:   reader.ReadUint16s(uint64(attributeLength)),
		}
	case "InnerClasses":
		numberOfClasses := reader.ReadUint16()

		return InnerClasses{
			numberOfClasses: numberOfClasses,
			classes:         reader.ReadInnerClasses(numberOfClasses),
		}
	case "EnclosingMethod":
		return EnclosingMethod{
			classIndex:  reader.ReadUint16(),
			methodIndex: reader.ReadUint16(),
		}
	case "Synthetic":
		return Synthetic{}
	case "SourceFile":
		return SourceFile{
			sourceFileIndex: reader.ReadUint16(),
		}
	case "SourceDebugExtension":
		return SourceDebugExtension{
			debugExtension: reader.ReadUint8s(uint64(attributeLength)),
		}
	case "LineNumberTable":
		return LineNumberTable{
			lineNumberTable: reader.ReadLineNumberTable(),
		}
	case "LocalVariableTable":
		return LocalVariableTable{
			localVariableTable: reader.ReadLocalVariableTable(),
		}
	case "LocalVariableTypeTable":
		return LocalVariableTypeTable{
			localVariableTypeTable: reader.ReadLocalVariableTypeTable(),
		}
	case "Deprecated":
		return Deprecated{}
	case "Signature":
		return Signature{
			signatureIndex: reader.ReadUint16(),
		}
	case "RuntimeVisibleAnnotations":
		return RuntimeVisibleAnnotations{
			annotations: reader.ReadAnnotations(),
		}
	case "RuntimeInvisibleAnnotations":
		return RuntimeInvisibleAnnotations{
			annotations: reader.ReadAnnotations(),
		}
	case "RuntimeVisibleParameterAnnotations":
		return RuntimeVisibleParameterAnnotations{
			parameterAnnotations: reader.ReadParameterAnnotations(),
		}
	case "RuntimeInvisibleParameterAnnotations":
		return RuntimeInvisibleParameterAnnotations{
			parameterAnnotations: reader.ReadParameterAnnotations(),
		}
	case "AnnotationDefault":
		return AnnotationDefault{
			defaultValue: reader.ReadElementValue(),
		}
	case "BootstrapMethods":
		numberOfBootstrapMethods := reader.ReadUint16()

		return BootstrapMethods{
			numberOfBootstrapMethods: numberOfBootstrapMethods,
			bootstrapMethods:         reader.ReadBootstrapMethods(numberOfBootstrapMethods),
		}
	case "MethodParameters":
		numberOfParameters := reader.ReadUint8()

		return MethodParameters{
			numberOfParameters: numberOfParameters,
			parameters:         reader.ReadMethodParameters(numberOfParameters),
		}
	case "Module":
		return Module{
			moduleNameIndex:    reader.ReadUint16(),
			moduleFlags:        reader.ReadUint16(),
			moduleVersionIndex: reader.ReadUint16(),
			moduleInfo:         reader.ReadModuleInfo(),
		}
	case "ModulePackages":
		numberOfPackages := reader.ReadUint16()

		return ModulePackages{
			numberOfPackages: numberOfPackages,
			packages:         reader.ReadUint16s(uint64(numberOfPackages)),
		}
	case "ModuleMainClass":
		return ModuleMainClass{
			mainClassIndex: reader.ReadUint16(),
		}
	case "NestHost":
		return NestHost{
			classIndex: reader.ReadUint16(),
		}
	case "NestMembers":
		return NestMembers{
			numberOfClasses: reader.ReadUint16(),
			classes:         reader.ReadUint16s(uint64(attributeLength)),
		}
	case "Record":
		componentsCount := reader.ReadUint16()

		return Record{
			componentsCount: componentsCount,
			components:      reader.ReadRecordComponentInfos(componentsCount, cp),
		}
	default:
		return UnknownAttributeType{
			data: reader.ReadUint8s(uint64(attributeLength)),
		}
	}
}

func (reader *ClassReader) ReadLocalVariableTypeTable() []LocalVariableType {
	numberOfEntries := reader.ReadUint16()

	localVariableTypes := make([]LocalVariableType, numberOfEntries)

	for i := 0; i < int(numberOfEntries); i++ {
		localVariableTypes[i] = LocalVariableType{
			startPc:        reader.ReadUint16(),
			length:         reader.ReadUint16(),
			nameIndex:      reader.ReadUint16(),
			signatureIndex: reader.ReadUint16(),
			index:          reader.ReadUint16(),
		}
	}

	return localVariableTypes
}

func (reader *ClassReader) ReadRecordComponentInfos(componentsCount uint16, cp []ClassPoolInfo) []RecordComponentInfo {
	components := make([]RecordComponentInfo, componentsCount)

	for i := uint16(0); i < componentsCount; i++ {
		components[i] = RecordComponentInfo{
			nameIndex:       reader.ReadUint16(),
			descriptorIndex: reader.ReadUint16(),
			attributeInfo:   reader.ReadAttributes(reader.ReadUint16(), cp),
		}
	}

	return components
}

func (reader *ClassReader) ReadModuleInfo() ModuleInfo {
	requiresCount := reader.ReadUint16()
	requires := reader.ReadModuleRequires(requiresCount)
	exportsCount := reader.ReadUint16()
	exports := reader.ReadModuleExports(exportsCount)
	opensCount := reader.ReadUint16()
	opens := reader.ReadModuleOpens(opensCount)
	usesCount := reader.ReadUint16()
	usesIndex := reader.ReadUint16s(uint64(usesCount))
	providesCount := reader.ReadUint16()
	provides := reader.ReadModuleProvides(providesCount)

	return ModuleInfo{
		requiresCount: requiresCount,
		requires:      requires,
		exportsCount:  exportsCount,
		exports:       exports,
		opensCount:    opensCount,
		opens:         opens,
		usesCount:     usesCount,
		usesIndex:     usesIndex,
		providesCount: providesCount,
		provides:      provides,
	}
}

func (reader *ClassReader) ReadModuleRequires(count uint16) []ModuleRequire {
	moduleRequires := make([]ModuleRequire, count)

	for i := uint16(0); i < count; i++ {
		moduleRequires[i] = ModuleRequire{
			requiresIndex:        reader.ReadUint16(),
			requiresFlags:        reader.ReadUint16(),
			requiresVersionIndex: reader.ReadUint16(),
		}
	}

	return moduleRequires
}

func (reader *ClassReader) ReadModuleExports(count uint16) []ModuleExport {
	moduleExports := make([]ModuleExport, count)

	for i := uint16(0); i < count; i++ {
		exportsIndex := reader.ReadUint16()
		exportsFlags := reader.ReadUint16()
		exportsToCount := reader.ReadUint16()
		exportsToIndex := reader.ReadUint16s(uint64(exportsToCount))

		moduleExports[i] = ModuleExport{
			exportsIndex:   exportsIndex,
			exportsFlags:   exportsFlags,
			exportsToCount: exportsToCount,
			exportsToIndex: exportsToIndex,
		}
	}

	return moduleExports
}

func (reader *ClassReader) ReadModuleOpens(count uint16) []ModuleOpen {
	moduleOpens := make([]ModuleOpen, count)

	for i := uint16(0); i < count; i++ {
		opensIndex := reader.ReadUint16()
		opensFlags := reader.ReadUint16()
		opensToCount := reader.ReadUint16()
		opensToIndex := reader.ReadUint16s(uint64(opensToCount))

		moduleOpens[i] = ModuleOpen{
			opensIndex:   opensIndex,
			opensFlags:   opensFlags,
			opensToCount: opensToCount,
			opensToIndex: opensToIndex,
		}
	}

	return moduleOpens
}

func (reader *ClassReader) ReadModuleProvides(count uint16) []ModuleProvide {
	moduleProvides := make([]ModuleProvide, count)

	for i := uint16(0); i < count; i++ {
		providesIndex := reader.ReadUint16()
		providesWithCount := reader.ReadUint16()
		providesWithIndex := reader.ReadUint16s(uint64(providesWithCount))

		moduleProvides[i] = ModuleProvide{
			providesIndex:     providesIndex,
			providesWithCount: providesWithCount,
			providesWithIndex: providesWithIndex,
		}
	}

	return moduleProvides
}

func (reader *ClassReader) ReadExceptionTable() []ExceptionTableEntry {
	exceptionTableLength := reader.ReadUint16()
	exceptionTable := make([]ExceptionTableEntry, exceptionTableLength)

	for i := 0; i < int(exceptionTableLength); i++ {
		exceptionTable[i] = ExceptionTableEntry{
			startPc:   reader.ReadUint16(),
			endPc:     reader.ReadUint16(),
			handlerPc: reader.ReadUint16(),
			catchType: reader.ReadUint16(),
		}
	}

	return exceptionTable
}

func (reader *ClassReader) ReadAttributes(attributeCount uint16, cp []ClassPoolInfo) []AttributeInfo {
	attributes := make([]AttributeInfo, attributeCount)

	for i := 0; i < int(attributeCount); i++ {
		attributes[i] = reader.ReadAttributeInfo(cp)
	}

	return attributes
}

func (reader *ClassReader) ReadInnerClasses(numberOfClasses uint16) []InnerClass {
	innerClasses := make([]InnerClass, numberOfClasses)

	for i := 0; i < int(numberOfClasses); i++ {
		innerClasses[i] = InnerClass{
			innerClassInfoIndex:   reader.ReadUint16(),
			outerClassInfoIndex:   reader.ReadUint16(),
			innerNameIndex:        reader.ReadUint16(),
			innerClassAccessFlags: reader.ReadUint16(),
		}
	}

	return innerClasses
}

func (reader *ClassReader) ReadLineNumberTable() []LineNumber {
	lineNumberTableLength := reader.ReadUint16()
	lineNumberTable := make([]LineNumber, lineNumberTableLength)

	for i := 0; i < int(lineNumberTableLength); i++ {
		lineNumberTable[i] = LineNumber{
			startPc:    reader.ReadUint16(),
			lineNumber: reader.ReadUint16(),
		}
	}

	return lineNumberTable
}

func (reader *ClassReader) ReadLocalVariableTable() []LocalVariable {
	localVariableTableLength := reader.ReadUint16()
	localVariableTable := make([]LocalVariable, localVariableTableLength)

	for i := 0; i < int(localVariableTableLength); i++ {
		localVariableTable[i] = LocalVariable{
			startPc:         reader.ReadUint16(),
			length:          reader.ReadUint16(),
			nameIndex:       reader.ReadUint16(),
			descriptorIndex: reader.ReadUint16(),
			index:           reader.ReadUint16(),
		}
	}

	return localVariableTable
}

func (reader *ClassReader) ReadAnnotations() []Annotation {
	numberOfAnnotations := reader.ReadUint16()
	annotations := make([]Annotation, numberOfAnnotations)

	for i := 0; i < int(numberOfAnnotations); i++ {
		annotations[i] = Annotation{
			typeIndex:    reader.ReadUint16(),
			elementPairs: reader.ReadElementPairs(),
		}
	}

	return annotations
}

func (reader *ClassReader) ReadElementPairs() []ElementPair {
	numberOfElementPairs := reader.ReadUint16()
	elementPairs := make([]ElementPair, numberOfElementPairs)

	for i := 0; i < int(numberOfElementPairs); i++ {
		elementPairs[i] = ElementPair{
			elementNameIndex: reader.ReadUint16(),
			value:            reader.ReadElementValue(),
		}
	}

	return elementPairs
}

func (reader *ClassReader) ReadElementValue() ElementValue {
	tag := reader.ReadUint8()

	switch tag {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z', 's':
		return ElementValue{
			tag: tag,
			value: &ConstantValueIndex{
				constantValueIndex: reader.ReadUint16(),
			},
		}
	case 'e':
		return ElementValue{
			tag: tag,
			value: &EnumConstantValue{
				typeNameIndex:     reader.ReadUint16(),
				constantNameIndex: reader.ReadUint16(),
			},
		}
	case 'c':
		return ElementValue{
			tag: tag,
			value: &ClassInfoIndex{
				classInfoIndex: reader.ReadUint16(),
			},
		}
	case '@':
		return ElementValue{
			tag: tag,
			value: &AnnotationValue{
				annotationValue: reader.ReadAnnotations(),
			},
		}
	case '[':
		return ElementValue{
			tag: tag,
			value: &ArrayValue{
				values: reader.ReadElementValues(),
			},
		}
	default:
		panic(fmt.Sprintf("Unexpected tag: %v", tag))
	}
}

func (reader *ClassReader) ReadElementValues() []ElementValue {
	numberOfElementValues := reader.ReadUint16()
	elementValues := make([]ElementValue, numberOfElementValues)

	for i := 0; i < int(numberOfElementValues); i++ {
		elementValues[i] = reader.ReadElementValue()
	}

	return elementValues
}

func (reader *ClassReader) ReadParameterAnnotations() []ParameterAnnotation {
	numberOfParameterAnnotations := reader.ReadUint16()
	parameterAnnotations := make([]ParameterAnnotation, numberOfParameterAnnotations)

	for i := 0; i < int(numberOfParameterAnnotations); i++ {
		parameterAnnotations[i] = ParameterAnnotation{
			parameterIndex: reader.ReadUint8(),
			annotations:    reader.ReadAnnotations(),
		}
	}

	return parameterAnnotations
}

func (reader *ClassReader) ReadBootstrapMethods(numberOfBootstrapMethods uint16) []BootstrapMethod {
	bootstrapMethods := make([]BootstrapMethod, numberOfBootstrapMethods)

	for i := 0; i < int(numberOfBootstrapMethods); i++ {
		bootstrapMethods[i] = BootstrapMethod{
			bootstrapMethodRef: reader.ReadUint16(),
			bootstrapArguments: reader.ReadUint16s(uint64(reader.ReadUint16())),
		}
	}

	return bootstrapMethods
}

func (reader *ClassReader) ReadMethodParameters(numberOfMethodParameters uint8) []MethodParameter {
	methodParameters := make([]MethodParameter, numberOfMethodParameters)

	for i := 0; i < int(numberOfMethodParameters); i++ {
		methodParameters[i] = MethodParameter{
			nameIndex:   reader.ReadUint16(),
			accessFlags: reader.ReadUint16(),
		}
	}

	return methodParameters
}

func (reader *ClassReader) ReadStackMapTableEntries(numberOfEntries uint16) []StackMapFrame {
	stackMapTableEntries := make([]StackMapFrame, numberOfEntries)

	for i := 0; i < int(numberOfEntries); i++ {
		stackMapTableEntries[i] = reader.ReadStackMapFrame()
	}

	return stackMapTableEntries
}

func (reader *ClassReader) ReadStackMapFrame() StackMapFrame {
	frameType := reader.ReadUint8()

	switch {
	case frameType <= 63:
		return SameFrame{
			frameType: frameType,
		}
	case frameType <= 127:
		return SameLocals1StackItemFrame{
			frameType: frameType,
			stack:     reader.ReadVerificationTypeInfo(),
		}
	case frameType == 247:
		return SameLocals1StackItemFrameExtended{
			frameType:   frameType,
			offsetDelta: reader.ReadUint16(),
			stack:       reader.ReadVerificationTypeInfo(),
		}
	case frameType <= 250:
		return ChopFrame{
			frameType:   frameType,
			offsetDelta: reader.ReadUint16(),
		}
	case frameType == 251:
		return SameFrameExtended{
			frameType:   frameType,
			offsetDelta: reader.ReadUint16(),
		}
	case frameType <= 254:
		return AppendFrame{
			frameType:   frameType,
			offsetDelta: reader.ReadUint16(),
			locals:      reader.ReadVerificationTypeInfos(uint16(frameType - 251)),
		}
	case frameType == 255:
		return FullFrame{
			frameType:          frameType,
			offsetDelta:        reader.ReadUint16(),
			numberOfLocals:     reader.ReadUint16(),
			locals:             reader.ReadVerificationTypeInfos(reader.ReadUint16()),
			numberOfStackItems: reader.ReadUint16(),
			stack:              reader.ReadVerificationTypeInfos(reader.ReadUint16()),
		}
	default:
		panic(fmt.Sprintf("Unexpected frame type: %v", frameType))
	}
}

func (reader *ClassReader) ReadVerificationTypeInfos(amount uint16) []VerificationTypeInfo {
	verificationTypes := make([]VerificationTypeInfo, amount)

	for i := 0; i < int(amount); i++ {
		verificationTypes[i] = reader.ReadVerificationTypeInfo()
	}

	return verificationTypes
}

func (reader *ClassReader) ReadVerificationTypeInfo() VerificationTypeInfo {
	tag := reader.ReadUint8()

	switch tag {
	case 0:
		return TopVariable{}
	case 1:
		return IntegerVariable{}
	case 2:
		return FloatVariable{}
	case 3:
		return DoubleVariable{}
	case 4:
		return LongVariable{}
	case 5:
		return NullVariable{}
	case 6:
		return UninitializedThisVariable{}
	case 7:
		return ObjectVariable{
			cpoolIndex: reader.ReadUint16(),
		}
	case 8:
		return UninitializedVariable{
			offset: reader.ReadUint16(),
		}
	default:
		panic(fmt.Sprintf("Unexpected frame type: %v", tag))
	}
}
