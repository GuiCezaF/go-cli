package files

import (
	"os"
)

func WriteMainFile(mainPath string) error {
	packageContent := []byte(
		`package main
		
import "fmt"

func main() {
	fmt.Println("Hello World!")
}
`)

	mainFile, err := os.Create(mainPath)
	if err != nil {
		return err
	}
	defer mainFile.Close()

	_, err = mainFile.Write(packageContent)
	if err != nil {
		return err
	}

	return nil
}
