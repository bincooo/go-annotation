package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

var (
	g_fset = token.NewFileSet()
)

func load(path string) (*ast.File, error) {
	fileSpec, err := parser.ParseFile(g_fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("unable to parse go file AST: %w", err)
	}
	return fileSpec, nil
}

func GlobalFSet() *token.FileSet {
	return g_fset
}
