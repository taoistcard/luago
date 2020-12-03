local t = {"a","b","c","d"}
t[2] = nil
t[4] = nil
--[[main <ch7.lua:0,0> (9 instructions at 0096B2D0)
0+ params, 5 slots, 1 upvalue, 1 local, 7 constants, 0 functions
        1       [1]     NEWTABLE        0 4 0
        2       [1]     LOADK           1 -1    ; "a"
        3       [1]     LOADK           2 -2    ; "b"
        4       [1]     LOADK           3 -3    ; "c"
        5       [1]     LOADK           4 -4    ; "d"
        6       [1]     SETLIST         0 4 1   ; 1
        7       [2]     SETTABLE        0 -5 -6 ; 2 nil
        8       [3]     SETTABLE        0 -7 -6 ; 4 nil
        9       [3]     RETURN          0 1
]]

--[[local t = {"a","b","c"}
t[2] = "B"
t["foo"] = "Bar"
local s = t[3] .. t [2] .. t[1] .. t["foo"] .. #t
]]