package main

import "fmt"

type Runtime struct {
	resources *Resources
}

func CreateRuntime(resources map[string]ClassFile) Runtime {
	return Runtime{
		resources: &Resources{
			classes: resources,
		},
	}
}

func (r *Runtime) AddResource(name string, class ClassFile) {
	r.resources.classes[name] = class
}

func (r *Runtime) AddResources(resources map[string]ClassFile) {
	for name, class := range resources {
		r.resources.classes[name] = class
	}
}

func (r *Runtime) GetResource(name string) ClassFile {
	if resource, ok := r.resources.classes[name]; ok {
		return resource
	}

	panic(fmt.Sprintf("Unknown class %s", name))
}

func (r *Runtime) GetResources() map[string]ClassFile {
	return r.resources.classes
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

func (r *Runtime) makeEnv() *JEnv {
	return &JEnv{
		runtime: r,
		frame:   &Frame{
			/* root frame */
		},
	}
}

func (r *Runtime) Run(class string) {
	env := r.makeEnv()
	classLoader := env.createClassLoader(nil)
	klass := classLoader.findOrLoadClass(env, class)

	klass.FindMethod("main", "([Ljava/lang/String;)V").Invoke(env, klass, nil, []interface{}{})
}

func (c *RuntimeClass) FindMethod(name, descriptor string) *Method {
	for _, method := range c.methods {
		if method.name == name && method.rawDescriptor == descriptor {
			return method
		}
	}

	return nil
}

func (jmo *JMetaObject) init(env *JEnv, cdesc string /* default ()V */, args []interface{}) {
	jmo.class.FindMethod("<init>", cdesc).Invoke(env, jmo.class, jmo, args)
}

func (m *Method) Invoke(env *JEnv, klass RuntimeClass, this *JMetaObject, args []interface{}) bool {
	if m == nil {
		// this method doesn't exist, tf you're tryna do?

		return false // fail
	}

	// TODO call <clinit>

	if this != nil && !this.initialized {
		this.initialized = true

		//klass.FindMethod("<init>", "()V").Invoke(env, klass, this, []interface{}{})
		this.init(env, "()V", []interface{}{})
	}

	env.frame = &Frame{
		stack:       StackOf(m.maxStack),
		root:        env.frame,
		classLoader: klass.classLoader,
	}

	RunCode(m.code, putLocals(m.locals, this, args), env, this, klass)

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
	runtime         *Runtime
	classLoaders    []*ClassLoader
	rootClassLoader *ClassLoader
	frame           *Frame
}

type Resources struct {
	classes map[string]ClassFile
}

type Frame struct {
	stack       *Stack
	root        *Frame
	classLoader *ClassLoader
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

func (s *Stack) PopRef() (o interface{}) {
	s.check(1)

	o = s.values[len(s.values)-1]

	s.values = s.values[:len(s.values)-1]

	return
}

func (s *Stack) PopArray() JArray {
	return s.PopRef().(JArray)
}

func (s *Stack) PopInt() PrimitiveInt {
	return s.PopRef().(JPrimitive).primitive.(PrimitiveInt)
}

func (s *Stack) PopByte() PrimitiveByte {
	return s.PopRef().(JPrimitive).primitive.(PrimitiveByte)
}

func (s *Stack) PopFloat() PrimitiveFloat {
	return s.PopRef().(JPrimitive).primitive.(PrimitiveFloat)
}

func (s *Stack) PopDouble() PrimitiveDouble {
	return s.PopRef().(JPrimitive).primitive.(PrimitiveDouble)
}

func (s *Stack) PopLong() PrimitiveLong {
	return s.PopRef().(JPrimitive).primitive.(PrimitiveLong)
}

func (s *Stack) PushRef(o interface{}) {
	s.ready(1)

	s.values = append(s.values, o)
}

func (s *Stack) PushInt(i PrimitiveInt) {
	s.PushRef(NewPrimitiveInt(i))
}

func (s *Stack) PushByte(b PrimitiveByte) {
	s.PushRef(NewPrimitiveByte(b))
}

func (s *Stack) PushFloat(f PrimitiveFloat) {
	s.PushRef(NewPrimitiveFloat(f))
}

func (s *Stack) PushDouble(d PrimitiveDouble) {
	s.PushRef(NewPrimitiveDouble(d))
}

func (s *Stack) PushLong(l PrimitiveLong) {
	s.PushRef(NewPrimitiveLong(l))
}

func (s *Stack) PushShort(p PrimitiveShort) {
	s.PushRef(NewPrimitiveShort(p))
}

func (s *Stack) PushChar(p PrimitiveChar) {
	s.PushRef(NewPrimitiveChar(p))
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

func NewPrimitiveByte(b PrimitiveByte) interface{} {
	return JPrimitive{
		primitive: b,
	}
}

func NewPrimitiveByteD(b byte) interface{} {
	return JPrimitive{
		primitive: PrimitiveByte{
			value: b,
		},
	}
}

func NewPrimitiveByteRD(b byte) PrimitiveByte {
	return PrimitiveByte{
		value: b,
	}
}

func NewPrimitiveFloat(f PrimitiveFloat) interface{} {
	return JPrimitive{
		primitive: f,
	}
}

func NewPrimitiveFloatRD(f float32) PrimitiveFloat {
	return PrimitiveFloat{
		value: f,
	}
}

func NewPrimitiveDouble(d PrimitiveDouble) interface{} {
	return JPrimitive{
		primitive: d,
	}
}

func NewPrimitiveDoubleRD(d float64) PrimitiveDouble {
	return PrimitiveDouble{
		value: d,
	}
}

func NewPrimitiveLong(l PrimitiveLong) interface{} {
	return JPrimitive{
		primitive: l,
	}
}

func NewPrimitiveLongRD(l int64) PrimitiveLong {
	return PrimitiveLong{
		value: l,
	}
}

func NewPrimitiveShort(s PrimitiveShort) interface{} {
	return JPrimitive{
		primitive: s,
	}
}

func NewPrimitiveShortRD(l int16) PrimitiveShort {
	return PrimitiveShort{
		value: l,
	}
}

func NewPrimitiveChar(c PrimitiveChar) interface{} {
	return JPrimitive{
		primitive: c,
	}
}

func NewPrimitiveCharRD(c int32) PrimitiveChar {
	return PrimitiveChar{
		value: c,
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

	class       RuntimeClass
	initialized bool
	data        map[string]interface{} // fields
}

type JArray struct {
	//JObject

	atype byte
	data  []interface{}
}

func (env *JEnv) createArray(atype byte, length int32) JArray {
	return JArray{
		atype: atype,
		data:  make([]interface{}, length),
	}
}

func (cl *ClassLoader) makeString(env *JEnv, value string) *JMetaObject {
	// create the object
	obj := cl.createObject(env, "java/lang/String")

	// init it with the str value
	// by first creating a char array
	carr := env.createArray(jaChar, int32(len(value)))

	// put the chars into the array
	for i, c := range value {
		carr.SetRefD(int32(i), NewPrimitiveCharRD(c))
	}

	// call the constructor
	obj.init(env, "([C)V", []interface{}{carr})

	// finally, return the string object
	return obj
}

const (
	jaReference = 0
	jaArray     = 1
	jaBoolean   = 2
	jaChar      = 3
	jaFloat     = 4
	jaDouble    = 5
	jaByte      = 6
	jaShort     = 7
	jaInt       = 8
	jaLong      = 9
)

type JPrimitive struct {
	//JObject

	primitive Primitive
}

type Primitive interface{}

type PrimitiveInt struct {
	//Primitive

	value int32
}

func (i PrimitiveInt) ToByte() PrimitiveByte {
	return NewPrimitiveByteRD(byte(i.value))
}

func (i PrimitiveInt) ToChar() PrimitiveChar {
	return NewPrimitiveCharRD(i.value)
}

func (i PrimitiveInt) ToDouble() PrimitiveDouble {
	return NewPrimitiveDoubleRD(float64(i.value))
}

func (i PrimitiveInt) ToFloat() PrimitiveFloat {
	return NewPrimitiveFloatRD(float32(i.value))
}

func (i PrimitiveInt) ToLong() PrimitiveLong {
	return NewPrimitiveLongRD(int64(i.value))
}

func (i PrimitiveInt) ToShort() PrimitiveShort {
	return NewPrimitiveShortRD(int16(i.value))
}

func (i PrimitiveInt) Add(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(i.value + j.value)
}

func (i PrimitiveInt) And(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(i.value & j.value)
}

func (i PrimitiveInt) Div(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(i.value / j.value)
}

func (i PrimitiveInt) Mul(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(i.value * j.value)
}

func (i PrimitiveInt) Neg() PrimitiveInt {
	return NewPrimitiveIntRD(-i.value)
}

func (i PrimitiveInt) Or(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(i.value | j.value)
}

func (i PrimitiveInt) Rem(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(i.value % j.value)
}

func (i PrimitiveInt) Shl(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(i.value << (j.value & 0x1f))
}

func (i PrimitiveInt) Shr(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(i.value >> (j.value & 0x1f))
}

func (i PrimitiveInt) Sub(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(i.value - j.value)
}

func (i PrimitiveInt) Xor(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(i.value ^ j.value)
}

func (i PrimitiveInt) Ushr(j PrimitiveInt) PrimitiveInt {
	return NewPrimitiveIntRD(int32(uint32(i.value) >> (j.value & 0x1f)))
}

type PrimitiveLong struct {
	//Primitive

	value int64
}

func (l PrimitiveLong) ToDouble() PrimitiveDouble {
	return NewPrimitiveDoubleRD(float64(l.value))
}

func (l PrimitiveLong) ToFloat() PrimitiveFloat {
	return NewPrimitiveFloatRD(float32(l.value))
}

func (l PrimitiveLong) ToInt() PrimitiveInt {
	return NewPrimitiveIntRD(int32(l.value))
}

func (l PrimitiveLong) Add(j PrimitiveLong) PrimitiveLong {
	return NewPrimitiveLongRD(l.value + j.value)
}

func (l PrimitiveLong) Div(j PrimitiveLong) PrimitiveLong {
	return NewPrimitiveLongRD(l.value / j.value)
}

func (l PrimitiveLong) Mul(j PrimitiveLong) PrimitiveLong {
	return NewPrimitiveLongRD(l.value * j.value)
}

func (l PrimitiveLong) Neg() PrimitiveLong {
	return NewPrimitiveLongRD(-l.value)
}

func (l PrimitiveLong) Rem(j PrimitiveLong) PrimitiveLong {
	return NewPrimitiveLongRD(l.value % j.value)
}

func (l PrimitiveLong) Or(j PrimitiveLong) PrimitiveLong {
	return NewPrimitiveLongRD(l.value | j.value)
}

func (l PrimitiveLong) Shl(j PrimitiveLong) PrimitiveLong {
	return NewPrimitiveLongRD(l.value << (j.value & 0x3f))
}

func (l PrimitiveLong) Shr(j PrimitiveLong) PrimitiveLong {
	return NewPrimitiveLongRD(l.value >> (j.value & 0x3f))
}

func (l PrimitiveLong) Sub(j PrimitiveLong) PrimitiveLong {
	return NewPrimitiveLongRD(l.value - j.value)
}

func (l PrimitiveLong) Xor(j PrimitiveLong) PrimitiveLong {
	return NewPrimitiveLongRD(l.value ^ j.value)
}

func (l PrimitiveLong) Ushr(j PrimitiveLong) PrimitiveLong {
	return NewPrimitiveLongRD(int64(uint64(l.value) >> (j.value & 0x3f)))
}

type PrimitiveFloat struct {
	//Primitive

	value float32
}

func (f PrimitiveFloat) Add(o PrimitiveFloat) PrimitiveFloat {
	return PrimitiveFloat{
		value: f.value + o.value,
	}
}

func (f PrimitiveFloat) Neg() PrimitiveFloat {
	return PrimitiveFloat{
		value: -f.value,
	}
}

func (f PrimitiveFloat) Sub(o PrimitiveFloat) PrimitiveFloat {
	return f.Add(o.Neg())
}

func (f PrimitiveFloat) Mul(o PrimitiveFloat) PrimitiveFloat {
	return PrimitiveFloat{
		value: f.value * o.value,
	}
}

func (f PrimitiveFloat) Div(o PrimitiveFloat) PrimitiveFloat {
	return PrimitiveFloat{
		value: f.value / o.value,
	}
}

func (f PrimitiveFloat) Rem(o PrimitiveFloat) PrimitiveFloat {
	return PrimitiveFloat{
		value: float32(int32(f.value) % int32(o.value)),
	}
}

func (f PrimitiveFloat) ToDouble() PrimitiveDouble {
	return NewPrimitiveDoubleRD(float64(f.value))
}

func (f PrimitiveFloat) ToInt() PrimitiveInt {
	return NewPrimitiveIntRD(int32(f.value))
}

func (f PrimitiveFloat) ToLong() PrimitiveLong {
	return NewPrimitiveLongRD(int64(f.value))
}

type PrimitiveDouble struct {
	//Primitive

	value float64
}

func (d PrimitiveDouble) Add(o PrimitiveDouble) PrimitiveDouble {
	return PrimitiveDouble{
		value: d.value + o.value,
	}
}

func (d PrimitiveDouble) Neg() PrimitiveDouble {
	return PrimitiveDouble{
		value: -d.value,
	}
}

func (d PrimitiveDouble) Sub(o PrimitiveDouble) PrimitiveDouble {
	return d.Add(o.Neg())
}

func (d PrimitiveDouble) Mul(o PrimitiveDouble) PrimitiveDouble {
	return PrimitiveDouble{
		value: d.value * o.value,
	}
}

func (d PrimitiveDouble) Div(o PrimitiveDouble) PrimitiveDouble {
	return PrimitiveDouble{
		value: d.value / o.value,
	}
}

func (d PrimitiveDouble) Rem(o PrimitiveDouble) PrimitiveDouble {
	return PrimitiveDouble{
		value: float64(int64(d.value) % int64(o.value)), // TODO might not work
	}
}

func (d PrimitiveDouble) ToFloat() PrimitiveFloat {
	return NewPrimitiveFloatRD(float32(d.value))
}

func (d PrimitiveDouble) ToInt() PrimitiveInt {
	return NewPrimitiveIntRD(int32(d.value))
}

func (d PrimitiveDouble) ToLong() PrimitiveLong {
	return NewPrimitiveLongRD(int64(d.value))
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

	value int32
}

type PrimitiveVoid struct {
	Primitive
}

type ClassLoader struct {
	parent  *ClassLoader
	classes map[string]*RuntimeClass
}

type RuntimeClass struct {
	classLoader  *ClassLoader
	methods      map[string]*Method
	fields       map[string]*Field
	name         string
	flags        []ClassAccessFlag
	constantPool *RuntimeConstantPool
	staticFields map[string]interface{}
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
		classLoader:  cl,
		methods:      mapMethods(class.methods, class.constantPool),
		fields:       mapFields(class, class.constantPool),
		name:         AsString(class.constantPool[class.constantPool[class.thisClass].(*Class).nameIndex]),
		flags:        mapCAccessFlags(class.accessFlags),
		constantPool: Transform(class), // added in post lmao
	}

	cls.FindMethod("<clinit>", "()V").Invoke(env, cls, nil, []interface{}{})

	return
}

func (cl *ClassLoader) findOrLoadClass(env *JEnv, name string) RuntimeClass {
	if cls, ok := cl.classes[name]; ok {
		return *cls
	}

	return cl.createRuntimeClass(env, env.runtime.GetResource(name))
}

func (cl *ClassLoader) createObject(env *JEnv, class string) *JMetaObject {
	return &JMetaObject{
		class: cl.findOrLoadClass(env, class),
		data:  map[string]interface{}{},
	}
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
