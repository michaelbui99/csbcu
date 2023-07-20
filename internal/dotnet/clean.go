package dotnet

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/michaelbui99/csbcu/internal/ioutil"
)

func CleanBinaries(projectPaths []string) (int, error) {
	projectsCleaned := 0
	for _, projectPath := range projectPaths {
		cleaned, err := CleanBinary(projectPath)
		if err != nil {
			return projectsCleaned, err
		}

		if cleaned {
			projectsCleaned++
		}
	}

	return projectsCleaned, nil
}

func CleanBinary(projectPath string) (bool, error) {
	binFolders := [2]string{"obj", "bin"}

	for _, binFolder := range binFolders {

		absPath, err := filepath.Abs(projectPath)
		if err != nil {
			return false, err
		}

		csprojPath := fmt.Sprintf("%v/%v", projectPath, path.Base(projectPath)+".csproj")
		if !ioutil.FileExists(csprojPath) {
			return false, nil
		}

		err = os.RemoveAll(fmt.Sprintf("%v/%v", absPath, binFolder))
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
