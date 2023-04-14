package pkg

import "fmt"

type File struct {
	filePath    string
	fileContent string
}

func (f *File) getFileContent() {
	fmt.Println(f.filePath)
}
