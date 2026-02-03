package utils

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Helper(embedddedReadme string) {

	// Check for help flag
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		// Create a temporary file with the embedded README content
		tmpDir := os.TempDir()
		tmpFile := filepath.Join(tmpDir, "SecondBrainManager-README.md")

		if err := os.WriteFile(tmpFile, []byte(embedddedReadme), 0644); err != nil {
			// Fallback: print to terminal if we can't create temp file
			fmt.Println(embedddedReadme)
			os.Exit(0)
		}

		// Open the temporary README in default browser/application
		cmd := exec.Command("open", tmpFile)
		if err := cmd.Run(); err != nil {
			// If opening fails, print to terminal
			fmt.Println(embedddedReadme)
			os.Exit(1)
		}

		fmt.Println("Opening documentation...")
		os.Exit(0)
	}
}
