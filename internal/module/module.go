package module

import (
	"strings"

	"golang.org/x/mod/modfile"
	module2 "golang.org/x/mod/module"

	. "github.com/YReshetko/go-annotation/internal/utils/stream"
)

type module struct {
	root  string
	files []string
	mod   *modfile.File
	// file path (absolute path) - module name
	subModFiles map[string]string
}

func newModule(root string, mod *modfile.File, goFiles []string, subModFiles map[string]string) module {
	return module{
		root:        root,
		files:       goFiles,
		mod:         mod,
		subModFiles: subModFiles,
	}
}

func (m *module) Files() []string {
	return m.files
}

func (m *module) Root() string {
	return m.root
}

func (m *module) hasModFile() bool {
	return m.mod != nil
}

func (m *module) hasSubModules() bool {
	return len(m.subModFiles) > 0
}

func (m *module) isSubModule(importPath string) bool {
	for _, modName := range m.subModFiles {
		if strings.Contains(importPath, modName) {
			return true
		}
	}
	return false
}

func (m *module) subModuleRoot(importPath string) string {
	var modPath string
	var modName string
	for fp, mn := range m.subModFiles {
		if strings.Contains(importPath, mn) && len(mn) > len(modName) {
			modPath = fp
			modName = mn
		}
	}
	return modPath
}

// Usecase:
// m.root: /home/yury/go/src/github.com/YReshetko/go-annotation/examples/constructor
// m.files: /internal/common/common.go, ...
// importPath: github.com/YReshetko/go-annotation/examples/constructor/internal/common
// result - true as m.root + m.files[i] contains importPath
func (m *module) hasImportPath(importPath string) bool {
	path := OfSlice(m.files).
		Map(joinPath(m.root)).
		Filter(contains(importPath)).
		One()
	return len(path) > 0
}

func (m *module) findClosestModuleName(importPath string) string {
	var out []string
	if m.mod != nil && len(m.mod.Require) > 0 {
		for _, sm := range m.mod.Require {
			if sm.Indirect || sm.Syntax == nil {
				continue
			}
			for _, t := range sm.Syntax.Token {
				mn, _, ok := module2.SplitPathVersion(t)
				if !ok {
					continue
				}
				if len(mn) > 0 && strings.HasPrefix(importPath, mn) {
					out = append(out, mn)
				}
			}

		}
	}
	if len(out) == 0 {
		return ""
	}

	ind := 0
	l := len(out[ind])

	for i := 1; i < len(out); i++ {
		if len(out[i]) > l {
			l = len(out[i])
			ind = i
		}
	}

	return out[ind]
}

func (m *module) escapedPath(moduleName string) (string, bool) {
	if m.mod == nil || len(m.mod.Require) == 0 {
		return "", false
	}
	line := findModuleLine(m.mod.Require, moduleName)
	if line == nil {
		return "", false
	}

	if len(line.Token) < 2 {
		return "", false
	}
	p, err := module2.EscapePath(line.Token[0])
	if err != nil {
		return "", false
	}
	v, err := module2.EscapeVersion(line.Token[1])
	if err != nil {
		return "", false
	}
	return p + "@" + v, true
}

func findModuleLine(r []*modfile.Require, moduleName string) *modfile.Line {
	for _, sm := range r {
		if sm.Indirect || sm.Syntax == nil {
			continue
		}
		for _, t := range sm.Syntax.Token {
			mn, _, ok := module2.SplitPathVersion(t)
			if !ok || mn != moduleName {
				continue
			}
			return sm.Syntax
		}
	}
	return nil
}

func moduleName(mod *modfile.File) string {
	if mod == nil ||
		mod.Module == nil ||
		mod.Module.Syntax == nil ||
		len(mod.Module.Syntax.Token) < 2 {
		return ""
	}
	return mod.Module.Mod.Path
}
