#include "textflag.h"
#include "funcdata.h"

TEXT ·Strings(SB), 0, $48-48
        NO_LOCAL_POINTERS // Hack.
        MOVQ x+0(FP), DX
        MOVQ x+8(FP), DI
        MOVQ y+16(FP), CX
        MOVQ y+24(FP), SI
        TESTQ DI, DI
        JZ maybe_return_y // x is "", maybe we can return y without allocs
        TESTQ SI, SI
        JZ maybe_return_x // y is "", maybe we can return x without allocs
concatenate:
        LEAQ (DI)(SI*1), R8 // len(x) + len(y)
        // Allocate storage for new string.
        MOVQ R8, 0(SP)
        MOVQ $0, 8(SP)
        MOVB $0, 16(SP)
        CALL runtime·mallocgc(SB)
        MOVQ 24(SP), AX // allocated str
        MOVQ AX, newstr-8(SP)
        // Copy x into allocated str.
        MOVQ x+0(FP), DX
        MOVQ x+8(FP), DI
        MOVQ AX, 0(SP)
        MOVQ DX, 8(SP)
        MOVQ DI, 16(SP)
        CALL runtime·memmove(SB)
        // Copy y into allocated str at the offset of len(x).
        MOVQ x+8(FP), DI
        MOVQ y+16(FP), CX
        MOVQ y+24(FP), SI
        MOVQ newstr-8(SP), AX
        LEAQ (AX)(DI*1), BX
        MOVQ BX, 0(SP)
        MOVQ CX, 8(SP)
        MOVQ SI, 16(SP)
        CALL runtime·memmove(SB)
        // Return new string.
        MOVQ newstr-8(SP), AX
        MOVQ x+8(FP), R8
        ADDQ y+24(FP), R8
        MOVQ AX, ret+32(FP)
        MOVQ R8, ret+40(FP)
        RET
maybe_return_y:
        MOVQ (TLS), AX // stack
        CMPQ CX, (AX)
        JL return_y // if y_ptr < stk.lo
        CMPQ CX, 8(AX)
        JGE return_y // if y_ptr >= stk.hi
        JMP concatenate // y is on stack, must do a new alloc
return_y:
        MOVQ CX, ret+32(FP)
        MOVQ SI, ret+40(FP)
        RET
maybe_return_x:
        MOVQ (TLS), AX // stack
        CMPQ DX, (AX)
        JL return_x // if x_ptr < stk.lo
        CMPQ DX, 8(AX)
        JGE return_x // if x_ptr >= stk.hi
        JMP concatenate // x is on stack, must do a new alloc
return_x:
        MOVQ DX, ret+32(FP)
        MOVQ DI, ret+40(FP)
        RET
