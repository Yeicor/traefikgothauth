package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
)

const providersPackage = "github.com/markbates/goth/providers"

var providersDir = "vendor/" + providersPackage

func main() {
	err := os.WriteFile("providers.go", []byte(GenerateProviders().String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command("go", "fmt", "providers.go").Run()
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateProviders generates the providers list of metadata and constructor functions
func GenerateProviders() *strings.Builder {
	imports := &strings.Builder{}
	imports.WriteString("package traefikgothauth\n\n")
	imports.WriteString("import (\n")
	imports.WriteString("\t\"fmt\"\n")
	imports.WriteString("\t\"" + path.Dir(providersPackage) + "\"\n")
	providerDirs, err := os.ReadDir(providersDir)
	if err != nil {
		log.Fatal(err)
	}
	contents := &strings.Builder{}
	for _, file := range providerDirs {
		if file.IsDir() && GenerateProvider(file.Name(), contents) {
			imports.WriteString("\t\"" + providersPackage + "/" + file.Name() + "\"\n")
			// TODO: Save snapshot, test yaegi build and rollback on failure
		}
	}
	imports.WriteString(")\n\n")
	imports.WriteString("var allProviders = []*ProviderInfo{\n")
	imports.WriteString(contents.String())
	imports.WriteString("}\n\n")
	return imports
}

// GenerateProvider generates the metadata and constructor function for a provider
func GenerateProvider(providerName string, contents *strings.Builder) bool {
	providerNameTitle := strings.ToUpper(string(providerName[0])) + providerName[1:]
	providerPkgs, err := parser.ParseDir(token.NewFileSet(), providersDir+"/"+providerName, nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return false
	}
	contents.WriteString("\t{\n")
	contents.WriteString("\t\tName: \"" + providerName + "\",\n")
	contents.WriteString("\t\tNew: func(callback string, custom map[string]interface{}) (goth.Provider, error) {\n")
	for name, pkg := range providerPkgs {
		GenerateProviderPackage(providerName, providerNameTitle, name, pkg, contents)
	}
	contents.WriteString("\t\t},\n")
	contents.WriteString("\t},\n")
	return true
}

// GenerateProviderPackage generates the metadata and constructor function for a provider package
func GenerateProviderPackage(providerName, providerNameTitle, packageName string, pkg *ast.Package, contents *strings.Builder) {
	var allFuncs []*ast.FuncDecl
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				if strings.HasPrefix(fn.Name.Name, "New") {
					allFuncs = append(allFuncs, fn)
				}
			}
		}
	}
	// Sort functions by more specific parameters first
	sort.Slice(allFuncs, func(i, j int) bool {
		count1, count2 := 0, 0
		for _, stmt := range allFuncs[i].Type.Params.List {
			count1 += len(stmt.Names)
		}
		for _, stmt := range allFuncs[j].Type.Params.List {
			count2 += len(stmt.Names)
		}
		return count1 > count2
	})
	// Render constructor functions
	contents.WriteString("\t\t\tvar provider goth.Provider\n")
	for _, fn := range allFuncs {
		if strings.HasPrefix(fn.Name.Name, "New") {
			contents.WriteString("\t\t\t// " + providerNameTitle + " > " + fn.Name.Name + "\n")
			contents.WriteString("\t\t\tif provider == nil {\n")
			contents.WriteString("\t\t\tcustom[\"callbackURL\"] = callback\n")
			argCount := 0
			callbackUrlFound := false
			for _, stmt := range fn.Type.Params.List {
				for _, ident := range stmt.Names {
					_, isOptional := stmt.Type.(*ast.Ellipsis)
					if isOptional {
						contents.WriteString("\t\t\tif _, ok := custom[\"" + ident.Name + "\"]; !ok {\n")
						contents.WriteString("\t\t\t\tcustom[\"" + ident.Name + "\"] = make([]" + stmt.Type.(*ast.Ellipsis).Elt.(*ast.Ident).Name + ", 0)\n")
						contents.WriteString("\t\t\t}\n")
					}
					contents.WriteString("\t\t\tif arg" + ident.Name + ", ok := custom[\"" + ident.Name + "\"]; ok {\n")
					if ident.Name == "callbackURL" {
						callbackUrlFound = true
					}
					argCount++
				}
			}
			if !callbackUrlFound {
				panic("Missing callbackURL")
			}
			if len(fn.Type.Results.List) == 1 {
				contents.WriteString("\t\t\tprovider")
			} else if len(fn.Type.Results.List) == 2 {
				contents.WriteString("\t\t\tvar err error\n")
				contents.WriteString("\t\t\tprovider, err")
			} else {
				panic("Unknown number of results")
			}
			contents.WriteString("= " + packageName + "." + fn.Name.Name + "(")
			argCount2 := 0
			for _, stmt := range fn.Type.Params.List {
				for _, ident := range stmt.Names {
					contents.WriteString("arg" + ident.Name)
					if t, ok := stmt.Type.(*ast.Ident); ok {
						contents.WriteString(".(" + t.Name + ")")
					} else if t, ok := stmt.Type.(*ast.Ellipsis); ok {
						if t2, ok := t.Elt.(*ast.Ident); ok {
							contents.WriteString(".([]" + t2.Name + ")...")
						} else {
							panic("Unknown ellipsis type")
						}
					} else if t, ok := stmt.Type.(*ast.ArrayType); ok {
						if t2, ok := t.Elt.(*ast.Ident); ok {
							contents.WriteString(".([]" + t2.Name + ")")
						} else {
							panic("Unknown array type")
						}
					} else {
						panic("Unknown type")
					}
					if argCount2++; argCount2 < argCount {
						contents.WriteString(", ")
					}
				}
			}
			contents.WriteString(")\n")
			if len(fn.Type.Results.List) == 2 {
				contents.WriteString("\t\t\tif err != nil {\n")
				contents.WriteString("\t\t\t\treturn nil, err\n")
				contents.WriteString("\t\t\t}\n")
			}
			for _, stmt := range fn.Type.Params.List {
				for range stmt.Names {
					contents.WriteString("\t\t\t}\n")
				}
			}
			contents.WriteString("\t\t\t}\n")
		}
	}
	contents.WriteString("\t\t\tif provider == nil {\n")
	contents.WriteString("\t\t\t\treturn nil, fmt.Errorf(\"failed to create provider with parameters: %v. Required parameters:")
	for _, fn := range allFuncs {
		if strings.HasPrefix(fn.Name.Name, "New") {
			contents.WriteString("\\n - For " + fn.Name.Name + ": ")
			for _, stmt := range fn.Type.Params.List {
				for _, ident := range stmt.Names {
					contents.WriteString(ident.Name + " ")
				}
			}
		}

	}
	contents.WriteString("\", custom)\n")
	contents.WriteString("\t\t\t}\n")
	contents.WriteString("\t\t\treturn provider, nil\n")
}
