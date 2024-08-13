package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/GuiCezaF/go-cli/internal/files"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	projectName string
	projectPath string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create boilerplate for a new project",
	Run: func(cmd *cobra.Command, args []string) {
		if projectName == "" {
			fmt.Println("You must supply a project name.")
			return
		}
		if projectPath == "" {
			fmt.Println("You must supply a project path.")
			return
		}
		fmt.Println("Creating project...")

		globalPath := filepath.Join(projectPath, projectName)

		if _, err := os.Stat(globalPath); err == nil {
			color.Red.Println("Project directory already exists.")
			return
		}

		if err := os.MkdirAll(globalPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		// TODO: quando possuir mais tecnologias verificar se for go e colocar o user do github automaticamente
		gitHubUser := "" // aqui você pode colocar seu usuario do github para criar com o padrão do GO
		startGo := exec.Command("go", "mod", "init", gitHubUser+projectName)
		startGo.Dir = globalPath
		startGo.Stdout = os.Stdout
		startGo.Stderr = os.Stderr
		err := startGo.Run()
		if err != nil {
			log.Fatal(err)
		}

		cmdPath := filepath.Join(globalPath, "cmd")
		if err := os.MkdirAll(cmdPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		internalPath := filepath.Join(globalPath, "internal")
		if err := os.MkdirAll(internalPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		handlerPath := filepath.Join(internalPath, "handler")
		if err := os.MkdirAll(handlerPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		routesPath := filepath.Join(handlerPath, "routes")
		if err := os.MkdirAll(routesPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		mainPath := filepath.Join(cmdPath, "main.go")
		if err := files.WriteMainFile(mainPath); err != nil {
			log.Fatal(err)
		}

		routesFilePath := filepath.Join(routesPath, "routes.go")
		if err := files.WriteRoutesFile(routesFilePath); err != nil {
			log.Fatal(err)
		}

		color.Green.Println("Project created successfully.")
	},
}

func init() {
	createCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project")
	createCmd.Flags().StringVarP(&projectPath, "path", "p", "", "Path where the project will be created")

	rootCommand.AddCommand(createCmd)
}
