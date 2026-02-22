// Package runner contains functions for running the compiler pipeline for compiling tarka source code
package runner

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/adinovap20/tarkac/internal/asmgen/lnx64"
	"github.com/adinovap20/tarkac/internal/ast"
	"github.com/adinovap20/tarkac/internal/astprinter"
	"github.com/adinovap20/tarkac/internal/ir"
	"github.com/adinovap20/tarkac/internal/irgen"
	"github.com/adinovap20/tarkac/internal/lexer"
	"github.com/adinovap20/tarkac/internal/parser"
	"github.com/adinovap20/tarkac/internal/semantic"
	"github.com/adinovap20/tarkac/internal/token"
)

// Flags structure contains all the parsed data from the command line arguments or flags
type Flags struct {
	inputFile  *string // Pointer to the path of the string input file path
	outputFile *string // Pointer to the path of the string output file path
	debugFlag  *bool   // Pointer to the debug flag
	fasmFile   *string // Pointer to the path of the string output file path
}

// Run runs the compiler pipeline according to the command line options
func Run() {
	// Parse the command line arguments
	outputFile := flag.String("o", "out", "Output file path")
	debug := flag.Bool("d", false, "Enable debug mode prints")
	fasmFile := flag.String("f", "out.asm", "Fasm file path")
	flag.Parse()
	fmt.Println(*fasmFile)
	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("Usage: tarkac <input_file> [options]\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	inputFile := args[0]
	cmdLineFlags := &Flags{inputFile: &inputFile, outputFile: outputFile, debugFlag: debug, fasmFile: fasmFile}
	if *cmdLineFlags.debugFlag {
		fmt.Printf("Compiling %s to %s with debug flag %v\n",
			*cmdLineFlags.inputFile, *cmdLineFlags.outputFile, *cmdLineFlags.debugFlag,
		)
	}

	// Run lexical analysis phas
	toks := runLexicalAnalysis(cmdLineFlags)
	program := runSyntaxAnalysis(cmdLineFlags, toks)
	runSemanticAnalysis(cmdLineFlags, program)
	irProgram := runIRGeneration(cmdLineFlags, program)
	runCodeGeneration(cmdLineFlags, irProgram)
	runExecutableGeneration(cmdLineFlags)
}

// runLexicalAnalysis runs the lexical analysis phase of the compiler pipeline
func runLexicalAnalysis(flags *Flags) []token.Token {
	if *flags.debugFlag {
		fmt.Println("=== Lexical Analysis ===")
	}
	content, err := os.ReadFile(*flags.inputFile)
	if err != nil {
		log.Fatal(err)
	}
	code := string(content)
	lexer := lexer.NewLexer(code)
	tokens := lexer.GetTokens()
	if *flags.debugFlag {
		token.PrintTokens(tokens)
		fmt.Println("Lexical analysis successful...")
	}
	return tokens
}

// runSyntaxAnalysis runs the syntax analysis phase of the compiler pipeline
func runSyntaxAnalysis(flags *Flags, tokens []token.Token) *ast.Program {
	if *flags.debugFlag {
		fmt.Println("=== Syntax Analysis ===")
	}
	parser := parser.NewParser(tokens)
	program := parser.Parse()
	if program == nil {
		log.Fatal("Syntax analysis failed... Parser returned <nil> abstract syntax tree!")
	}
	parser.PrintErrors()
	if *flags.debugFlag {
		printer := astprinter.NewASTPrinter()
		program.Accept(printer)
		fmt.Println("Syntax analysis successful...")
	}
	return program
}

// runSemanticAnalysis runs the semantic analysis phase of the compiler pipeline
func runSemanticAnalysis(flags *Flags, program *ast.Program) {
	if *flags.debugFlag {
		fmt.Println("=== Semantic Analysis ===")
	}
	analyzer := semantic.NewSemanticAnalyzer()
	program.Accept(analyzer)
	analyzer.PrintErrors()
	if *flags.debugFlag {
		fmt.Println("Semantic analysis successful...")
	}
}

// runIRGeneration runs the IR Generation phase of the compiler pipeline
func runIRGeneration(flags *Flags, program *ast.Program) *ir.IRProgram {
	if *flags.debugFlag {
		fmt.Println("=== IR Generation ===")
	}
	irGenerator := irgen.NewIRGenerator()
	program.Accept(irGenerator)
	if *flags.debugFlag {
		irGenerator.Print()
		fmt.Println("IR generation successful...")
	}
	return irGenerator.IRProgram
}

// runCodeGeneration generates FASM code for the IR
func runCodeGeneration(flags *Flags, program *ir.IRProgram) {
	if *flags.debugFlag {
		fmt.Println("=== Code Generation ===")
	}
	lnx64Generator := lnx64.NewGenerator()
	program.Accept(lnx64Generator)
	if *flags.debugFlag {
		fmt.Println(lnx64Generator.Code)
		fmt.Println("Code generation successful...")
	}
	err := os.WriteFile(*flags.fasmFile, []byte(lnx64Generator.Code), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// runExecutableGeneration generates executable from the FASM code
func runExecutableGeneration(flags *Flags) {
	if *flags.debugFlag {
		fmt.Println("=== Executable Generation ===")
	}
	cmd := exec.Command("fasm", *flags.fasmFile, *flags.outputFile)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	if *flags.debugFlag {
		fmt.Println("Executable generation successful...")
	}
}
