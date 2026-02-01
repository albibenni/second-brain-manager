package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joho/godotenv"
)

//go:embed README.md
var readmeContent string

func main() {
	// Check for help flag
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		// Create a temporary file with the embedded README content
		tmpDir := os.TempDir()
		tmpFile := filepath.Join(tmpDir, "SecondBrainManager-README.md")

		if err := os.WriteFile(tmpFile, []byte(readmeContent), 0644); err != nil {
			// Fallback: print to terminal if we can't create temp file
			fmt.Println(readmeContent)
			os.Exit(0)
		}

		// Open the temporary README in default browser/application
		cmd := exec.Command("open", tmpFile)
		if err := cmd.Run(); err != nil {
			// If opening fails, print to terminal
			fmt.Println(readmeContent)
			os.Exit(1)
		}

		fmt.Println("Opening documentation...")
		os.Exit(0)
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "SecondBrainManager"
	}

	fmt.Printf("Welcome to %s!\n", appName)
}
