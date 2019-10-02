package walker

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindMusicFiles(dir string) []string {

	listOfFilePath := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && isKnownExtension(info.Name()) {
			// fmt.Printf("visited file or dir: %q\n", path)
			listOfFilePath = append(listOfFilePath, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path  %v\n", err)
		return []string{}
	}

	fmt.Printf("total files: %d\n", len(listOfFilePath))

	return listOfFilePath
}

func isKnownExtension(fileName string) bool {

	audioFiles := []string{
		".mp3",
		".MP3",
		".wav",
		".wma",
	}

	for _, a := range audioFiles {
		if a == filepath.Ext(fileName) {
			return true
		}
	}
	return false

}
