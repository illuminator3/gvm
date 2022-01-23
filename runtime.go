package main

import "fmt"

type Runtime struct{}

func CreateRuntime() Runtime {
	return Runtime{}
}

func (env *JEnv) createClassLoader(parent *ClassLoader) *ClassLoader {
	cl := ClassLoader{
		parent: parent,
	}

	env.classLoaders = append(env.classLoaders, &cl)

	if parent != nil {
		env.rootClassLoader = &cl
	}

	return &cl
}

func makeEnv() *JEnv {
	return &JEnv{}
}

func (r *Runtime) Run(class ClassFile) {
	env := makeEnv()
	classLoader := env.createClassLoader(nil)
	klass := classLoader.createRuntimeClass(env, class)

	klass.FindMethod("main", "([Ljava/lang/String;)V").Invoke(env, klass, klass.metaObject, []interface{}{})
}

func (c *RuntimeClass) FindMethod(name, descriptor string) *Method {
	for _, method := range c.methods {
		if method.name == name && method.rawDescriptor == descriptor {
			return method
		}
	}

	return nil
}

func (m *Method) Invoke(env *JEnv, klass RuntimeClass, this *JMetaObject, args []interface{}) bool {
	if m == nil {
		// this method doesn't exist, tf you're tryna do?

		return false // fail
	}

	if !this.initialized {
		this.initialized = true

		klass.FindMethod("<init>", "()V").Invoke(env, klass, this, []interface{}{})
	}

	RunCode(m.code, putLocals(m.locals, this, args), env, this, m.maxStack)

	return true // success!
}

func putLocals(locals []RuntimeLocalVariable, this *JMetaObject, args []interface{}) []RuntimeLocalVariable {
	locals[0].obj = this

	for i := 1; i < len(args); i++ {
		locals[i].obj = args[i-1]
	}

	return locals
}

type RuntimeLocalVariable struct {
	name string
	slot int
	obj  interface{}
}

type JEnv struct {
	stack           *Stack
	classLoaders    []*ClassLoader
	rootClassLoader *ClassLoader
}

type Stack struct {
	values    []interface{}
	resizable bool
	maxSize   int
}

func CreateStack() *Stack {
	return ResizableStack()
}

func StackOf(size int) *Stack {
	return &Stack{
		values:    make([]interface{}, size),
		resizable: false,
		maxSize:   size,
	}
}

func ResizableStack() *Stack {
	return &Stack{
		values:    make([]interface{}, 0),
		resizable: true,
		maxSize:   0,
	}
}

func (s *Stack) ready(n int) {
	if s.resizable {
		if n > s.maxSize {
			panic("Stack overflow")
		}
	} else {
		if n > len(s.values)+1 { // no idea if this actually works
			panic("Stack overflow")
		}
	}
}

func (s *Stack) check(n int) {
	if n >= len(s.values) {
		panic("Stack underflow")
	}
}

func (s *Stack) PopInt() (i PrimitiveInt) {
	s.check(1)

	i = s.PopRef().(JPrimitive).primitive.(PrimitiveInt)

	s.values = s.values[:len(s.values)-1]

	return
}

func (s *Stack) PopRef() (o interface{}) {
	s.check(1)

	o = s.values[len(s.values)-1]

	s.values = s.values[:len(s.values)-1]

	return
}

func (s *Stack) PushInt(i PrimitiveInt) {
	s.ready(1)

	pri := NewPrimitiveInt(i)

	s.values = append(s.values, &pri)
}

func (s *Stack) PushRef(o interface{}) {
	s.ready(1)

	s.values = append(s.values, o)
}

func NewPrimitiveInt(i PrimitiveInt) interface{} {
	return JPrimitive{
		primitive: i,
	}
}

func NewPrimitiveIntD(i int32) interface{} {
	return JPrimitive{
		primitive: PrimitiveInt{
			value: i,
		},
	}
}

func NewPrimitiveIntRD(i int32) PrimitiveInt {
	return PrimitiveInt{
		value: i,
	}
}

