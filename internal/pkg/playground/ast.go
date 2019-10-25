package playground

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"strconv"

	"golang.org/x/lint"
	"golang.org/x/tools/go/ast/astutil"
)

// MergePackageFiles merge multiple go files into 1 single file.
// The readers will be closed if they implement io.Closer.
// Remove  the main function out of the result
func mergePackageFiles(pkgName string, fileName string, files map[string]io.Reader) ([]byte, error) {
	fset := token.NewFileSet()
	astFiles := make(map[string]*ast.File)
	imports := make([]*ast.ImportSpec, 0)
	for name, f := range files {
		src, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("invalid input file %s: %w", name, err)
		}
		astFile, err := parser.ParseFile(fset, name, string(src), parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("failed to parse input file %s: %w", name, err)
		}
		astFiles[name] = astFile
		if closer, ok := f.(io.Closer); ok {
			_ = closer.Close()
		}
		// Seem ast has a bug and cause sometime the imports is in middle of the file and failed to print/build.
		// Hence remove them first and will add back later in the merged file.
		for _, specs := range astutil.Imports(fset, astFile) {
			for _, spec := range specs {
				pth, _ := strconv.Unquote(spec.Path.Value)
				_ = astutil.DeleteImport(fset, astFile, pth)
				imports = append(imports, spec)
			}
		}
	}
	pkg := &ast.Package{
		Name:    pkgName,
		Files:   astFiles,
		Scope:   ast.NewScope(nil),
		Imports: make(map[string]*ast.Object),
	}
	pkgFile := ast.MergePackageFiles(pkg, ast.FilterImportDuplicates)
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, pkgFile); err != nil {
		return nil, fmt.Errorf("failed to print merged file: %w", err)
	}

	// Parse the file again to have correct position in the ast.File.
	// And re-import all the original imports.
	pkgFset := token.NewFileSet()
	pkgFile, err := parser.ParseFile(pkgFset, fileName, buf.String(), parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse merged file: %w", err)
	}
	for _, imp := range imports {
		path, err := strconv.Unquote(imp.Path.Value)
		if err != nil {
			fmt.Printf("ERR: invalid import path: %v\n", err)
			continue
		}
		name := ""
		if imp.Name != nil {
			name = imp.Name.Name
		}
		_ = astutil.AddNamedImport(pkgFset, pkgFile, name, path)
	}

	// remove main function as it will cause playground thinks it's not unit test
	for i, decl := range pkgFile.Decls {
		if gen, ok := decl.(*ast.FuncDecl); ok && gen.Name.Name == "main" {
			copy(pkgFile.Decls[i:], pkgFile.Decls[i+1:])
			pkgFile.Decls = pkgFile.Decls[:len(pkgFile.Decls)-1]
			break
		}
	}

	pkgBuf := bytes.Buffer{}
	if err := format.Node(&pkgBuf, pkgFset, pkgFile); err != nil {
		return nil, fmt.Errorf("failed to format merged file: %w", err)
	}
	return pkgBuf.Bytes(), nil
}

// LintFile evaluate the source file against Go lint rules.
func LintFile(file string, src []byte) ([]lint.Problem, error) {
	linter := &lint.Linter{}
	problems, err := linter.Lint(file, src)
	if err != nil {
		return nil, fmt.Errorf("failed to parse lint file: %w", err)
	}
	return problems, nil
}
