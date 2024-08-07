package file

import (
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/google/uuid"
)

func GetRootDirectory() string {
	projectDir := os.Getenv("PROJECT_DIR")
	projectName := regexp.MustCompile(`^(.*` + projectDir + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}

func GetFileContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func GenerateUniqueFileName(prefix string) string {
	u := uuid.New()
	hash := generateHash(u.String())
	uniqueFilename := fmt.Sprintf("%s_%s.png", prefix, hash[:10])
	return uniqueFilename
}

func generateHash(input string) string {
	// You can use your preferred hash function here
	// This example uses a simple hash to keep it short
	// You may want to use a more secure hash function in a real application
	hash := 0
	for _, char := range input {
		hash = (hash << 5) + int(char)
	}
	return fmt.Sprintf("%x", hash)
}
