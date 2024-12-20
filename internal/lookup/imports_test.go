package lookup_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ast2 "github.com/bincooo/go-annotation/internal/ast"
	"github.com/bincooo/go-annotation/internal/lookup"
	"github.com/bincooo/go-annotation/internal/module"
)

func TestFindImportByAlias(t *testing.T) {

	m, err := module.Load("./fixtures")
	require.NoError(t, err)

	mast, err := ast2.LoadFileAST("./fixtures/imports.go")
	require.NoError(t, err)

	toTest := []struct {
		alias     string
		importPkg string
	}{
		{"fmt", "fmt"},
		{"mlog", "log"},
		{"slog", "log/slog"},
		{".", "github.com/davecgh/go-spew/spew"},
		{"_", "github.com/davecgh/go-spew/spew"},
		{"spew", "github.com/davecgh/go-spew/spew"},
		{"anythingelse", "github.com/bincooo/go-annotation/internal/lookup/fixtures/dashed-package"},
	}

	for _, s := range toTest {
		t.Run(s.alias, func(t *testing.T) {
			pkg, ok := lookup.FindImportByAlias(m, mast, s.alias)
			require.True(t, ok)
			assert.Equal(t, s.importPkg, pkg)
		})
	}
}
