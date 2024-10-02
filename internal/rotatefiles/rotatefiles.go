package rotatefiles

import (
	"fmt"
	"os"
	"path/filepath"
)

func getFilesInfo(path string, d os.DirEntry, err error) error {
	// ignore dirs
	if d.IsDir() {
		return nil
	}

	// fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\tModTime: %s\n", d.Name(), d.Size(), path, d.ModTime())
	fmt.Println(os.Stat(path))

	return nil
}

// Rotate files: keep <num> of most recent files and delete other
func RotateFilesByMtime(filesDir string, filesToKeep int) error {
	if err := filepath.WalkDir(filesDir, getFilesInfo); err != nil {
		return err
	}

	return nil
}
