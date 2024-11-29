/*
Copyright © 2024 Pavel Kadera info@pavelkadera.cz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package install

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func UpdateUserPath(newPath string) error {
	if newPath == "" {
		return fmt.Errorf("new path is empty")
	}

	// Open the registry key where the user's PATH is stored
	key, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		return fmt.Errorf("failed to open registry key: %w", err)
	}
	defer key.Close()

	// Get the current PATH value
	currentPath, _, err := key.GetStringValue("Path")
	if err != nil && err != registry.ErrNotExist {
		return fmt.Errorf("failed to read current PATH: %w", err)
	}

	// Check if the new path already exists
	if containsPath(currentPath, newPath) {
		return fmt.Errorf("403 new path already exists in PATH")
	}
	finPath := currentPath + ";" + newPath
	// Update the PATH variable
	err = key.SetStringValue("Path", finPath)
	if err != nil {
		return fmt.Errorf("failed to set new PATH: %w", err)
	}

	return nil
}

func containsPath(currentPath, newPath string) bool {
	paths := splitEnvPath(currentPath)
	for _, p := range paths {
		if p == newPath {
			return true
		}
	}
	return false
}

func splitEnvPath(envPath string) []string {
	return strings.Split(envPath, ";")
}

func GetCurrentPwd() (string, error) {
	// Get the absolute path of the executable
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}

	// Get the directory of the executable
	workingDir := filepath.Dir(execPath)

	return workingDir, nil
}
