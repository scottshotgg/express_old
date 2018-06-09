define i32 @main() {
main:
	%0 = alloca i32
	store i32 10, i32* %0
	%1 = alloca double
	store double 1.1, double* %1
	%2 = alloca i32
	store i32 20, i32* %2
	br label %return
return:
	ret i32 0
}
