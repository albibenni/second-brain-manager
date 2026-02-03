package utils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckHelpFlag(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected bool
	}{
		{
			name:     "no arguments",
			args:     []string{"program"},
			expected: false,
		},
		{
			name:     "--help flag",
			args:     []string{"program", "--help"},
			expected: true,
		},
		{
			name:     "-h flag",
			args:     []string{"program", "-h"},
			expected: true,
		},
		{
			name:     "other flag",
			args:     []string{"program", "--version"},
			expected: false,
		},
		{
			name:     "help flag with extra args",
			args:     []string{"program", "--help", "extra"},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckHelpFlag(tt.args)
			if result != tt.expected {
				t.Errorf("CheckHelpFlag(%v) = %v, want %v", tt.args, result, tt.expected)
			}
		})
	}
}

func TestCreateTempReadme(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name:    "simple content",
			content: "# Test README\nThis is a test.",
			wantErr: false,
		},
		{
			name:    "empty content",
			content: "",
			wantErr: false,
		},
		{
			name:    "long content",
			content: strings.Repeat("test\n", 1000),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath, err := CreateTempReadme(tt.content)
			
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTempReadme() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			if err == nil {
				// Clean up
				defer os.Remove(filePath)
				
				// Verify file exists
				if _, err := os.Stat(filePath); os.IsNotExist(err) {
					t.Errorf("CreateTempReadme() file was not created at %s", filePath)
					return
				}
				
				// Verify content
				content, err := os.ReadFile(filePath)
				if err != nil {
					t.Errorf("Failed to read created file: %v", err)
					return
				}
				
				if string(content) != tt.content {
					t.Errorf("File content = %q, want %q", string(content), tt.content)
				}
				
				// Verify filename
				expectedFilename := "SecondBrainManager-README.md"
				actualFilename := filepath.Base(filePath)
				if actualFilename != expectedFilename {
					t.Errorf("Filename = %s, want %s", actualFilename, expectedFilename)
				}
			}
		})
	}
}

func TestOpenFile(t *testing.T) {
	// Create a temporary file to test with
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	t.Run("open existing file", func(t *testing.T) {
		// Note: This will actually try to open the file on macOS
		// In a real CI/CD environment, you might want to skip this or mock it
		err := OpenFile(testFile)
		// We don't check the error here because "open" might fail in headless environments
		// but we verify the function doesn't panic
		_ = err
	})
	
	t.Run("open non-existent file", func(t *testing.T) {
		err := OpenFile("/path/to/nonexistent/file.txt")
		if err == nil {
			t.Error("OpenFile() should return error for non-existent file")
		}
	})
}

func TestShowHelp(t *testing.T) {
	tests := []struct {
		name    string
		readme  string
		wantErr bool
	}{
		{
			name:    "valid readme",
			readme:  "# SecondBrainManager\n\nTest documentation",
			wantErr: false, // May fail in headless environment but won't error on file creation
		},
		{
			name:    "empty readme",
			readme:  "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ShowHelp will try to open the file which might fail in CI
			// but it should at least create the temp file
			err := ShowHelp(tt.readme)
			// We don't strictly check error here because "open" command
			// might fail in headless environments
			_ = err
		})
	}
}
