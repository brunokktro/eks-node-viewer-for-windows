# EKS Node Viewer - Windows Installation and Usage Guide

This guide provides Windows-specific instructions for installing and using EKS Node Viewer.

## Prerequisites

### Required Software
- **Go 1.24.2 or later** - [Download from golang.org](https://golang.org/downloads/)
- **Git** - [Download from git-scm.com](https://git-scm.com/downloads)
- **AWS CLI v2** - [Installation Guide](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- **kubectl** - [Installation Guide](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows/)

### Optional Software
- **Windows Terminal** - Recommended for better color support and Unicode rendering
- **PowerShell 7+** - Enhanced scripting capabilities
- **GoReleaser** - For building releases

## Installation Methods

### Method 1: Pre-built Binary (Recommended)

1. Download the latest Windows binary from [GitHub Releases](https://github.com/awslabs/eks-node-viewer/releases)
2. Extract `eks-node-viewer.exe` to a directory in your PATH
3. Open Command Prompt or PowerShell and verify installation:
   ```cmd
   eks-node-viewer --version
   ```

### Method 2: Build from Source

#### Using PowerShell (Recommended)
```powershell
# Clone the repository
git clone https://github.com/awslabs/eks-node-viewer.git
cd eks-node-viewer

# Build using the provided PowerShell script
.\build.ps1 -Generate -Test

# Or build without tests
.\build.ps1
```

#### Using Command Prompt
```cmd
# Clone the repository
git clone https://github.com/awslabs/eks-node-viewer.git
cd eks-node-viewer

# Build manually
set CGO_ENABLED=0
go build -ldflags="-s -w -X main.version=local -X main.builtBy=Windows" -o eks-node-viewer.exe ./cmd/eks-node-viewer
```

#### Using Go Install
```cmd
go install github.com/awslabs/eks-node-viewer/cmd/eks-node-viewer@latest
```

## Configuration

### AWS Credentials
Ensure your AWS credentials are configured. EKS Node Viewer supports all standard AWS credential methods:

```cmd
# Using AWS CLI
aws configure

# Or set environment variables
set AWS_ACCESS_KEY_ID=your-access-key
set AWS_SECRET_ACCESS_KEY=your-secret-key
set AWS_REGION=us-west-2
```

### Kubernetes Configuration
Ensure kubectl is configured to access your EKS cluster:

```cmd
# Update kubeconfig for your EKS cluster
aws eks update-kubeconfig --region us-west-2 --name your-cluster-name

# Verify connection
kubectl get nodes
```

### Configuration File
Create a configuration file at `%USERPROFILE%\.eks-node-viewer`:

```ini
# Example Windows configuration
node-selector=karpenter.sh/nodepool
resources=cpu,memory
extra-labels=topology.kubernetes.io/zone,karpenter.sh/nodepool
node-sort=creation=asc
style=#2E91D2,#ffff00,#D55E00
```

## Usage

### Basic Usage
```cmd
# Standard usage
eks-node-viewer.exe

# With specific options
eks-node-viewer.exe --resources cpu,memory --extra-labels topology.kubernetes.io/zone
```

### PowerShell Usage
```powershell
# Standard usage
.\eks-node-viewer.exe

# With AWS profile and region
$env:AWS_PROFILE = "myprofile"
$env:AWS_REGION = "us-west-2"
.\eks-node-viewer.exe
```

### Advanced Examples
```cmd
# Karpenter nodes only
eks-node-viewer.exe --node-selector karpenter.sh/nodepool

# Display both CPU and Memory Usage
eks-node-viewer.exe --resources cpu,memory

# Display extra labels (AZ)
eks-node-viewer.exe --extra-labels topology.kubernetes.io/zone

# Sort by CPU usage in descending order
eks-node-viewer.exe --node-sort=eks-node-viewer/node-cpu-usage=dsc

# Disable pricing (faster startup)
eks-node-viewer.exe --disable-pricing
```

## Windows-Specific Features

### Terminal Compatibility
- **Command Prompt**: Basic functionality, limited color support
- **PowerShell**: Enhanced color support and Unicode characters
- **Windows Terminal**: Full color support, best experience

### Color Support
EKS Node Viewer automatically detects your terminal capabilities:
- Windows Terminal: Full 24-bit color support
- PowerShell: 16-color support
- Command Prompt: Basic color support

### File Paths
All file paths use Windows conventions:
- Config file: `%USERPROFILE%\.eks-node-viewer`
- Kubeconfig: `%USERPROFILE%\.kube\config`
- Temp files: `%TEMP%\eks-node-viewer-*`

## Troubleshooting

### Common Issues

#### "eks-node-viewer is not recognized as an internal or external command"
- Ensure the executable is in your PATH
- Use the full path to the executable: `C:\path\to\eks-node-viewer.exe`

#### Color/Unicode Issues
- Use Windows Terminal for the best experience
- Update to PowerShell 7+ for better Unicode support
- Set console code page: `chcp 65001`

#### AWS Credential Issues
```cmd
# Verify AWS credentials
aws sts get-caller-identity

# Check environment variables
echo %AWS_PROFILE%
echo %AWS_REGION%
```

#### Kubernetes Connection Issues
```cmd
# Verify kubectl configuration
kubectl config current-context
kubectl get nodes

# Update kubeconfig if needed
aws eks update-kubeconfig --region your-region --name your-cluster
```

### Performance Tips
- Use `--disable-pricing` for faster startup
- Use `--node-selector` to filter nodes and improve performance
- Close other resource-intensive applications for better terminal rendering

## Development on Windows

### Building from Source
```powershell
# Install dependencies
go mod download

# Generate code and run tests
.\build.ps1 -Generate -Test

# Build release
.\build.ps1 -Release
```

### Testing
```powershell
# Run all tests
go test ./pkg/... ./cmd/...

# Run tests with coverage
go test -coverprofile=coverage.out ./pkg/... ./cmd/...
go tool cover -html=coverage.out
```

## Support

For Windows-specific issues:
1. Check this documentation first
2. Verify your terminal supports the required features
3. Test with different terminals (CMD, PowerShell, Windows Terminal)
4. Open an issue on [GitHub](https://github.com/awslabs/eks-node-viewer/issues) with:
   - Windows version
   - Terminal type
   - PowerShell/Go versions
   - Complete error messages
