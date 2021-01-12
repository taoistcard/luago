package state

func (st *luaState) GetTop() int {
	return st.stack.top
}

func (st *luaState) AbsIndex(idx int) int {
	return st.stack.absIndex(idx)
}

func (st *luaState) CheckStack(n int) bool {
	st.stack.check(n)
	return true
}

func (st *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		st.stack.pop()
	}
	//st.SetTop(-n-1)
}

func (st *luaState) Copy(fromIdx, toIdx int) {
	val := st.stack.get(fromIdx)
	st.stack.set(toIdx, val)
}

func (st *luaState) PushValue(idx int) {
	val := st.stack.get(idx)
	st.stack.push(val)
}

func (st *luaState) Replace(idx int) {
	val := st.stack.pop()
	st.stack.set(idx, val)
}

func (st *luaState) Insert(idx int) {
	st.Rotate(idx, 1)
}

func (st *luaState) Remove(idx int) {
	st.Rotate(idx, -1)
	st.Pop(1)
}

func (st *luaState) Rotate(idx, n int) {
	t := st.stack.top - 1
	p := st.stack.absIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	st.stack.reverse(p, m)
	st.stack.reverse(m+1, t)
	st.stack.reverse(p, t)
}

func (st *luaState) SetTop(idx int) {
	newTop := st.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow!")
	}
	n := st.stack.top - newTop
	if n > 0 {
		for i := 0; i < n; i++ {
			st.stack.pop()
		}
	} else if n < 0 {
		for i := 0; i > n; i-- {
			st.stack.push(nil)
		}
	}
}