func (a JArray) SetRef(i PrimitiveInt, o interface{}) {
	a.SetRefD(i.value, o)
}

func (a JArray) SetRefD(i int32, o interface{}) {
	a.data[i] = o
}

func (a JArray) GetRef(i PrimitiveInt) interface{} {
	return a.GetRefD(i.value)
}

func (a JArray) GetRefD(i int32) interface{} {
	return a.data[i]
}

func (a JArray) ArrayLength() PrimitiveInt {
	return NewPrimitiveIntRD(int32(len(a.data)))
}

func (a JArray) ArrayLengthD() int32 {
	return int32(len(a.data))
}

//type JObject interface{}
//type JObject A

type JMetaObject struct {
	//JObject

	class        *RuntimeClass
	constantPool *RuntimeConstantPool
	initialized  bool
}

type JRObject struct {
	//JObject

	class *RuntimeClass
	data  map[string]interface{}
}

type JArray struct {
	//JObject

	data []interface{}
}

type JPrimitive struct {
	//JObject

	primitive Primitive
}

type Primitive interface{}

type PrimitiveInt struct {
	Primitive

	value int32
}

type PrimitiveLong struct {
	Primitive

	value int64
}

type PrimitiveFloat struct {
	Primitive

	value float32
}

type PrimitiveDouble struct {
	Primitive

	value float64
}

type PrimitiveBoolean struct {
	Primitive

	value bool
}

type PrimitiveByte struct {
	Primitive

	value byte
}

type PrimitiveShort struct {
	Primitive

	value int16
}

type PrimitiveChar struct {
	Primitive

	value uint16
}

type PrimitiveVoid struct {
	Primitive
}

type ClassLoader struct {
	parent  *ClassLoader
	classes map[string]*Class
}

type RuntimeClass struct {
	classLoader *ClassLoader
	methods     map[string]*Method
	fields      map[string]*Field
	name        string
	flags       []ClassAccessFlag
	metaObject  *JMetaObject
}

type Method struct {
	name          string
	rawDescriptor string
	accessFlags   []MethodAccessFlag
	code          []Instruction
	maxStack      int
	maxLocals     int
	descriptor    MethodDescriptor
	locals        []RuntimeLocalVariable
}

type Field struct {
	name        string
	descriptor  string
	tipe        Type
	accessFlags []FieldAccessFlag
}

type Type interface{}

type TypeByte struct {
	Type
}

type TypeChar struct {
	Type
}

type TypeDouble struct {
	Type
}

type TypeFloat struct {
	Type
}

type TypeInt struct {
	Type
}

type TypeLong struct {
	Type
}

type TypeShort struct {
	Type
}

type TypeBoolean struct {
	Type
}

type TypeVoid struct {
	Type
}

type TypeArray struct {
	Type

	componentType Type
}

type TypeClass struct {
	Type

	name string
}

type AccessFlag interface{}

type ClassAccessFlag struct {
	AccessFlag

	value uint16
}

type MethodAccessFlag struct {
	AccessFlag

	value uint16
}

type FieldAccessFlag struct {
	AccessFlag

	value uint16
}

// Class Access Flags
var (
	CAccPublic     = cCAF(0x0001)
	CAccFinal      = cCAF(0x0010)
	CAccSuper      = cCAF(0x0020)
	CAccInterface  = cCAF(0x0200)
	CAccAbstract   = cCAF(0x0400)
	CAccSynthetic  = cCAF(0x1000)
	CAccAnnotation = cCAF(0x2000)
	CAccEnum       = cCAF(0x4000)
	CAccModule     = cCAF(0x8000)
)

// Method Access Flags
var (
	MAccPublic       = cMAF(0x0001)
	MAccPrivate      = cMAF(0x0002)
	MAccProtected    = cMAF(0x0004)
	MAccStatic       = cMAF(0x0008)
	MAccFinal        = cMAF(0x0010)
	MAccSynchronized = cMAF(0x0020)
	MAccBridge       = cMAF(0x0040)
	MAccVarargs      = cMAF(0x0080)
	MAccNative       = cMAF(0x0100)
	MAccAbstract     = cMAF(0x0400)
	MAccStrict       = cMAF(0x0800)
	MAccSynthetic    = cMAF(0x1000)
)

