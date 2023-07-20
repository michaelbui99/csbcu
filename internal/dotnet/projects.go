package dotnet

import (
	"io/fs"
	"path/filepath"
	"strings"
)

const projectFileExtension = "csproj"

func GetProjectPaths(searchDir string) ([]string, error) {
	projectPaths := make([]string, 0)

	err := filepath.Walk(searchDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() && (info.Name() == "obj" || info.Name() == "bin") {
			return filepath.SkipDir
		}

		if strings.Contains(info.Name(), projectFileExtension) {
			projectPaths = append(projectPaths, filepath.Dir(path))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return projectPaths, nil
}
