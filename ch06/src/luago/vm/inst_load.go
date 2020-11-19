package vm

import . "luago/api"

func loadNil(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	vm.PushNil()
	for i := a; i < a+b; i++ {
		vm.Copy(-1, i)
	}
	vm.Pop(1)
}

func loadBool(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1 //字节码的索引从0开始

	vm.PushBoolean(b != 0)
	vm.Replace(a)
	if c != 0 {
		vm.AddPC(1)
	}
}

func loadK(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	a += 1
	vm.GetConst(bx)
	vm.Replace(a)
}

func loadKx(i Instruction, vm LuaVM) {
	a, _ := i.ABx()
	a += 1
	ax := Instruction(vm.Fetch()).Ax()
	vm.GetConst(ax)
	vm.Replace(a)
}

func _binaryArith(i Instruction, vm LuaVM, op ArithOp) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRK(b)
	vm.GetRK(c)
	vm.Arith(op)
	vm.Replace(a)
}

func _unaryArith(i Instruction, vm LuaVM, op ArithOp) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.PushValue(b)
	vm.Arith(op)
	vm.Replace(a)
}
