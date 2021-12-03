package play

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	writeToEnvFile("12345")
	pageID, _ := getMetadataPageID()
	fmt.Println(pageID)
}

func writeToEnvFile(metadataPageID string) error {
	content := fmt.Sprintf(`CONFLUENCE_METADATA_PAGE_ID=%s`, metadataPageID)
	file, err := os.Create("newFolder/test.env")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content + "\n")
	if err != nil {
		return err
	}
	file.Sync()
	return err
}

func getMetadataPageID() (string, error) {
	contents, err := os.ReadFile("newFolder/test.env")
	if err != nil {
		return "", err
	}
	splitContents := strings.Split(string(contents), "=")
	return splitContents[1], nil
}
