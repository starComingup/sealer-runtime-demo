package kubernets

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// LoadMetadata :read metadata via ClusterImage name.
func LoadMetadata(rootfs string) (*Metadata, error) {
	metadataPath := filepath.Join(rootfs, "Metadata")
	var metadataFile []byte
	var err error
	var md Metadata

	if !IsFileExist(metadataPath) {
		return nil, nil
	}

	metadataFile, err = os.ReadFile(filepath.Clean(metadataPath))

	if err != nil {
		return nil, fmt.Errorf("failed to read ClusterImage metadata %v", err)
	}
	err = json.Unmarshal(metadataFile, &md)
	if err != nil {
		return nil, fmt.Errorf("failed to load ClusterImage metadata %v", err)
	}
	return &md, nil
}

func IsFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}