// Field Access Flags
var (
	FAccPublic    = cFAF(0x0001)
	FAccPrivate   = cFAF(0x0002)
	FAccProtected = cFAF(0x0004)
	FAccStatic    = cFAF(0x0008)
	FAccFinal     = cFAF(0x0010)
	FAccVolatile  = cFAF(0x0040)
	FAccTransient = cFAF(0x0080)
	FAccSynthetic = cFAF(0x1000)
	FAccEnum      = cFAF(0x4000)
)

func cCAF(value uint16) ClassAccessFlag {
	return ClassAccessFlag{
		value: value,
	}
}

func cMAF(value uint16) MethodAccessFlag {
	return MethodAccessFlag{
		value: value,
	}
}

func cFAF(value uint16) FieldAccessFlag {
	return FieldAccessFlag{
		value: value,
	}
}

func (cl *ClassLoader) createRuntimeClass(env *JEnv, class ClassFile) (cls RuntimeClass) {
	cls = RuntimeClass{
		classLoader: cl,
		methods:     mapMethods(class.methods, class.constantPool),
		fields:      mapFields(class, class.constantPool),
		name:        AsString(class.constantPool[class.constantPool[class.thisClass].(*Class).nameIndex]),
		flags:       mapCAccessFlags(class.accessFlags),
		metaObject:  nil,
	}

	cls.metaObject = &JMetaObject{
		class:        &cls,
		constantPool: Transform(class),
	}

	cls.FindMethod("<clinit>", "()V").Invoke(env, cls, cls.metaObject, []interface{}{})

	return
}

func mapMethods(members []MemberInfo, cp []ClassPoolInfo) map[string]*Method {
	methods := make(map[string]*Method, len(members))

	for _, member := range members {
		name := AsString(cp[member.nameIndex])

		mai := findMethodAttributeInfo(member, cp)

		methods[name] = &Method{
			name:          name,
			rawDescriptor: AsString(cp[member.descriptorIndex]),
			accessFlags:   mapMAccessFlags(member.accessFlags),
			code:          makeCode(mai),
			maxStack:      int(mai.maxStack),
			maxLocals:     int(mai.maxLocals),
			descriptor:    MakeMethodDescriptor(CreateStringConsumer(AsString(cp[member.descriptorIndex]))),
			locals:        mapLocalVariables(int(mai.maxLocals), findLocals(mai.codeAttribute.attributes, cp), cp),
		}
	}

	return methods
}

func findLocals(attributes []AttributeInfo, cp []ClassPoolInfo) LocalVariableTable {
	name := func(info AttributeInfo) string {
		return AsString(cp[info.attributeNameIndex])
	}

	for _, attribute := range attributes {
		if name(attribute) == "LocalVariableTable" {
			return attribute.attrType.(LocalVariableTable)
		}
	}

	return LocalVariableTable{}
}

func mapLocalVariables(maxLocals int, lvt LocalVariableTable, cp []ClassPoolInfo) (rlv []RuntimeLocalVariable) {
	rlv = make([]RuntimeLocalVariable, maxLocals)

	for i, lvt := range lvt.localVariableTable {
		rlv[i] = RuntimeLocalVariable{
			name: AsString(cp[lvt.nameIndex]),
			slot: int(lvt.index),
			obj:  nil,
		}
	}

	return
}

func makeCode(mai MethodAttributeInfo) []Instruction {
	if mai.code == nil {
		return nil
	}

	return ParseInstructions(mai.code)
}

type MethodAttributeInfo struct {
	maxStack      uint16
	maxLocals     uint16
	code          []uint8
	codeAttribute Code
}

