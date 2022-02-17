package main

import (
	"fmt"
	"math"
)

func RunCode(instr []Instruction, lv []RuntimeLocalVariable, env *JEnv, this *JMetaObject, klass RuntimeClass) {
	for _, i := range instr {
		fmt.Printf("Stack before instruction %v\n", env.frame.stack)

		i.Execute(env, this, *klass.constantPool, lv, checkStack(env.frame), env.frame)

		fmt.Printf("Stack after instruction %v\n", env.frame.stack)
	}
}

func checkStack(frame *Frame) *Stack {
	if frame != nil {
		return frame.stack
	}

	panic("Current frame is nil")
}

func aload(env *JEnv) {
	stack := env.frame.stack
	index := stack.PopInt()
	arrayref := stack.PopArray()

	if index.value < 0 {
		panic("Array index is negative")
	}

	if index.value > arrayref.ArrayLengthD() {
		panic("Array index is out of bounds")
	}

	stack.PushRef(arrayref.GetRef(index))
}

func astore(env *JEnv) {
	stack := env.frame.stack
	ref := stack.PopRef()
	index := stack.PopInt()
	arrayref := stack.PopArray()

	if index.value < 0 {
		panic("Array index is negative")
	}

	if index.value > arrayref.ArrayLengthD() {
		panic("Array index is out of bounds")
	}

	arrayref.SetRef(index, ref)
}

func (aaload Instructionaaload) Execute(env *JEnv, _ *JMetaObject, _ RuntimeConstantPool, _ []RuntimeLocalVariable, _ *Stack, frame *Frame) {
	aload(env)
}

func (aastore Instructionaastore) Execute(env *JEnv, _ *JMetaObject, _ RuntimeConstantPool, _ []RuntimeLocalVariable, _ *Stack, frame *Frame) {
	astore(env)
}

func (aconstNull InstructionaconstNull) Execute(_ *JEnv, _ *JMetaObject, _ RuntimeConstantPool, _ []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(nil)
}

func (aload Instructionaload) Execute(_ *JEnv, _ *JMetaObject, _ RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[aload.Index].obj)
}

func (aload0 Instructionaload0) Execute(_ *JEnv, _ *JMetaObject, _ RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[0].obj)
}

func (aload1 Instructionaload1) Execute(_ *JEnv, _ *JMetaObject, _ RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[1].obj)
}

func (aload2 Instructionaload2) Execute(_ *JEnv, _ *JMetaObject, _ RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[2].obj)
}

func (aload3 Instructionaload3) Execute(_ *JEnv, _ *JMetaObject, _ RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[3].obj)
}

func (anewarray Instructionanewarray) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	array := env.createArray(jaReference, stack.PopInt().value)

	stack.PushRef(array)

	// TODO, sorta class-type of the array?
}

func _return(env *JEnv, control interface{}) {
	fmt.Println(env.frame)

	frame := env.frame
	root := frame.root

	if root != nil {
		env.frame = root

		if control != nil {
			root.stack.PushRef(frame.stack.PopRef()) // doubles/ints/arrays can be treated as refs due to the implementation
		}

		return
	}

	panic("JVM exit (root frame is nil)")
}

func (areturn Instructionareturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_return(env, -1)
}

func (arraylength Instructionarraylength) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopArray().ArrayLength())
}

func (astore Instructionastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[astore.Index].obj = stack.PopRef()
}

func (astore0 Instructionastore0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[0].obj = stack.PopRef()
}

func (astore1 Instructionastore1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[1].obj = stack.PopRef()
}

func (astore2 Instructionastore2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[2].obj = stack.PopRef()
}

func (astore3 Instructionastore3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[3].obj = stack.PopRef()
}

func (athrow Instructionathrow) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (baload Instructionbaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	aload(env)
}

func (bastore Instructionbastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	astore(env)
}

func (bipush Instructionbipush) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushByte(NewPrimitiveByteRD(bipush.Value))
}

func (breakpoint Instructionbreakpoint) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	panic("No implementation for breakpoint instruction")
}

func (caload Instructioncaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	aload(env)
}

func (castore Instructioncastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	astore(env)
}

func (checkcast Instructioncheckcast) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	obj := stack.PopRef()

	// TODO arrays

	if rcp.infos[checkcast.Index].(RClass).name != obj.(JMetaObject).class.name {
		// TODO throw exception
	}

	stack.PushRef(obj) // repush obj
}

func (d2f Instructiond2f) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(stack.PopDouble().ToFloat())
}

func (d2i Instructiond2i) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopDouble().ToInt())
}

func (d2l Instructiond2l) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopDouble().ToLong())
}

func (dadd Instructiondadd) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(stack.PopDouble().Add(stack.PopDouble()))
}

