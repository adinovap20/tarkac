package ast

// PrintIndentation prints the indentation based on the given depth
func PrintIndentation(depth int) {
	for range depth {
		print("    ")
	}
}

// PrintNilIndentation prints the <nil> with the given indentation
func PrintNilIndentation(depth int) {
	PrintIndentation(depth)
	println("<nil>")
}
