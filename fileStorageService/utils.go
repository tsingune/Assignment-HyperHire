package fileStorageService

import (
	"crypto/rand"
	"fmt"
	"mime/multipart"
)

// SplitFile splits a file into chunks
func SplitFile(file multipart.File, tempDir string) ([]string, error) {
	// Logic to split the file into smaller chunks
	return []string{"chunk1", "chunk2"}, nil
}

// GenerateFileID creates a unique file IDs
func GenerateFileID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
