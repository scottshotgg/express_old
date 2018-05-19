# TODO: need to make a check for the filename here

# Run the compiler which is currently written in Go
go run main.go $1 &&

# Use LLVM to output the binary file from the tokens that we created
llc program.expr.ll -o $1.s &&

# Use clang to take those tokens and compile them into an exe
clang program.expr.s -o $1.exe &&

# Run the exe
./$1.exe