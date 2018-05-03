# TODO

## MAIN SHIT: <br>
- rebuild data/program flow
- rearch the entire lexing process
- rearch the parsing process and use the `Parsemeta` object
- make a `lex_test.go` and a `parse_test.go` file for individualized tests
- fix all other `TODO:` and `FIXME:` tags before starting any more feature
- implement `getX()` chain: 
  - `getStatement()`
  - `getExpr()`
  - `getTerm()`
  - `getFactor()`

- ultimately end up with a parser that can parse primitives, as well as arrays, and multi-type objects with different types of assignment statements for now


enclosers <br>
array type <br>
object type <br>
Nexted enclosers <br>

add debug printouts

## for now the lexer does the literal identification; for instance, with floats and string, but this should probably move to the parser