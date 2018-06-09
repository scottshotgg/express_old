# Architecture Notes

- Simple lexer
  - portable token format

- Marpa Parser
  - pluggable parsers/state machines

- Semantic VM
  - light optimization

- LLVM backend for creating the binary
  - use Boehmgc for garbage collection
  - lightweight runtime for dynamic typing, managing green threads, memory management, etc