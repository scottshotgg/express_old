	.section	__TEXT,__text,regular,pure_instructions
	.macosx_version_min 10, 12
	.globl	_main                   ## -- Begin function main
	.p2align	4, 0x90
_main:                                  ## @main
	.cfi_startproc
## %bb.0:
	movl	$666, -4(%rsp)          ## imm = 0x29A
	xorl	%eax, %eax
	retq
	.cfi_endproc
                                        ## -- End function

.subsections_via_symbols
