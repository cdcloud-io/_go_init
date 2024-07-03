package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Expected major version for compatibility
const expectedMajorVersion = 1

// Function to extract template version
func extractTemplateVersion(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "<<<tmpl_version:") && strings.Contains(line, ">>>") {
			start := strings.Index(line, "<<<tmpl_version:") + len("<<<tmpl_version:")
			end := strings.Index(line, ">>>")
			if start < end {
				version := line[start:end]
				return version, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("template version not found in template")
}

// Function to parse the version string and return the major version
func parseMajorVersion(version string) (int, error) {
	var major int
	_, err := fmt.Sscanf(version, "%d.%d.%d", &major, new(int), new(int))
	if err != nil {
		return 0, fmt.Errorf("invalid version format: %s", version)
	}
	return major, nil
}

func main() {
	templateFile := "main.go.tmpl"

	version, err := extractTemplateVersion(templateFile)
	if err != nil {
		panic(err)
	}

	majorVersion, err := parseMajorVersion(version)
	if err != nil {
		panic(err)
	}

	if majorVersion != expectedMajorVersion {
		fmt.Printf("Template version %s is not compatible with expected major version %d\n", version, expectedMajorVersion)
		os.Exit(1)
	}

	fmt.Println("Template version is compatible")

	data := struct {
		AppName string
		App     string
		Version string
	}{
		AppName: "myapp",
		App:     "github.com/myorg/myapp",
		Version: version,
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("cmd/%s/main.go", data.AppName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}
