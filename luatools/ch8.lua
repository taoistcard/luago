--[[main <ch8.lua:0,0> (4 instructions at 0037B2D0)
0+ params, 5 slots, 1 upvalue, 5 locals, 0 constants, 2 functions
        1       [2]     LOADNIL         0 2
        2       [3]     CLOSURE         3 0     ; 0037B350
        3       [4]     CLOSURE         4 1     ; 0037B3D0
        4       [4]     RETURN          0 1

function <ch8.lua:3,3> (1 instruction at 0037B350)
0 params, 2 slots, 0 upvalues, 0 locals, 0 constants, 0 functions
        1       [3]     RETURN          0 1

function <ch8.lua:4,4> (1 instruction at 0037B3D0)
0 params, 2 slots, 0 upvalues, 0 locals, 0 constants, 0 functions
        1       [4]     RETURN          0 1

local a,b,c
local function f() end
local g = function() end
]]


--[[main <ch8.lua:0,0> (7 instructions at 0036B2D0)
0+ params, 3 slots, 1 upvalue, 2 locals, 3 constants, 1 function
        1       [3]     CLOSURE         0 0     ; 0036B350
        2       [1]     SETTABUP        0 -1 0  ; _ENV "test"
        3       [5]     GETTABUP        0 0 -1  ; _ENV "test"
        4       [5]     LOADK           1 -2    ; 11
        5       [5]     LOADK           2 -3    ; 22
        6       [5]     CALL            0 3 3
        7       [5]     RETURN          0 1
]]
--[[function test(a,b)
    return a + b, a - b
end

local add,sub=test(11,22)
]]

local function max(...)
    local args = {...}
    local val, idx
    for i=1, #args do
        if val == nil or args[i] > val then
            val, idx = args[i], i
        end
    end
    return val, idx
end

local function assert(v)
    if not v then fail() end
end

local vl = max(3,9,7,128,35)
assert(v1 == 128)

local v2, i2 = max(3,9,7,128,35)
assert(v2 == 128 and i2 == 4)

local v3, i3 = max(3,9,7,128,35)
assert(v3 == 128 and i3 == 4)

local t = {max(3,9,7,128,35)}
assert(t[1] == 128 and t[2] == 4)

--[[main <ch8.lua:0,0> (66 instructions at 0043B2D0)
0+ params, 14 slots, 1 upvalue, 8 locals, 9 constants, 2 functions
        1       [48]    CLOSURE         0 0     ; 0043B350
        2       [52]    CLOSURE         1 1     ; 0043B550
        3       [54]    MOVE            2 0
        4       [54]    LOADK           3 -1    ; 3
        5       [54]    LOADK           4 -2    ; 9
        6       [54]    LOADK           5 -3    ; 7
        7       [54]    LOADK           6 -4    ; 128
        8       [54]    LOADK           7 -5    ; 35
        9       [54]    CALL            2 6 2
        10      [55]    MOVE            3 1
        11      [55]    GETTABUP        4 0 -6  ; _ENV "v1"
        12      [55]    EQ              1 4 -4  ; - 128
        13      [55]    JMP             0 1     ; to 15
        14      [55]    LOADBOOL        4 0 1
        15      [55]    LOADBOOL        4 1 0
        16      [55]    CALL            3 2 1
        17      [57]    MOVE            3 0
        18      [57]    LOADK           4 -1    ; 3
        19      [57]    LOADK           5 -2    ; 9
        20      [57]    LOADK           6 -3    ; 7
        21      [57]    LOADK           7 -4    ; 128
        22      [57]    LOADK           8 -5    ; 35
        23      [57]    CALL            3 6 3
        24      [58]    MOVE            5 1
        25      [58]    EQ              0 3 -4  ; - 128
        26      [58]    JMP             0 2     ; to 29
        27      [58]    EQ              1 4 -7  ; - 4
        28      [58]    JMP             0 1     ; to 30
        29      [58]    LOADBOOL        6 0 1
        30      [58]    LOADBOOL        6 1 0
        31      [58]    CALL            5 2 1
        32      [60]    MOVE            5 0
        33      [60]    LOADK           6 -1    ; 3
        34      [60]    LOADK           7 -2    ; 9
        35      [60]    LOADK           8 -3    ; 7
        36      [60]    LOADK           9 -4    ; 128
        37      [60]    LOADK           10 -5   ; 35
        38      [60]    CALL            5 6 3
        39      [61]    MOVE            7 1
        40      [61]    EQ              0 5 -4  ; - 128
        41      [61]    JMP             0 2     ; to 44
        42      [61]    EQ              1 6 -7  ; - 4
        43      [61]    JMP             0 1     ; to 45
        44      [61]    LOADBOOL        8 0 1
        45      [61]    LOADBOOL        8 1 0
        46      [61]    CALL            7 2 1
        47      [63]    NEWTABLE        7 0 0
        48      [63]    MOVE            8 0
        49      [63]    LOADK           9 -1    ; 3
        50      [63]    LOADK           10 -2   ; 9
        51      [63]    LOADK           11 -3   ; 7
        52      [63]    LOADK           12 -4   ; 128
        53      [63]    LOADK           13 -5   ; 35
        54      [63]    CALL            8 6 0
        55      [63]    SETLIST         7 0 1   ; 1
        56      [64]    MOVE            8 1
        57      [64]    GETTABLE        9 7 -8  ; 1
        58      [64]    EQ              0 9 -4  ; - 128
        59      [64]    JMP             0 3     ; to 63
        60      [64]    GETTABLE        9 7 -9  ; 2
        61      [64]    EQ              1 9 -7  ; - 4
        62      [64]    JMP             0 1     ; to 64
        63      [64]    LOADBOOL        9 0 1
        64      [64]    LOADBOOL        9 1 0
        65      [64]    CALL            8 2 1
        66      [64]    RETURN          0 1

function <ch8.lua:39,48> (21 instructions at 0043B350)
0+ params, 8 slots, 0 upvalues, 7 locals, 2 constants, 0 functions
        1       [40]    NEWTABLE        0 0 0
        2       [40]    VARARG          1 0
        3       [40]    SETLIST         0 0 1   ; 1
        4       [41]    LOADNIL         1 1
        5       [42]    LOADK           3 -1    ; 1
        6       [42]    LEN             4 0
        7       [42]    LOADK           5 -1    ; 1
        8       [42]    FORPREP         3 8     ; to 17
        9       [43]    EQ              1 1 -2  ; - nil
        10      [43]    JMP             0 3     ; to 14
        11      [43]    GETTABLE        7 0 6
        12      [43]    LT              0 1 7
        13      [43]    JMP             0 3     ; to 17
        14      [44]    GETTABLE        7 0 6
        15      [44]    MOVE            2 6
        16      [44]    MOVE            1 7
        17      [42]    FORLOOP         3 -9    ; to 9
        18      [47]    MOVE            3 1
        19      [47]    MOVE            4 2
        20      [47]    RETURN          3 3
        21      [48]    RETURN          0 1

function <ch8.lua:50,52> (5 instructions at 0043B550)
1 param, 2 slots, 1 upvalue, 1 local, 1 constant, 0 functions
        1       [51]    TEST            0 1
        2       [51]    JMP             0 2     ; to 5
        3       [51]    GETTABUP        1 0 -1  ; _ENV "fail"
        4       [51]    CALL            1 1 1
        5       [52]    RETURN          0 1]]