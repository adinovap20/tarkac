# Lexical Analysis Phase

$$
\begin{aligned}
\text{KW\_EXIT} &\to \text{exit} \\

\\[1ex]
\text{LIT\_INT} &\to [1-9][0-9]* \\
\text{LIT\_IDENT}  &\to [a-zA-Z\_][a-zA-Z0-9\_]* \\

\\[1ex]
\text{EX\_NEWLINE} &\to \backslash n
\end{aligned}
$$

# Syntax Grammar

$$
\begin{aligned}

\text{Program} &\to [\text{Statement}]^* \\

\\[1ex]
\text{Statement} &\to \text{StmtExit} \\
\text{StmtExit} &\to \textit{KW\_EXIT} \quad \text{Expression} \quad \textit{EX\_NEWLINE} \\

\\[1ex]
\text{Expression} &\to \text{ExprIntLit} \\
\text{ExprIntLit} &\to \textit{LIT\_INT} \\

\end{aligned}
$$