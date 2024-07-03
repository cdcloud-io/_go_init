package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReadModuleName reads the module name from the go.mod file.
func ReadModuleName(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}
	return "", fmt.Errorf("module name not found in %s", filePath)
}

// ProcessTemplateFile replaces {{ .ARTIFACT_NAME }} with the module name in the given file.
func ProcessTemplateFile(filePath, moduleName string) error {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	output := strings.ReplaceAll(string(input), "{{ .ARTIFACT_NAME }}", moduleName)

	if err := os.WriteFile(filePath, []byte(output), 0644); err != nil {
		return err
	}

	return nil
}

// RenameArtifactNameDirs renames any directory named __ARTIFACT_NAME__ to the module name.
func RenameArtifactNameDirs(root, moduleName string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "__ARTIFACT_NAME__" {
			newPath := filepath.Join(filepath.Dir(path), moduleName)
			if err := os.Rename(path, newPath); err != nil {
				return err
			}
		}
		return nil
	})
}

// WalkAndProcess walks through directories and processes .tmpl files and renames directories.
func WalkAndProcess(root, moduleName string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process .tmpl files
		if !info.IsDir() && filepath.Ext(path) == ".tmpl" {
			if err := ProcessTemplateFile(path, moduleName); err != nil {
				return err
			}
		}

		// Rename __ARTIFACT_NAME__ directories
		if info.IsDir() && info.Name() == "__ARTIFACT_NAME__" {
			newPath := filepath.Join(filepath.Dir(path), moduleName)
			if err := os.Rename(path, newPath); err != nil {
				return err
			}
		}

		return nil
	})
}

func main() {
	moduleName, err := ReadModuleName("go.mod")
	if err != nil {
		fmt.Println("Error reading module name:", err)
		return
	}

	root := "."

	if err := WalkAndProcess(root, moduleName); err != nil {
		fmt.Println("Error processing templates and directories:", err)
	}
}
