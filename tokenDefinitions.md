# Semantic definitions in Express:

  ## Number:
  Binary: `0b(0 | 1)+`

  Octal: `0o[0-7]+`

  Hex: `0x[0-9A-za-z]+`

  Int: `[0-9]+`

  Float: `[0-9]*.[0-9]+`

  <br />

  ## Group:
  `(` _variables_ `)`
  
  <br />

  ## Block:
  `{` _statements_ `}`

  <br />

  ## Array:
  `[` _variables_ `]`

  <br />

  ## Function:
  A function is a special case of an object: <br />
  - In `Express`, objects can be thought of (and are, technically) an anonymous scope declaration that can be externally referenced by a variable, the _`object`_. By doing so you are allowed external access to that scope declaration and any variables within it.

  - A function essentially acts as a way to 'reinvoke' the object in order to allow the object to redetermine itself. In doing so, you may also de-structure the return object in order to modify the return values. When de-structuring, the memory space of the return values is statically mapped with the undesired contents being "garbage collected" after the scope is exited. (if garbage collection is allowed) 
  - What really happens to the memory is that in runtime, if a function is called dynamically, then yes, the only way is garbage collection. However, if the function is statically called (read: _compile-time determinant_) then the proper corresponding LLVM tokens are injected to clean up the memory after usage.

  `<group>` `<block>` : `() {` _statements_ `}` <br />
  `<group>` `<group>` `<block>` : `(` _arg1_, _arg2_ `) (` _ret1_, _ret2_ `) {` _statements_ `}` <br />
  `<group>` `<block>` `<group>` : `(` _arg1_, _arg2_ `) {` _statements_ `}` `(` _ret1_, _ret2_ `)` <br />