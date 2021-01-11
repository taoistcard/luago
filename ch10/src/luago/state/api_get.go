package state

import . "luago/api"

func (self *luaState) CreateTable(nArr, nRec int) {
	t := newLuaTable(nArr, nRec)
	self.stack.push(t)
}

func (self *luaState) NewTable() {
	self.CreateTable(0, 0)
}

func (self *luaState) GetTable(idx int) LuaType {
	t := self.stack.get(idx)
	k := self.stack.pop()
	return self.getTable(t, k)
}

func (self *luaState) getTable(t, k luaValue) LuaType {
	if tbl, ok := t.(*luaTable); ok {
		v := tbl.get(k)
		self.stack.push(v)
		return typeOf(v)
	}
	panic("not a table!")
}

/*
也可以先把k入栈，然后调用GetTable(idx)实现，但是不够高效
*/
func (self *luaState) GetField(idx int, k string) LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, k)
}

func (self *luaState) GetI(idx int, i int64) LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, i)
}

func (self *luaState) GetGlobal(name string) LuaType {
	t := self.registry.get(LUA_RIDX_GLOBALS)
	return self.getTable(t, name)
}

/*
func (self *luaState) GetGlobal(name string) LuaType {
	self.PushGlobalTable()
	return self.GetField(-1, name)
}
*/
