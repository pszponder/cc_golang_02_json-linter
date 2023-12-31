package main

import (
	"fmt"

	"github.com/pszponder/cc_golang_02_json-linter/internal/args"
	"github.com/pszponder/cc_golang_02_json-linter/internal/lexer"
)

func main() {
	// Retrieve filepath to the file to validate
	filePath := args.GetFilePath()
	fmt.Println(filePath)

	// Pass in the file to a lexer in order to generate a token representation of the file (tokenize it)
	tokens := lexer.Lex(filePath)
	for _, tok := range tokens {
		fmt.Println(tok)
	}

	// TODO: Pass the tokens to a parser to generate an AST

	// TODO: Validate the JSON using the generated AST

	// TODO: Output 0 or 1 to stdout depending if parsed file is valid JSON
}
