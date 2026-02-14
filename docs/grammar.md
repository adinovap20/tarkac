# Grammar

## Lexical Analysis Phase

```ebnf
KW_EXIT    -> "exit"

LIT_INT    -> [1-9][0-9]*
LIT_IDENT  -> [a-zA-Z_][a-zA-Z0-9_]*

EX_NEWLINE -> "\n"
```

## Syntax Grammar

```ebnf
Program     -> Statement*

Statement   -> StmtExit
StmtExit    -> KW_EXIT Expression EX_NEWLINE

Expression  -> ExprIntLit
ExprIntLit  -> LIT_INT
```
