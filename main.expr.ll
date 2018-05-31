define i32 @main() {
main:
	%0 = alloca double
	store double 55.0, double* %0
	%1 = alloca [11 x i32]
	store [11 x i32] [i32 104, i32 101, i32 121, i32 32, i32 105, i32 116, i32 115, i32 32, i32 121, i32 111, i32 117], [11 x i32]* %1
	%2 = alloca [3 x i32]
	store [3 x i32] [i32 103, i32 101, i32 101], [3 x i32]* %2
	%3 = alloca [9 x i32]
	store [9 x i32] [i32 103, i32 101, i32 101, i32 32, i32 97, i32 110, i32 100, i32 32, i32 105], [9 x i32]* %3
	%4 = alloca i32
	store i32 40, i32* %4
	%5 = alloca double
	store double 10.16, double* %5
	%6 = alloca i8
	store i8 0, i8* %6
	%7 = alloca i8
	store i8 1, i8* %7
	%8 = alloca i8
	store i8 1, i8* %8
	br label %9
; <label>:9
	ret i32 0
}
