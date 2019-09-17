
TEXT ·selectU64WithPDEP(SB),$0-24
        // 1<<ith in AX
	MOVQ 	ith+8(FP), CX
	MOVQ    $0x1, AX
	SHLQ    CX, AX

	// mask is in BX
	MOVQ 	word+0(FP), BX

	// place "1" in AX to the position specified by mask
	PDEPQ   BX, AX, CX

	// backward scan to the highest "1"
	BSRQ    CX, CX
	MOVQ 	CX, ret+16(FP)
	RET
