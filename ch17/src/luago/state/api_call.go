package state

import (
	. "luago/api"
	"luago/binchunk"
	"luago/vm"
)

func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	var proto *binchunk.Prototype
	if binchunk.IsBinaryChunk(chunk) {
		proto = binchunk.Undump(chunk)
	} else {
		proto = compiler.Compile(string(chunk), chunkName)
	}

	proto := binchunk.Undump(chunk)
	c := newLuaClosure(proto)
	self.stack.push(c)
	if len(proto.Upvalues) > 0 {
		env := self.registry.get(LUA_RIDX_GLOBALS)
		c.upvals[0] = &upvalue{&env}
	}
	return LUA_OK
}

func (self *luaState) PCall(nArgs, nResults, msgh int) (status int) {
	caller := self.stack
	status = LUA_ERRRUN
	defer func() {
		if err := recover(); err != nil {
			for self.stack != caller {
				self.popLuaStack()
			}
			self.stack.push(err)
		}
	}()
	self.Call(nArgs, nResults)
	status = LUA_OK
	return
}
func (self *luaState) Call(nArgs, nResults int) {
	val := self.stack.get(-(nArgs + 1))
	c, ok := val.(*closure)
	if !ok {
		if mf := getMetafield(val, "__call", self); mf != nil {
			if c, ok = mf.(*closure); ok {
				self.stack.push(val)
				self.Insert(-(nArgs + 2))
				nArgs += 1
			}
		}
	}
	if ok {
		if c.proto != nil {
			// fmt.Printf("call %s<%d,%d>\n",
			// 	c.proto.Source, c.proto.LineDefined,
			// 	c.proto.LastLineDefined)
			self.callLuaClosure(nArgs, nResults, c)
		} else {
			self.callGoClosure(nArgs, nResults, c)
		}
	} else {
		panic("not function!")
	}
}

func (self *luaState) callLuaClosure(nArgs, nResults int, c *closure) {
	nRegs := int(c.proto.MaxStackSize)
	nParams := int(c.proto.NumParams)
	isVararg := (c.proto.IsVararg == 1)

	newStack := newLuaStack(nRegs+LUA_MINSTACK, self)
	newStack.closure = c

	funcAndArgs := self.stack.popN(nArgs + 1)
	newStack.pushN(funcAndArgs[1:], nParams)
	newStack.top = nRegs
	if nArgs > nParams && isVararg { //实际传入的参数比定义的参数多，且此函数有变长参数，需要包多余的参数保存下来
		newStack.varargs = funcAndArgs[nParams+1:]
	}

	self.pushLuaStack(newStack)
	self.runLuaClosure()
	self.popLuaStack()
	if nResults != 0 {
		results := newStack.popN(newStack.top - nRegs)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
	}
}
func (self *luaState) callGoClosure(nArgs, nResults int, c *closure) {
	newStack := newLuaStack(nArgs+LUA_MINSTACK, self)
	newStack.closure = c

	args := self.stack.popN(nArgs)
	newStack.pushN(args, nArgs)
	self.stack.pop()

	self.pushLuaStack(newStack)
	r := c.goFunc(self)
	self.popLuaStack()

	if nResults != 0 {
		results := newStack.popN(r)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
	}
}

func (self *luaState) runLuaClosure() {
	for {
		inst := vm.Instruction(self.Fetch())
		inst.Execute(self)
		if inst.Opcode() == vm.OP_RETURN {
			break
		}
	}
}

func (self *luaStack) popN(n int) []luaValue {
	vals := make([]luaValue, n)
	for i := n - 1; i >= 0; i-- {
		vals[i] = self.pop()
	}
	return vals
}

func (self *luaStack) pushN(vals []luaValue, n int) {
	nVals := len(vals)
	if n < 0 { //这种情况把所有vals入栈
		n = nVals
	}
	for i := 0; i < n; i++ {
		if i < nVals {
			self.push(vals[i])
		} else {
			self.push(nil)
		}
	}
}