func (daload Instructiondaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	aload(env)
}

func (dastore Instructiondastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	astore(env)
}

func (dcmpg Instructiondcmpg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	v2 := stack.PopDouble().value
	v1 := stack.PopDouble().value

	if math.IsNaN(v1) || math.IsNaN(v2) || v1 > v2 {
		stack.PushInt(NewPrimitiveIntRD(1))
	} else if v1 == v2 {
		stack.PushInt(NewPrimitiveIntRD(0))
	} else {
		stack.PushInt(NewPrimitiveIntRD(-1))
	}
}

func (dcmpl Instructiondcmpl) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	v2 := stack.PopDouble().value
	v1 := stack.PopDouble().value

	if math.IsNaN(v1) || math.IsNaN(v2) || v1 < v2 {
		stack.PushInt(NewPrimitiveIntRD(-1))
	} else if v1 == v2 {
		stack.PushInt(NewPrimitiveIntRD(0))
	} else {
		stack.PushInt(NewPrimitiveIntRD(1))
	}
}

func (dconst0 Instructiondconst0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(NewPrimitiveDoubleRD(0))
}

func (dconst1 Instructiondconst1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(NewPrimitiveDoubleRD(1))
}

func (ddiv Instructionddiv) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(stack.PopDouble().Div(stack.PopDouble()))
}

func _dload(lv []RuntimeLocalVariable, index int, stack *Stack) {
	stack.PushDouble(lv[index].obj.(JPrimitive).primitive.(PrimitiveDouble))
}

func (dload Instructiondload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_dload(lv, int(dload.Index), stack)
}

func (dload0 Instructiondload0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_dload(lv, 0, stack)
}

func (dload1 Instructiondload1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_dload(lv, 1, stack)
}

func (dload2 Instructiondload2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_dload(lv, 2, stack)
}

func (dload3 Instructiondload3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_dload(lv, 3, stack)
}

func (dmul Instructiondmul) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(stack.PopDouble().Mul(stack.PopDouble()))
}

func (dneg Instructiondneg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(stack.PopDouble().Neg())
}

func (drem Instructiondrem) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(stack.PopDouble().Rem(stack.PopDouble()))
}

func (dreturn Instructiondreturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_return(env, -1)
}

func (dstore Instructiondstore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[dstore.Index].obj = stack.PopDouble() // no generic helper function to assure type
}

func (dstore0 Instructiondstore0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[0].obj = stack.PopDouble()
}

func (dstore1 Instructiondstore1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[1].obj = stack.PopDouble()
}

func (dstore2 Instructiondstore2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[2].obj = stack.PopDouble()
}

func (dstore3 Instructiondstore3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[3].obj = stack.PopDouble()
}

func (dsub Instructiondsub) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(stack.PopDouble().Sub(stack.PopDouble()))
}

func (dup Instructiondup) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	el := stack.PopRef() // generic pop once again

	stack.PushRef(el)
	stack.PushRef(el)
}

func (dupX1 InstructiondupX1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	el1 := stack.PopRef()
	el2 := stack.PopRef()

	stack.PushRef(el1)
	stack.PushRef(el2)
	stack.PushRef(el1)
}

func (dupX2 InstructiondupX2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (dup2 Instructiondup2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (dup2X1 Instructiondup2X1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (dup2X2 Instructiondup2X2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (f2d Instructionf2d) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(stack.PopFloat().ToDouble())
}

func (f2i Instructionf2i) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopFloat().ToInt())
}

func (f2l Instructionf2l) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopFloat().ToLong())
}

func (fadd Instructionfadd) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(stack.PopFloat().Add(stack.PopFloat()))
}

func (faload Instructionfaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	aload(env)
}

func (fastore Instructionfastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	astore(env)
}

func (fcmpg Instructionfcmpg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fcmpl Instructionfcmpl) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fconst0 Instructionfconst0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(NewPrimitiveFloatRD(0))
}

func (fconst1 Instructionfconst1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(NewPrimitiveFloatRD(1))
}

func (fconst2 Instructionfconst2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(NewPrimitiveFloatRD(2))
}

func (fdiv Instructionfdiv) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(stack.PopFloat().Div(stack.PopFloat()))
}

func (fload Instructionfload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fload0 Instructionfload0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fload1 Instructionfload1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fload2 Instructionfload2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fload3 Instructionfload3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fmul Instructionfmul) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(stack.PopFloat().Mul(stack.PopFloat()))
}

func (fneg Instructionfneg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(stack.PopFloat().Neg())
}

func (frem Instructionfrem) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(stack.PopFloat().Rem(stack.PopFloat()))
}

func (freturn Instructionfreturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_return(env, -1)
}

