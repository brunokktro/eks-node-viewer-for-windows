# Changelog - EKS Node Viewer for Windows

All notable changes to this Windows-optimized fork will be documented in this file.

## [1.0.0] - 2025-06-08

### 🆕 Added - Windows Enhancements

#### PowerShell Integration
- **build.ps1**: Comprehensive PowerShell build script with multiple options
  - `-Generate`: Code generation and dependency download
  - `-Test`: Run all tests
  - `-Clean`: Clean build artifacts
  - `-Release`: Build with GoReleaser
- **hack/gen_licenses.ps1**: PowerShell version of license generation script

#### Windows-Specific Configuration
- **pkg/config/config.go**: New configuration package with Windows-specific features
  - Smart platform detection (Windows Terminal, PowerShell, Command Prompt)
  - Cross-platform path normalization
  - Windows-specific temporary directory handling
  - Shell and terminal environment detection

#### Documentation
- **WINDOWS.md**: Comprehensive Windows installation and usage guide
  - Multiple installation methods
  - Terminal-specific instructions
  - Troubleshooting section
  - Performance optimization tips
- **README.md**: Updated with Windows-specific sections and examples

#### Testing & CI/CD
- **pkg/config/config_test.go**: Unit tests for Windows-specific functionality
- **.github/workflows/test-windows.yaml**: Windows-specific GitHub Actions workflow
  - Windows Server testing
  - Cross-platform build validation
  - PowerShell script testing

### 🔧 Enhanced - Cross-Platform Compatibility

#### Path Handling
- Improved file path handling for Windows vs Unix systems
- Automatic path separator normalization
- Enhanced home directory detection

#### Terminal Support
- Windows Terminal color optimization
- PowerShell Unicode support
- Command Prompt compatibility
- Automatic terminal capability detection

#### Build System
- GoReleaser already supported Windows (maintained)
- Added PowerShell alternative to Makefile
- Enhanced cross-platform build scripts

### 📖 Documentation Improvements

#### Windows-Specific Guides
- Step-by-step Windows installation
- Terminal comparison and recommendations
- PowerShell vs Command Prompt usage
- Windows-specific troubleshooting

#### Enhanced Examples
- PowerShell environment variable syntax
- Windows path examples
- Terminal-specific command variations

### 🧪 Testing Enhancements

#### Windows Testing
- Dedicated Windows CI/CD pipeline
- Cross-platform build validation
- PowerShell script testing
- Windows-specific unit tests

#### Platform Coverage
- Windows Server 2019/2022
- Windows 10/11
- Multiple Go versions
- Cross-compilation testing

### 🔄 Maintained Compatibility

#### Original Features
- All original functionality preserved
- Linux and macOS compatibility maintained
- Same command-line interface
- Same configuration file format

#### Dependencies
- No new external dependencies added
- All Windows enhancements use Go standard library
- Maintained compatibility with existing AWS and Kubernetes integrations

## [Upstream] - Based on AWS Labs EKS Node Viewer

This fork is based on the original [AWS Labs EKS Node Viewer](https://github.com/awslabs/eks-node-viewer) and maintains compatibility with all upstream features while adding Windows-specific enhancements.

### Original Features Maintained
- EKS cluster node visualization
- Karpenter integration
- AWS pricing integration
- Kubernetes resource monitoring
- Terminal UI with Bubble Tea
- Cross-platform Go build support

---

## Contributing

We welcome contributions that improve Windows compatibility and user experience. Please see the main README.md for contribution guidelines.

## License

This project maintains the same Apache 2.0 license as the original AWS Labs project.