func findMethodAttributeInfo(member MemberInfo, cp []ClassPoolInfo) MethodAttributeInfo {
	name := func(info AttributeInfo) string {
		return AsString(cp[info.attributeNameIndex])
	}

	mai := MethodAttributeInfo{}

	for _, attr := range member.attributes {
		if name(attr) == "Code" {
			mai.maxStack = attr.attrType.(Code).maxStack
			mai.maxLocals = attr.attrType.(Code).maxLocals
			mai.code = attr.attrType.(Code).code
			mai.codeAttribute = attr.attrType.(Code)

			break
		}
	}

	return mai
}

type MethodDescriptor struct {
	arguments  []Type
	returnType Type
}

func MakeMethodDescriptor(consumer *StringConsumer) MethodDescriptor {
	consumer.Next() // (

	args := MakeDescriptor(CreateStringConsumer(consumer.TakeWhile(func(c string) bool {
		return c != ")"
	})))

	consumer.Next() // )

	return MethodDescriptor{
		arguments:  args,
		returnType: ReadType(consumer),
	}
}

func MakeDescriptor(consumer *StringConsumer) (types []Type) {
	types = make([]Type, 0)

	for consumer.HasNext() {
		types = append(types, ReadType(consumer))
	}

	return
}

func ReadType(consumer *StringConsumer) Type {
	next := consumer.Next()

	switch next {
	case "B":
		return TypeByte{}
	case "C":
		return TypeChar{}
	case "D":
		return TypeDouble{}
	case "F":
		return TypeFloat{}
	case "I":
		return TypeInt{}
	case "J":
		return TypeLong{}
	case "S":
		return TypeShort{}
	case "Z":
		return TypeBoolean{}
	case "V":
		return TypeVoid{}
	case "[":
		return TypeArray{
			componentType: ReadType(consumer),
		}
	case "L":
		defer consumer.Next() // ;

		return TypeClass{
			name: AsString(consumer.TakeWhile(func(c string) bool {
				return c != ";"
			})),
		}
	default:
		panic(fmt.Sprintf("Unknown type: %s", next))
	}
}

func mapFields(class ClassFile, cp []ClassPoolInfo) map[string]*Field {
	fields := make(map[string]*Field, len(class.fields))

	for _, field := range class.fields {
		name := AsString(cp[field.nameIndex])

		fields[name] = &Field{
			name:        name,
			descriptor:  AsString(cp[field.descriptorIndex]),
			accessFlags: mapFAccessFlags(field.accessFlags),
		}
	}

	return fields
}

func mapCAccessFlags(flags uint16) []ClassAccessFlag {
	var accessFlags []ClassAccessFlag

	for _, flag := range []ClassAccessFlag{
		CAccPublic,
		CAccFinal,
		CAccSuper,
		CAccInterface,
		CAccAbstract,
		CAccSynthetic,
		CAccAnnotation,
		CAccEnum,
		CAccModule,
	} {
		if flags&flag.value != 0 {
			accessFlags = append(accessFlags, flag)
		}
	}

	return accessFlags
}

func mapMAccessFlags(flags uint16) []MethodAccessFlag {
	var accessFlags []MethodAccessFlag

	for _, flag := range []MethodAccessFlag{
		MAccPublic,
		MAccPrivate,
		MAccProtected,
		MAccStatic,
		MAccFinal,
		MAccSynchronized,
		MAccBridge,
		MAccVarargs,
		MAccNative,
		MAccAbstract,
		MAccStrict,
		MAccSynthetic,
	} {
		if flags&flag.value != 0 {
			accessFlags = append(accessFlags, flag)
		}
	}

	return accessFlags
}

func mapFAccessFlags(flags uint16) []FieldAccessFlag {
	var accessFlags []FieldAccessFlag

	for _, flag := range []FieldAccessFlag{
		FAccPublic,
		FAccPrivate,
		FAccProtected,
		FAccStatic,
		FAccFinal,
		FAccVolatile,
		FAccTransient,
		FAccSynthetic,
		FAccEnum,
	} {
		if flags&flag.value != 0 {
			accessFlags = append(accessFlags, flag)
		}
	}

	return accessFlags
}
