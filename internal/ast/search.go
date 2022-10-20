package ast

import (
	"go/ast"
)

// findHighLevelNode finds all types, functions, methods, constants and variables defined in the file at the package level
func findHighLevelNode(file *ast.File, nodeName string) (ast.Node, bool) {
	return findNodeInDeclsByName(file.Decls, nodeName)
}

func findNodeInDeclsByName(decls []ast.Decl, name string) (ast.Node, bool) {
	for _, decl := range decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Name.Name == name {
				return d, true
			}
		case *ast.GenDecl:
			if n, ok := findNodeInSpecsByName(d.Specs, name, d.Doc); ok {
				return n, true
			}
		}
	}
	return nil, false
}

func findNodeInSpecsByName(specs []ast.Spec, name string, genCommentGroup *ast.CommentGroup) (ast.Node, bool) {
	for _, spec := range specs {
		switch s := spec.(type) {
		case *ast.ValueSpec:
			for _, ident := range s.Names {
				if ident.Name == name {
					if len(specs) == 1 {
						s.Doc = mergeCommentGroups(s.Doc, genCommentGroup)
					}
					return s, true
				}
			}
		case *ast.TypeSpec:
			if s.Name.Name == name {
				if len(specs) == 1 {
					s.Doc = mergeCommentGroups(s.Doc, genCommentGroup)
				}
				return s, true
			}
		}
	}
	return nil, false
}

// TODO Get rid of the hacky merge
func mergeCommentGroups(g1, g2 *ast.CommentGroup) *ast.CommentGroup {
	if g1 == nil {
		return g2
	}
	if g2 == nil {
		return g1
	}

	var comments []*ast.Comment
	copy(comments, g1.List)
	return &ast.CommentGroup{
		List: append(comments, g2.List...),
	}
}
