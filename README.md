# Express

## Status: `tock`

Code is pretty dirty right now; it is due for a "garbage collection" since a lot of the code is hacky garbage right now anyways. Will have to fix all of the FIXME and TODO stamps and clean up some of the workarounds later in the next phase.

The status follows a tick-tock style:

- `tick` : push out code to ensure that features are implemented
- `tock` : purge hacky workarounds, optimize, rearch

<br>

## TODO: `todo.md`

I will try to update the todo file as much as possible and keep it up to date as to what is currently needing support/work and what is currently being worked on

<br>

## Stages

### - `lex`:   Tokenize the file into raw lex tokens

### - `parse`: Squash and combine lex tokens into higher structures

### - `check`:   Check the tokens syntactically

### - `llvm`:    Generate LLVM tokens/code and output the binary

Not all stages are started/implemented fully yet.

- `lex` is 99% done
- `parse` is mostly finished
- `check` is just being started
- `llvm` will be worked on intermittently in a back-and-forth style with `check` for a bit while I find the best route for token generation within Go.

<br>

## Grammar

This is mostly correct but needs to be updated

```ebnf
program : stmt+ ;

stmt    : assignment_stmt |
          expr <EOS>      ;

assignment_stmt:  var assign_op expr <EOS>

expr    : term sec_op term |
          term             ;

assign_op : SET    |
            ASSIGN |
            INIT   ;

ASSIGN : `=`
SET    : `:`
INIT   : [ SET ASSIGN ]

sec_op  : ADD     |
          VEC_ADD |
          SUB     |
          VEC_SUB ;

ADD     : `+`  ;
VEC_ADD : `.+` ;
SUB     : `-`  ;
VEC_SUB : `.-` ;

term    : factor pri_op term |
          factor             ;

pri_op  : MULT     |
          DIV      |
          VEC_MULT |
          VEC_DIV  |
          MOD      ;

MULT      : `*`  ;
DIV       : `/`  ;
VEC_MULT  : `.*` ;
VEC_DIV   : `./` ;
MOD       : `%`  ;
VEC_MOD   : `.%` ;

factor  : L_PAREN expr R_PAREN |
          <literal>            |
          ident                ;

L_PAREN : `(` ;
R_PAREN : `)` ;

var     : type ident ;

type    : `var`  |
          number |
          char   |
          string |
          array  |
          bool   ;

ident   : <identifier> ;

(* 
  we need to make something for series/formulas
*)

number  : binary   |
          octal    |
          hex      |
          int      |
          float    |
          fraction |
          radical  |
          complex  ;

binary  : [ `0` | `1` ]+ ;
octal   : [ `0` | `1` | `2` | `3` | `4` | `5` | `6` | `7` ]+ ;
hex     : [ `0` | `1` | `2` | `3` | `4` | `5` | `6` | `7` |
            `8` | `9` |`A` | `B` | `C` | `D` | `E` | `F` ]+ ;
int     : [ `0` | `1` | `2` | `3` | `4` | `5` | `6` | `7` | `8` | `9` ]+ ;
float   : [ int `.` int ]
```
