#include "textflag.h" // 因为我们声明函数用到了 NOSPLIT 这样的 flag，所以需要将 textflag.h 包含进来

// func add(a, b int) int
TEXT ·add(SB), NOSPLIT, $0-16
    MOVQ a+0(FP), AX // 参数 a
    MOVQ b+8(FP), BX // 参数 b
    ADDQ BX, AX    // AX += BX
    MOVQ AX, ret+16(FP)   //在返回值得地方存入 AX
    RET                    // 返回

// func sub(a, b int) int
TEXT ·sub(SB), NOSPLIT, $0-16
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    SUBQ BX, AX    // AX -= BX
    MOVQ AX, ret+16(FP)
    RET

TEXT main·multiply(SB), NOSPLIT, $0-16
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    IMULQ BX, AX
    MOVQ AX, ret+16(FP)
    RET

