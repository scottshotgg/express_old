define i32 @main() {
main:
	%0 = alloca i32
	store i32 10, i32* %0
	br label %1
; <label>:1
	ret i32 0
}