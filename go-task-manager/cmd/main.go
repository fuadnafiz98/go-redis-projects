package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

type File struct {
	filePath    string
	fileContent string
}

func (f *File) Init() error {
	fmt.Println(f.filePath)
	content, err := os.ReadFile(f.filePath)
	f.fileContent = string(content)
	return err
}

func main() {

	file := File{
		filePath: "cmd/tasks.go",
	}

	err := file.Init()

	if err != nil {
		panic(err)
	}

	db, err := InitDB("database.db")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.db)

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "", file.fileContent, parser.ParseComments)
	if err != nil {
		fmt.Println("ERR", err)
	}
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if strings.HasPrefix(x.Name.Name, "task") {
				lpos := x.Pos()
				rpos := x.Body.Rbrace
				fmt.Println(file.fileContent[lpos-1 : rpos])
			}
		}
		return true
	})
}
