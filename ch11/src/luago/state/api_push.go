package state

import . "luago/api"

func (st *luaState) PushNil() {
	st.stack.push(nil)
}

func (st *luaState) PushBoolean(b bool) {
	st.stack.push(b)
}

func (st *luaState) PushInteger(n int64) {
	st.stack.push(n)
}

func (st *luaState) PushNumber(m float64) {
	st.stack.push(m)
}

func (st *luaState) PushString(s string) {
	st.stack.push(s)
}

func (st *luaState) PushGoFunction(f GoFunction, n int) {
	c := newGoClosure(f, n)
	for i := n; i > 0; i-- {
		val := st.stack.pop()
		c.upvals[i-1] = &upvalue{&val}
	}
	st.stack.push(c)
}

func (st *luaState) PushGlobalTable() {
	global := st.registry.get(LUA_RIDX_GLOBALS)
	st.stack.push(global)
}

/*
func (self *luaState) PushGlobalTable() {
	self.GetI(LUA_REGISTRYINDEX, LUA_RIDX_GLOBALS)
}*/
