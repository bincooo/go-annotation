package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bincooo/go-annotation/internal/logger"
	"go/ast"
	gfmt "go/format"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"text/template"

	"github.com/bincooo/go-annotation/internal/environment"
)

const (
	extGo   = ".go"
	extJSON = ".json"
)

func Store(path string, data []byte, meta map[string]string) error {
	meta["ver"] = environment.ToolVersion()
	meta["go_ver"] = environment.GoVersion()
	ext := filepath.Ext(path)

	data = ingest(ext, data, meta)
	data, err := format(filepath.Ext(path), data)
	if err != nil {
		return err
	}

	if err := mkDir(path); err != nil {
		return fmt.Errorf("unable to create new directories: %w", err)
	}

	if err := os.WriteFile(path, data, os.ModePerm); err != nil {
		return fmt.Errorf("unable to store file %s: %w", path, err)
	}

	return nil
}

func StoreAST(path string, file *ast.File, meta map[string]string) error {
	b := bytes.NewBufferString("")
	err := printer.Fprint(b, &token.FileSet{}, file)
	if err != nil {
		return fmt.Errorf("unable to print file AST: %w", err)
	}

	return Store(path, b.Bytes(), meta)
}

func mkDir(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}
	return nil
}

const doNotEditComment = `
// Code generated by {{if .annotation_name}}{{ .annotation_name }}{{else}}UNDEFINED{{end}} annotation processor. DO NOT EDIT.
// versions:
//		go: {{ .go_ver }}
//		go-annotation: {{ .ver }} 
//		{{if .annotation_name}}{{ .annotation_name }}: {{ .annotation_ver }}{{end}} 
`

func ingest(ext string, data []byte, meta map[string]string) []byte {
	switch ext {
	case extGo:
		b := bytes.NewBufferString("")
		tmpl, err := template.New("do not edit").Parse(doNotEditComment)
		if err != nil {
			return data
		}
		err = tmpl.Execute(b, meta)
		if err != nil {
			return data
		}

		return append(b.Bytes(), data...)
	}

	return data
}

func format(ext string, data []byte) ([]byte, error) {
	switch ext {
	case extGo:
		out, err := gfmt.Source(data)
		if err != nil {
			return nil, err
		}
		return out, nil
	case extJSON:
		var m any
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Warnf("unable to unmarshal JSON for pretty-print: %s", err)
			return data, nil
		}
		d, err := json.MarshalIndent(m, "", "\t")
		if err != nil {
			logger.Warnf("unable to marshal JSON for pretty-print: %w", err)
			return data, nil
		}
		return d, nil
	}
	return data, nil
}
