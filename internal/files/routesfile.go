package files

import (
	"os"
)

func WriteRoutesFile(routesFilePath string) error {
	packageContent := []byte(
		`package routes

// Seu código aqui
`)

	routesFile, err := os.Create(routesFilePath)
	if err != nil {
		return err
	}
	defer routesFile.Close()

	_, err = routesFile.Write(packageContent)
	if err != nil {
		return err
	}

	return nil
}
