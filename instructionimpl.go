package main

func RunCode(instr []Instruction, lv []RuntimeLocalVariable, env *JEnv, this *JMetaObject, stackSize int) {
	for _, i := range instr {
		i.Execute(env, this, *this.constantPool, lv, StackOf(stackSize))
	}
}

func (aaload Instructionaaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	index := stack.PopInt()
	arrayref := stack.PopRef().(JArray)

	if index.value < 0 || index.value >= arrayref.ArrayLengthD() {
		panic("java.lang.ArrayIndexOutOfBoundsException")
	}

	stack.PushRef(arrayref.GetRef(index))
}

func (aastore Instructionaastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	ref := stack.PopRef()
	index := stack.PopInt()
	arrayref := stack.PopRef().(JArray)

	if index.value < 0 || index.value >= arrayref.ArrayLengthD() {
		panic("java.lang.ArrayIndexOutOfBoundsException")
	}

	arrayref.SetRef(index, ref)
}

func (aconstNull InstructionaconstNull) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	stack.PushRef(nil)
}

func (aload Instructionaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	stack.PushRef(lv[aload.Index].obj)
}

func (aload0 Instructionaload0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	stack.PushRef(lv[0].obj)
}

func (aload1 Instructionaload1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	stack.PushRef(lv[1].obj)
}

func (aload2 Instructionaload2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	stack.PushRef(lv[2].obj)
}

func (aload3 Instructionaload3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	stack.PushRef(lv[3].obj)
}

func (anewarray Instructionanewarray) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	//count := stack.PopInt()
	//
	//if count.value < 0 {
	//	panic("java.lang.NegativeArraySizeException")
	//}
	//
	//stack.PushRef(NewJArray(count.value, rcp.GetClass(anewarray.Index)))

	// very much todo
}

