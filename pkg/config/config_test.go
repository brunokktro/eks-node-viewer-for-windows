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
	"testing"
)

func TestGetDefaultKubeconfigPath(t *testing.T) {
	path := GetDefaultKubeconfigPath()
	
	// Should contain .kube/config
	if !strings.Contains(path, ".kube") {
		t.Errorf("Expected path to contain .kube, got: %s", path)
	}
	
	// Should be absolute path
	if !filepath.IsAbs(path) {
		t.Errorf("Expected absolute path, got: %s", path)
	}
	
	// Should use correct separator for platform
	expectedSep := string(filepath.Separator)
	if !strings.Contains(path, expectedSep) && len(path) > 1 {
		t.Errorf("Expected path to use platform separator %s, got: %s", expectedSep, path)
	}
}

func TestGetConfigFilePath(t *testing.T) {
	path := GetConfigFilePath()
	
	// Should contain .eks-node-viewer
	if !strings.Contains(path, ".eks-node-viewer") {
		t.Errorf("Expected path to contain .eks-node-viewer, got: %s", path)
	}
	
	// Should be absolute path
	if !filepath.IsAbs(path) {
		t.Errorf("Expected absolute path, got: %s", path)
	}
}

func TestGetExecutableName(t *testing.T) {
	name := GetExecutableName()
	
	if runtime.GOOS == "windows" {
		if name != "eks-node-viewer.exe" {
			t.Errorf("Expected eks-node-viewer.exe on Windows, got: %s", name)
		}
	} else {
		if name != "eks-node-viewer" {
			t.Errorf("Expected eks-node-viewer on non-Windows, got: %s", name)
		}
	}
}

func TestNormalizePath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"path/to/file", filepath.Join("path", "to", "file")},
		{"path\\to\\file", filepath.Join("path", "to", "file")},
		{"./path/../file", "file"},
		{"/absolute/path", filepath.Join(string(filepath.Separator), "absolute", "path")},
	}
	
	for _, test := range tests {
		result := NormalizePath(test.input)
		if result != test.expected {
			t.Errorf("NormalizePath(%s) = %s, expected %s", test.input, result, test.expected)
		}
	}
}

func TestGetTempDir(t *testing.T) {
	tempDir := GetTempDir()
	
	// Should return a non-empty string
	if tempDir == "" {
		t.Error("Expected non-empty temp directory")
	}
	
	// Should be absolute path
	if !filepath.IsAbs(tempDir) {
		t.Errorf("Expected absolute path for temp directory, got: %s", tempDir)
	}
	
	// Platform-specific checks
	if runtime.GOOS == "windows" {
		// Should be a Windows path
		if !strings.Contains(tempDir, ":") && !strings.HasPrefix(tempDir, "\\\\") {
			t.Errorf("Expected Windows path format, got: %s", tempDir)
		}
	} else {
		// Should be Unix path
		if strings.Contains(tempDir, ":") {
			t.Errorf("Expected Unix path format, got: %s", tempDir)
		}
	}
}

func TestIsWindowsTerminal(t *testing.T) {
	// Save original environment
	originalWT := os.Getenv("WT_SESSION")
	originalPS := os.Getenv("PSModulePath")
	
	defer func() {
		// Restore original environment
		os.Setenv("WT_SESSION", originalWT)
		os.Setenv("PSModulePath", originalPS)
	}()
	
	if runtime.GOOS == "windows" {
		// Test Windows Terminal detection
		os.Setenv("WT_SESSION", "test-session")
		os.Setenv("PSModulePath", "")
		if !IsWindowsTerminal() {
			t.Error("Expected IsWindowsTerminal() to return true with WT_SESSION set")
		}
		
		// Test PowerShell detection
		os.Setenv("WT_SESSION", "")
		os.Setenv("PSModulePath", "C:\\Program Files\\PowerShell\\Modules")
		if !IsWindowsTerminal() {
			t.Error("Expected IsWindowsTerminal() to return true with PSModulePath set")
		}
		
		// Test no detection
		os.Setenv("WT_SESSION", "")
		os.Setenv("PSModulePath", "")
		if IsWindowsTerminal() {
			t.Error("Expected IsWindowsTerminal() to return false with no indicators")
		}
	} else {
		// Should always return false on non-Windows
		if IsWindowsTerminal() {
			t.Error("Expected IsWindowsTerminal() to return false on non-Windows platforms")
		}
	}
}

func TestGetShellInfo(t *testing.T) {
	info := GetShellInfo()
	
	// Should contain basic info
	if info["os"] == "" {
		t.Error("Expected os information")
	}
	
	if info["arch"] == "" {
		t.Error("Expected arch information")
	}
	
	// OS should match runtime
	if info["os"] != runtime.GOOS {
		t.Errorf("Expected os to be %s, got %s", runtime.GOOS, info["os"])
	}
	
	// Arch should match runtime
	if info["arch"] != runtime.GOARCH {
		t.Errorf("Expected arch to be %s, got %s", runtime.GOARCH, info["arch"])
	}
	
	// Platform-specific checks
	if runtime.GOOS == "windows" {
		if info["shell"] == "" {
			t.Error("Expected shell information on Windows")
		}
	}
}