func (fstore Instructionfstore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fstore0 Instructionfstore0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fstore1 Instructionfstore1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fstore2 Instructionfstore2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fstore3 Instructionfstore3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (fsub Instructionfsub) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (getfield Instructiongetfield) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (getstatic Instructiongetstatic) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (goto_ Instructiongoto) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (gotoW InstructiongotoW) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (i2b Instructioni2b) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushByte(stack.PopInt().ToByte())
}

func (i2c Instructioni2c) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushChar(stack.PopInt().ToChar())
}

func (i2d Instructioni2d) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(stack.PopInt().ToDouble())
}

func (i2f Instructioni2f) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(stack.PopInt().ToFloat())
}

func (i2l Instructioni2l) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopInt().ToLong())
}

func (i2s Instructioni2s) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushShort(stack.PopInt().ToShort())
}

func (iadd Instructioniadd) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Add(stack.PopInt()))
}

func (iaload Instructioniaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	aload(env)
}

func (iand Instructioniand) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().And(stack.PopInt()))
}

func (iastore Instructioniastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	astore(env)
}

func (iconstM1 InstructioniconstM1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(NewPrimitiveIntRD(-1))
}

func (iconst0 Instructioniconst0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(NewPrimitiveIntRD(0))
}

func (iconst1 Instructioniconst1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(NewPrimitiveIntRD(1))
}

func (iconst2 Instructioniconst2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(NewPrimitiveIntRD(2))
}

func (iconst3 Instructioniconst3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(NewPrimitiveIntRD(3))
}

func (iconst4 Instructioniconst4) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(NewPrimitiveIntRD(4))
}

func (iconst5 Instructioniconst5) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(NewPrimitiveIntRD(5))
}

func (idiv Instructionidiv) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Div(stack.PopInt()))
}

func (ifAcmpeq InstructionifAcmpeq) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifAcmpne InstructionifAcmpne) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifIcmpeq InstructionifIcmpeq) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifIcmpge InstructionifIcmpge) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifIcmpgt InstructionifIcmpgt) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifIcmple InstructionifIcmple) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifIcmplt InstructionifIcmplt) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifIcmpne InstructionifIcmpne) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifeq Instructionifeq) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifge Instructionifge) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifgt Instructionifgt) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifle Instructionifle) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (iflt Instructioniflt) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifne Instructionifne) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifnonnull Instructionifnonnull) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ifnull Instructionifnull) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (iinc Instructioniinc) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[iinc.Index].obj = lv[iinc.Index].obj.(JPrimitive).primitive.(PrimitiveInt).Add(NewPrimitiveIntRD(int32(iinc.Const))) // what am I even doing with my life at this point
}

func (iload Instructioniload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[iload.Index].obj)
}

func (iload0 Instructioniload0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[0].obj)
}

func (iload1 Instructioniload1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[1].obj)
}

func (iload2 Instructioniload2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[2].obj)
}

func (iload3 Instructioniload3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[3].obj)
}

func (impdep1 Instructionimpdep1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (impdep2 Instructionimpdep2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (imul Instructionimul) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Mul(stack.PopInt()))
}

func (ineg Instructionineg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Neg())
}

func (instanceof Instructioninstanceof) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (invokedynamic Instructioninvokedynamic) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (invokeinterface Instructioninvokeinterface) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (invokespecial Instructioninvokespecial) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (invokestatic Instructioninvokestatic) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (invokevirtual Instructioninvokevirtual) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ior Instructionior) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Or(stack.PopInt()))
}

func (irem Instructionirem) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Rem(stack.PopInt()))
}

func (ireturn Instructionireturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_return(env, -1)
}

func (ishl Instructionishl) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Shl(stack.PopInt()))
}

func (ishr Instructionishr) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Shr(stack.PopInt()))
}

func (istore Instructionistore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[istore.Index].obj = stack.PopInt()
}

func (istore0 Instructionistore0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[0].obj = stack.PopInt()
}

func (istore1 Instructionistore1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[1].obj = stack.PopInt()
}

func (istore2 Instructionistore2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[2].obj = stack.PopInt()
}

func (istore3 Instructionistore3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[3].obj = stack.PopInt()
}

func (isub Instructionisub) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Sub(stack.PopInt()))
}

func (iushr Instructioniushr) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Ushr(stack.PopInt()))
}

func (ixor Instructionixor) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopInt().Xor(stack.PopInt()))
}

func (jsr Instructionjsr) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (jsrW InstructionjsrW) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (l2d Instructionl2d) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushDouble(stack.PopLong().ToDouble())
}

func (l2f Instructionl2f) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushFloat(stack.PopLong().ToFloat())
}

func (l2i Instructionl2i) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushInt(stack.PopLong().ToInt())
}