func (areturn Instructionareturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (arraylength Instructionarraylength) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (astore Instructionastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (astore0 Instructionastore0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (astore1 Instructionastore1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (astore2 Instructionastore2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (astore3 Instructionastore3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (athrow Instructionathrow) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (baload Instructionbaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (bastore Instructionbastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (bipush Instructionbipush) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (breakpoint Instructionbreakpoint) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (caload Instructioncaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (castore Instructioncastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (checkcast Instructioncheckcast) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (d2f Instructiond2f) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (d2i Instructiond2i) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (d2l Instructiond2l) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dadd Instructiondadd) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (daload Instructiondaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dastore Instructiondastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dcmpg Instructiondcmpg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dcmpl Instructiondcmpl) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dconst0 Instructiondconst0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dconst1 Instructiondconst1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ddiv Instructionddiv) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dload Instructiondload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dload0 Instructiondload0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dload1 Instructiondload1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dload2 Instructiondload2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dload3 Instructiondload3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dmul Instructiondmul) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dneg Instructiondneg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (drem Instructiondrem) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dreturn Instructiondreturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dstore Instructiondstore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dstore0 Instructiondstore0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dstore1 Instructiondstore1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dstore2 Instructiondstore2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dstore3 Instructiondstore3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dsub Instructiondsub) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dup Instructiondup) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dupX1 InstructiondupX1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dupX2 InstructiondupX2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dup2 Instructiondup2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dup2X1 Instructiondup2X1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (dup2X2 Instructiondup2X2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (f2d Instructionf2d) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (f2i Instructionf2i) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (f2l Instructionf2l) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fadd Instructionfadd) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (faload Instructionfaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fastore Instructionfastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fcmpg Instructionfcmpg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fcmpl Instructionfcmpl) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fconst0 Instructionfconst0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fconst1 Instructionfconst1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fconst2 Instructionfconst2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fdiv Instructionfdiv) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fload Instructionfload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fload0 Instructionfload0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fload1 Instructionfload1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fload2 Instructionfload2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fload3 Instructionfload3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fmul Instructionfmul) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fneg Instructionfneg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (frem Instructionfrem) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (freturn Instructionfreturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fstore Instructionfstore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fstore0 Instructionfstore0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fstore1 Instructionfstore1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fstore2 Instructionfstore2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fstore3 Instructionfstore3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (fsub Instructionfsub) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (getfield Instructiongetfield) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (getstatic Instructiongetstatic) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (goto_ Instructiongoto) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (gotoW InstructiongotoW) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (i2b Instructioni2b) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (i2c Instructioni2c) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (i2d Instructioni2d) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (i2f Instructioni2f) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (i2l Instructioni2l) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (i2s Instructioni2s) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iadd Instructioniadd) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iaload Instructioniaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iand Instructioniand) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iastore Instructioniastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iconstM1 InstructioniconstM1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iconst0 Instructioniconst0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iconst1 Instructioniconst1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iconst2 Instructioniconst2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iconst3 Instructioniconst3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iconst4 Instructioniconst4) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iconst5 Instructioniconst5) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (idiv Instructionidiv) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifAcmpeq InstructionifAcmpeq) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifAcmpne InstructionifAcmpne) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifIcmpeq InstructionifIcmpeq) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifIcmpge InstructionifIcmpge) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifIcmpgt InstructionifIcmpgt) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifIcmple InstructionifIcmple) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifIcmplt InstructionifIcmplt) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifIcmpne InstructionifIcmpne) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifeq Instructionifeq) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifge Instructionifge) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifgt Instructionifgt) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifle Instructionifle) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iflt Instructioniflt) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifne Instructionifne) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifnonnull Instructionifnonnull) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ifnull Instructionifnull) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iinc Instructioniinc) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iload Instructioniload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iload0 Instructioniload0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iload1 Instructioniload1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iload2 Instructioniload2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iload3 Instructioniload3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (impdep1 Instructionimpdep1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (impdep2 Instructionimpdep2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (imul Instructionimul) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ineg Instructionineg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (instanceof Instructioninstanceof) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (invokedynamic Instructioninvokedynamic) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (invokeinterface Instructioninvokeinterface) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (invokespecial Instructioninvokespecial) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (invokestatic Instructioninvokestatic) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (invokevirtual Instructioninvokevirtual) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ior Instructionior) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (irem Instructionirem) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ireturn Instructionireturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ishl Instructionishl) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ishr Instructionishr) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (istore Instructionistore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (istore0 Instructionistore0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (istore1 Instructionistore1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (istore2 Instructionistore2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (istore3 Instructionistore3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (isub Instructionisub) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (iushr Instructioniushr) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ixor Instructionixor) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (jsr Instructionjsr) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (jsrW InstructionjsrW) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (l2d Instructionl2d) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (l2f Instructionl2f) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (l2i Instructionl2i) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ladd Instructionladd) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (laload Instructionlaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (land Instructionland) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lastore Instructionlastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lcmp Instructionlcmp) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lconst0 Instructionlconst0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lconst1 Instructionlconst1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ldc Instructionldc) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	rString := rcp.infos[ldc.Index].(*RString).value

	println(rString)

	// TODO
}

func (ldcW InstructionldcW) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ldc2W Instructionldc2W) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ldiv Instructionldiv) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lload Instructionlload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lload0 Instructionlload0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lload1 Instructionlload1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lload2 Instructionlload2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lload3 Instructionlload3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lmul Instructionlmul) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lneg Instructionlneg) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lookupswitch Instructionlookupswitch) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lor Instructionlor) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lrem Instructionlrem) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lreturn Instructionlreturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lshl Instructionlshl) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lshr Instructionlshr) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lstore Instructionlstore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lstore0 Instructionlstore0) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lstore1 Instructionlstore1) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lstore2 Instructionlstore2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lstore3 Instructionlstore3) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lsub Instructionlsub) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lushr Instructionlushr) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (lxor Instructionlxor) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (monitorenter Instructionmonitorenter) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (monitorexit Instructionmonitorexit) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (multianewarray Instructionmultianewarray) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (new Instructionnew) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (newarray Instructionnewarray) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (nop Instructionnop) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (pop Instructionpop) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (pop2 Instructionpop2) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (putfield Instructionputfield) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (putstatic Instructionputstatic) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (ret Instructionret) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (_return Instructionreturn) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (saload Instructionsaload) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (sastore Instructionsastore) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (sipush Instructionsipush) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (swap Instructionswap) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (tableswitch Instructiontableswitch) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}

func (wide Instructionwide) Execute(env *JEnv, this *JMetaObject, rcp RuntimeConstantPool, lv []RuntimeLocalVariable, stack *Stack) {
	// TODO
}
