package state

import . "luago/api"

type luaStack struct {
	//虚拟栈
	slots []luaValue
	top   int

	//函数调用信息
	state   *luaState
	closure *closure
	varargs []luaValue
	pc      int

	//链表结构
	prev *luaStack

	openuvs map[int]*upvalue
}

func newLuaStack(size int, state *luaState) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
		state: state,
	}
}

func (sk *luaStack) check(n int) {
	free := len(sk.slots) - sk.top
	for i := free; i < n; i++ {
		sk.slots = append(sk.slots, nil)
	}
}

func (sk *luaStack) push(val luaValue) {
	if sk.top == len(sk.slots) {
		panic("stack overflow!")
	}
	sk.slots[sk.top] = val
	sk.top++
}

func (sk *luaStack) pop() luaValue {
	if sk.top < 1 {
		panic("stack underflow!")
	}
	sk.top--
	val := sk.slots[sk.top]
	sk.slots[sk.top] = nil
	return val
}

func (sk *luaStack) absIndex(idx int) int {
	if idx >= 0 || idx <= LUA_REGISTRYINDEX {
		return idx
	}
	return idx + sk.top + 1
}

func (sk *luaStack) isValid(idx int) bool {
	if idx < LUA_REGISTRYINDEX {
		//upvalues
		uvIdx := LUA_REGISTRYINDEX - idx - 1 //从0开始的索引
		c := sk.closure
		return c != nil && uvIdx < len(c.upvals)
	}
	if idx == LUA_REGISTRYINDEX {
		return true
	}
	absIdx := sk.absIndex(idx)
	return absIdx > 0 && absIdx <= sk.top
}

func (sk *luaStack) get(idx int) luaValue {
	if idx < LUA_REGISTRYINDEX {
		//upvalues
		uvIdx := LUA_REGISTRYINDEX - idx - 1
		c := sk.closure
		if c == nil || uvIdx >= len(c.upvals) {
			return nil
		}
		return *(c.upvals[uvIdx].val)
	}
	if idx == LUA_REGISTRYINDEX {
		return sk.state.registry
	}
	absIdx := sk.absIndex(idx)
	if absIdx > 0 && absIdx <= sk.top {
		return sk.slots[absIdx-1]
	}
	return nil
}

func (sk *luaStack) set(idx int, val luaValue) {
	if idx < LUA_REGISTRYINDEX {
		//upvalues
		uvIdx := LUA_REGISTRYINDEX - idx - 1
		c := sk.closure
		if c != nil && uvIdx < len(c.upvals) {
			*(c.upvals[uvIdx].val) = val
		}
		return
	}
	if idx == LUA_REGISTRYINDEX {
		sk.state.registry = val.(*luaTable)
		return
	}
	absIdx := sk.absIndex(idx)
	if absIdx > 0 && absIdx <= sk.top {
		sk.slots[absIdx-1] = val
		return
	}
	panic("invalid index!")
}

func (sk *luaStack) reverse(from, to int) {
	slots := sk.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
