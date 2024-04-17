package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	// get current file path of this file
	fl, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current file path:", err)
		os.Exit(1)
	}

	// Create a scanner to read from standard input (stdin)

	// Open the go.mod file
	file, err := os.Open(filepath.Join(fl, "tools.go"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	// Read the tool dependencies line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Ignore empty lines
		if !strings.HasPrefix(line, "\t_ \"") {
			continue
		}

		line = strings.TrimSpace(line)
		line = strings.TrimSuffix(line, "\"")
		line = strings.TrimPrefix(line, "_ \"")

		finalLine := strings.Split(line, "/")

		gotit := finalLine[len(finalLine)-1]

		// check if gotit is a version
		if _, err := strconv.Atoi(strings.TrimPrefix(gotit, "v")); err == nil {
			gotit = finalLine[len(finalLine)-2]
		}

		fmt.Println("Installing", gotit, line)

		// Install the tool using "go install"
		// cmd := exec.Command("go", "build", "-mod=vendor", "-o", "bin/"+gotit, line)
		cmd := exec.Command("go", "build", "-o", "./bin/"+gotit, line)

		cmd.Stdout = os.Stdout // Display the output of 'go install'
		cmd.Stderr = os.Stderr // Display any errors

		if err := cmd.Run(); err != nil {
			fmt.Println("Error installing", line, ":", err)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
}
