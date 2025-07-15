package lookup

import (
	"go/ast"
	"os"
	"path/filepath"
	"slices"
	"strings"

	ast2 "github.com/bincooo/go-annotation/internal/ast"
	"github.com/bincooo/go-annotation/internal/logger"
	"github.com/bincooo/go-annotation/internal/module"
	"github.com/bincooo/go-annotation/internal/utils"

	. "github.com/bincooo/go-annotation/internal/utils/stream"
)

var (
	projectDir, _ = os.Getwd()
)

func FindImportByAlias(m module.Module, file *ast.File, alias string) (string, bool) {
	for _, spec := range file.Imports {
		if alias == getLocalPackageName(m, spec) {
			return utils.TrimQuotes(spec.Path.Value), true
		}
	}
	return "", false
}

func getLocalPackageName(m module.Module, spec *ast.ImportSpec) string {
	if spec.Name != nil {
		return spec.Name.String()
	}

	if spec.Path == nil {
		return ""
	}

	importPath := utils.TrimQuotes(spec.Path.Value)
	m, err := module.Find(m, importPath)
	if err != nil {
		logger.Warnf("unable to load module for import \"%s\", returning it as is: %w", importPath, err)
		return importPath
	}

	// logger.Debugf("module is found for %s", importPath)

	if m != nil {
		mod := module.Mod(m)
		children := ""
		if strings.HasPrefix(importPath, mod.Module.Mod.Path) {
			if importPath == mod.Module.Mod.Path {
				fast, err := ast2.LoadFileAST(filepath.Join(m.Root(), m.Files()[0]))
				if err != nil {
					panic(err)
				}
				return fast.Name.Name
			}
			children = strings.TrimPrefix(importPath, mod.Module.Mod.Path)
			children = children[1:]
			if runtime.GOOS == "windows" {
				children = strings.Join(strings.Split(children, "/"), string(filepath.Separator))
			}
		}

		filesInPackage := module.FilesInPackage(m, importPath)
		// TODO review the assamption that the first file is a good for check
		files := OfSlice(m.Files()).
			Filter(containsFile(filesInPackage)).
			Filter(func(file string) bool { return strings.HasPrefix(file, children) }).
			ToSlice()
		slices.SortFunc(files, func(file1, file2 string) int { return len(file1) - len(file2) })
		imp := OfSlice(files).
			Map(fileToPackageName(m.Root())).
			One()
		// logger.Debugf("module is found for %s, checking %s", importPath, imp)
		if strings.HasSuffix(imp, "_test") {
			imp = imp[:strings.Index(imp, "_test")]
		}
		return imp
	}

	return strings.ReplaceAll(utils.LastDir(importPath), "-", "_")
}

func fileToPackageName(root string) func(string) string {
	return func(file string) string {
		fast, err := ast2.LoadFileAST(filepath.Join(root, file))
		if err != nil {
			panic(err)
		}
		return fast.Name.Name
	}
}

func containsFile(files []string) func(string) bool {
	return func(file string) bool {
		for _, f := range files {
			if strings.HasSuffix(f, file) && isCompletePathSuffixMatch(f, file) {
				return true
			}
		}
		return false
	}
}

// The check is required, because if we just check suffix for similar package ends:
// log/example_test.go
// log/slog/example_test.go
// In this case both have equel suffix log/example_test.go, but this is completely different packages
func isCompletePathSuffixMatch(str, suffix string) bool {
	prefix := strings.TrimSuffix(str, suffix)
	if len(prefix) == 0 {
		return true
	}
	return prefix[len(prefix)-1] == '/' || prefix[len(prefix)-1] == '\\'
}

func isDir(src, target string) bool {
	src, _ = filepath.Abs(src)
	target, _ = filepath.Abs(target)
	return strings.HasPrefix(target, src)
}
