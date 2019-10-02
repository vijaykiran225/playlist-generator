package metadata

import (
	"os"

	"github.com/dhowden/tag"
)

func ExtractMetaData(filePath string) (tag.Metadata, error) {

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	metaData, err := tag.ReadFrom(f)
	if err != nil {
		return nil, err
	}

	return metaData, nil
}