func (ladd Instructionladd) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Add(stack.PopLong()))
}

func (laload Instructionlaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	aload(env)
}

func (land Instructionland) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (lastore Instructionlastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	astore(env)
}

func (lcmp Instructionlcmp) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (lconst0 Instructionlconst0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(NewPrimitiveLongRD(0))
}

func (lconst1 Instructionlconst1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(NewPrimitiveLongRD(1))
}

func (ldc Instructionldc) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	t := rcp.infos[ldc.Index]
	id := t.id()

	if id == RIntegerID {
		stack.PushInt(NewPrimitiveIntRD(t.(*RInteger).value))
	} else if id == RFloatID {
		stack.PushFloat(NewPrimitiveFloatRD(t.(*RFloat).value))
	} else if id == RStringID {
		stack.PushRef(frame.classLoader.makeString(env, t.(*RString).value))
	} else if id == RClassID {
		// TODO
	} else if id == RMethodTypeID {
		// TODO
	} else if id == RMethodHandleID {
		// TODO
	} else if id == RDynamicID {
		// TODO
	} else {
		panic(fmt.Sprintf("Cannot evaluate ldc %d", id))
	}
}

func (ldcW InstructionldcW) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ldc2W Instructionldc2W) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ldiv Instructionldiv) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Div(stack.PopLong()))
}

func (lload Instructionlload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[lload.Index].obj)
}

func (lload0 Instructionlload0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[0].obj)
}

func (lload1 Instructionlload1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[1].obj)
}

func (lload2 Instructionlload2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[2].obj)
}

func (lload3 Instructionlload3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushRef(lv[3].obj)
}

func (lmul Instructionlmul) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Mul(stack.PopLong()))
}

func (lneg Instructionlneg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Neg())
}

func (lookupswitch Instructionlookupswitch) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (lor Instructionlor) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Or(stack.PopLong()))
}

func (lrem Instructionlrem) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Rem(stack.PopLong()))
}

func (lreturn Instructionlreturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_return(env, -1)
}

func (lshl Instructionlshl) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Shl(stack.PopLong()))
}

func (lshr Instructionlshr) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Shr(stack.PopLong()))
}

func (lstore Instructionlstore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[lstore.Index].obj = stack.PopLong()
}

func (lstore0 Instructionlstore0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[0].obj = stack.PopLong()
}

func (lstore1 Instructionlstore1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[1].obj = stack.PopLong()
}

func (lstore2 Instructionlstore2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[2].obj = stack.PopLong()
}

func (lstore3 Instructionlstore3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	lv[3].obj = stack.PopLong()
}

func (lsub Instructionlsub) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Sub(stack.PopLong()))
}

func (lushr Instructionlushr) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Ushr(stack.PopLong()))
}

func (lxor Instructionlxor) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushLong(stack.PopLong().Xor(stack.PopLong()))
}

func (monitorenter Instructionmonitorenter) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (monitorexit Instructionmonitorexit) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (multianewarray Instructionmultianewarray) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (new Instructionnew) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	cls := rcp.infos[new.Index].(RClass)
	obj := frame.classLoader.createObject(env, cls.name)

	stack.PushRef(obj)
}

func virtualArrayType(type0 byte) byte {
	switch type0 {
	case 4:
		return jaBoolean
	case 5:
		return jaChar
	case 6:
		return jaFloat
	case 7:
		return jaDouble
	case 8:
		return jaByte
	case 9:
		return jaShort
	case 10:
		return jaInt
	case 11:
		return jaLong
	default:
		panic("Unknown virtual type")
	}
}

func (newarray Instructionnewarray) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	virtual := virtualArrayType(newarray.Atype)
	array := env.createArray(virtual, stack.PopInt().value)

	stack.PushRef(array)
}

func (nop Instructionnop) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// do nothing
}

func (pop Instructionpop) Execute(_ *JEnv, _ *JMetaObject, _ RuntimeConstantPool, _ []RuntimeLocalVariable, stack *Stack, _ *Frame) {
	stack.PopRef()
}

func (pop2 Instructionpop2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (putfield Instructionputfield) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (putstatic Instructionputstatic) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (ret Instructionret) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (_return0 Instructionreturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	_return(env, nil)
}

func (saload Instructionsaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	aload(env)
}

func (sastore Instructionsastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	astore(env)
}

func (sipush Instructionsipush) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	stack.PushShort(NewPrimitiveShortRD(sipush.Value))
}

func (swap Instructionswap) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	v1 := stack.PopRef()
	v2 := stack.PopRef()

	stack.PushRef(v1)
	stack.PushRef(v2)
}

func (tableswitch Instructiontableswitch) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}

func (wide Instructionwide) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack, frame *Frame) {
	// TODO
}
