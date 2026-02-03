package utils

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// CheckHelpFlag checks if the provided arguments contain a help flag
func CheckHelpFlag(args []string) bool {
	if len(args) > 1 {
		return args[1] == "--help" || args[1] == "-h"
	}
	return false
}

// CreateTempReadme creates a temporary file with the README content
func CreateTempReadme(content string) (string, error) {
	tmpDir := os.TempDir()
	tmpFile := filepath.Join(tmpDir, "SecondBrainManager-README.md")
	
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		return "", err
	}
	
	return tmpFile, nil
}

// OpenFile opens a file with the system's default application
func OpenFile(path string) error {
	cmd := exec.Command("open", path)
	return cmd.Run()
}

// ShowHelp displays the help documentation
func ShowHelp(readme string) error {
	tmpFile, err := CreateTempReadme(readme)
	if err != nil {
		// Fallback: print to terminal if we can't create temp file
		fmt.Println(readme)
		return err
	}
	
	if err := OpenFile(tmpFile); err != nil {
		// If opening fails, print to terminal
		fmt.Println(readme)
		return err
	}
	
	fmt.Println("Opening documentation...")
	return nil
}

// Helper checks for help flag and displays documentation
func Helper(embedddedReadme string) {
	if CheckHelpFlag(os.Args) {
		if err := ShowHelp(embedddedReadme); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}
}
