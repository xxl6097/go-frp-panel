package main

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"os"
	"path/filepath"
)

func main() {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}

	modFile, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Module Path: ", filepath.Base(filepath.Dir(modFile.Module.Mod.Path)))
	fmt.Println("Module Path: ", filepath.Base(modFile.Module.Mod.Path))
	fmt.Println("Module Path: ", modFile.Module.Mod.Path)
}
