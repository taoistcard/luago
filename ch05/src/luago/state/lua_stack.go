package state

type luaStack struct {
	slots []luaValue
	top   int
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
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
	if idx >= 0 {
		return idx
	}
	return idx + sk.top + 1
}

func (sk *luaStack) isValid(idx int) bool {
	absIdx := sk.absIndex(idx)
	return absIdx > 0 && absIdx <= sk.top
}

func (sk *luaStack) get(idx int) luaValue {
	absIdx := sk.absIndex(idx)
	if absIdx > 0 && absIdx <= sk.top {
		return sk.slots[absIdx-1]
	}
	return nil
}

func (sk *luaStack) set(idx int, val luaValue) {
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
