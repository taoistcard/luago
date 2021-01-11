package main

import (
	"fmt"
	"io/ioutil"
	. "luago/api"
	"luago/state"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		ls := state.New()
		ls.Register("print", print)
		ls.Load(data, os.Args[1], "b")
		ls.Call(0, 0)
	}
}

func printStack(ls LuaState) {
	top := ls.GetTop()
	for i := 1; i <= top; i++ {
		t := ls.Type(i)
		switch t {
		case LUA_TBOOLEAN:
			fmt.Printf("[%t]", ls.ToBoolean(i))
		case LUA_TNUMBER:
			fmt.Printf("[%g]", ls.ToNumber(i))
		case LUA_TSTRING:
			fmt.Printf("[%q]", ls.ToString(i))
		default:
			fmt.Printf("[%s]", ls.TypeName(t))
		}
	}
	fmt.Println()
}

func print(ls LuaState) int {
	nArgs := ls.GetTop()
	for i := 1; i <= nArgs; i++ {
		if ls.IsBoolean(i) {
			fmt.Printf("%t", ls.ToBoolean(i))
		} else if ls.IsString(i) {
			fmt.Print(ls.ToString(i))
		} else {
			fmt.Print(ls.TypeName(ls.Type(i)))
		}
		if i < nArgs {
			fmt.Print("\t")
		}
	}
	fmt.Println()
	return 0
}
