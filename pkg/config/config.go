/*
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
package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"k8s.io/client-go/util/homedir"
)

// GetDefaultKubeconfigPath returns the default kubeconfig path for the current platform
func GetDefaultKubeconfigPath() string {
	homeDir := homedir.HomeDir()
	return filepath.Join(homeDir, ".kube", "config")
}

// GetConfigFilePath returns the path to the eks-node-viewer config file
func GetConfigFilePath() string {
	homeDir := homedir.HomeDir()
	return filepath.Join(homeDir, ".eks-node-viewer")
}

// GetExecutableName returns the appropriate executable name for the current platform
func GetExecutableName() string {
	if runtime.GOOS == "windows" {
		return "eks-node-viewer.exe"
	}
	return "eks-node-viewer"
}

// NormalizePath normalizes file paths for cross-platform compatibility
func NormalizePath(path string) string {
	// Convert forward slashes to backslashes on Windows
	if runtime.GOOS == "windows" {
		path = strings.ReplaceAll(path, "/", string(filepath.Separator))
	}
	
	// Clean the path to remove redundant separators
	return filepath.Clean(path)
}

// GetTempDir returns the appropriate temporary directory for the current platform
func GetTempDir() string {
	if runtime.GOOS == "windows" {
		if temp := os.Getenv("TEMP"); temp != "" {
			return temp
		}
		if temp := os.Getenv("TMP"); temp != "" {
			return temp
		}
		return "C:\\Windows\\Temp"
	}
	
	if temp := os.Getenv("TMPDIR"); temp != "" {
		return temp
	}
	return "/tmp"
}

// IsWindowsTerminal checks if running in Windows Terminal or PowerShell
func IsWindowsTerminal() bool {
	if runtime.GOOS != "windows" {
		return false
	}
	
	// Check for Windows Terminal
	if os.Getenv("WT_SESSION") != "" {
		return true
	}
	
	// Check for PowerShell
	if strings.Contains(strings.ToLower(os.Getenv("PSModulePath")), "powershell") {
		return true
	}
	
	return false
}

// GetShellInfo returns information about the current shell environment
func GetShellInfo() map[string]string {
	info := make(map[string]string)
	info["os"] = runtime.GOOS
	info["arch"] = runtime.GOARCH
	
	if runtime.GOOS == "windows" {
		info["shell"] = "cmd"
		if os.Getenv("PSModulePath") != "" {
			info["shell"] = "powershell"
		}
		if os.Getenv("WT_SESSION") != "" {
			info["terminal"] = "windows-terminal"
		}
	} else {
		if shell := os.Getenv("SHELL"); shell != "" {
			info["shell"] = filepath.Base(shell)
		}
		if term := os.Getenv("TERM"); term != "" {
			info["terminal"] = term
		}
	}
	
	return info
}
