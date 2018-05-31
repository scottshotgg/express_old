define i32 @main() {
main:
	%0 = alloca { [1 x i32], i32, double, [1 x i32], i8 }
	store { [1 x i32], i32, double, [1 x i32], i8 } { [1 x i32] [i32 32], i32 0, double 0.0, [1 x i32] [i32 32], i8 0 }, { [1 x i32], i32, double, [1 x i32], i8 }* %0
	br label %1
; <label>:1
	ret i32 0
}
