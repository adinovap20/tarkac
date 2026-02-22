# Grammar

## Lexical Analysis Phase

```ebnf
KW_EXIT    -> "exit"
KW_INT     -> "int"

LIT_INT    -> [0-9]*
LIT_IDENT  -> [a-zA-Z_][a-zA-Z0-9_]*

OP_ASSIGN  -> "="

PUNC_COLON -> ":"

EX_NEWLINE -> "\n"
```

## Syntax Grammar

```ebnf
Program         -> Statement*

Statement       -> StmtExit | StmtIntVarDecl
StmtIntVarDecl  -> LIT_IDENT PUNC_COLON KW_INT OP_ASSIGN Expression EX_NEWLINE
StmtExit        -> KW_EXIT Expression EX_NEWLINE

Expression      -> ExprIntLit
ExprIntLit      -> LIT_INT
```
