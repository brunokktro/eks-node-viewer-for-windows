[![GitHub License](https://img.shields.io/badge/License-Apache%202.0-ff69b4.svg)](https://github.com/brunokktro/eks-node-viewer-for-windows/blob/main/LICENSE)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/brunokktro/eks-node-viewer-for-windows/issues)
[![Go code tests](https://github.com/brunokktro/eks-node-viewer-for-windows/actions/workflows/test.yaml/badge.svg)](https://github.com/brunokktro/eks-node-viewer-for-windows/actions/workflows/test.yaml)
[![Windows Tests](https://github.com/brunokktro/eks-node-viewer-for-windows/actions/workflows/test-windows.yaml/badge.svg)](https://github.com/brunokktro/eks-node-viewer-for-windows/actions/workflows/test-windows.yaml)

# EKS Node Viewer for Windows 🪟

This is a **Windows-optimized fork** of the original [AWS Labs EKS Node Viewer](https://github.com/awslabs/eks-node-viewer) with enhanced Windows support, PowerShell scripts, and comprehensive Windows documentation.

## 🆕 Windows Enhancements

- ✅ **PowerShell Build Scripts** - Native Windows build experience
- ✅ **Windows Terminal Support** - Optimized for modern Windows terminals
- ✅ **Comprehensive Windows Documentation** - Step-by-step Windows setup guide
- ✅ **Cross-Platform Configuration** - Smart platform detection and path handling
- ✅ **Windows CI/CD** - GitHub Actions for Windows testing
- ✅ **Enhanced Error Handling** - Windows-specific troubleshooting

## 🚀 Quick Start for Windows

### Option 1: PowerShell Build (Recommended)
```powershell
git clone https://github.com/brunokktro/eks-node-viewer-for-windows.git
cd eks-node-viewer-for-windows
.\build.ps1
```

### Option 2: Go Install
```cmd
go install github.com/brunokktro/eks-node-viewer-for-windows/cmd/eks-node-viewer@latest
```

### Option 3: Download Binary
Download the latest Windows binary from [Releases](https://github.com/brunokktro/eks-node-viewer-for-windows/releases)

📖 **For detailed Windows instructions, see [WINDOWS.md](WINDOWS.md)**

## About

`eks-node-viewer` is a tool for visualizing dynamic node usage within a cluster. It was originally developed as an internal tool at AWS for demonstrating consolidation with [Karpenter](https://karpenter.sh/). It displays the scheduled pod resource requests vs the allocatable capacity on the node. It *does not* look at the actual pod resource usage.

This **Windows-optimized fork** adds comprehensive Windows support while maintaining full compatibility with Linux and macOS.

![](./.static/screenshot.png)

## 🪟 Windows-Specific Features

- **PowerShell Integration**: Native PowerShell build and deployment scripts
- **Windows Terminal Optimization**: Enhanced color support and Unicode rendering
- **Smart Path Handling**: Automatic Windows path normalization and detection
- **Terminal Detection**: Automatic detection of Command Prompt, PowerShell, and Windows Terminal
- **Windows CI/CD**: Dedicated GitHub Actions for Windows testing and builds

### Talks Using eks-node-viewer

- [Containers from the Couch: Workload Consolidation with Karpenter](https://www.youtube.com/watch?v=BnksdJ3oOEs)
- [AWS re:Invent 2022 - Kubernetes virtually anywhere, for everyone](https://www.youtube.com/watch?v=OB7IZolZk78)

### Installation

#### Windows (Enhanced Support) 🪟

##### PowerShell Build Script (Recommended)
```powershell
git clone https://github.com/brunokktro/eks-node-viewer-for-windows.git
cd eks-node-viewer-for-windows
.\build.ps1 -Generate -Test
```

##### Pre-built Binary
1. Download the latest Windows binary from [GitHub Releases](https://github.com/brunokktro/eks-node-viewer-for-windows/releases)
2. Extract `eks-node-viewer.exe` to a directory in your PATH
3. Verify installation:
   ```cmd
   eks-node-viewer --version
   ```

##### Go Install
```cmd
go install github.com/brunokktro/eks-node-viewer-for-windows/cmd/eks-node-viewer@latest
```

**📖 For comprehensive Windows setup instructions, see [WINDOWS.md](WINDOWS.md)**

#### Homebrew (macOS/Linux)

```bash
# Note: Use original repository for Homebrew
brew tap aws/tap
brew install eks-node-viewer
```

#### Manual (All Platforms)
Please either fetch the latest [release](https://github.com/brunokktro/eks-node-viewer-for-windows/releases) or install manually using:
```shell
go install github.com/brunokktro/eks-node-viewer-for-windows/cmd/eks-node-viewer@latest
```

Note: This will install it to your `GOBIN` directory, typically `~/go/bin` on Unix systems or `%USERPROFILE%\go\bin` on Windows if it is unconfigured.

## Usage
```shell
Usage of ./eks-node-viewer:
  -attribution
    	Show the Open Source Attribution
  -context string
    	Name of the kubernetes context to use
  -disable-pricing
    	Disable pricing lookups
  -extra-labels string
    	A comma separated set of extra node labels to display
  -kubeconfig string
    	Absolute path to the kubeconfig file (default "~/.kube/config")
  -node-selector string
    	Node label selector used to filter nodes, if empty all nodes are selected
  -node-sort string
    	Sort order for the nodes, either 'creation' or a label name. The sort order can be controlled by appending =asc or =dsc to the value. (default "creation")
  -resources string
    	List of comma separated resources to monitor (default "cpu")
  -style string
    	Three color to use for styling 'good','ok' and 'bad' values. These are also used in the gradients displayed from bad -> good. (default "#04B575,#FFFF00,#FF0000")
  -v	Display eks-node-viewer version
  -version
    	Display eks-node-viewer version
```

### Examples
```shell
# Standard usage
eks-node-viewer

# Windows PowerShell with environment variables
$env:AWS_PROFILE = "myprofile"
$env:AWS_REGION = "us-west-2"
eks-node-viewer

# Karpenter nodes only
eks-node-viewer --node-selector karpenter.sh/nodepool

# Display both CPU and Memory Usage
eks-node-viewer --resources cpu,memory

# Display extra labels, i.e. AZ
eks-node-viewer --extra-labels topology.kubernetes.io/zone

# Sort by CPU usage in descending order
eks-node-viewer --node-sort=eks-node-viewer/node-cpu-usage=dsc

# Windows-specific: Disable pricing for faster startup
eks-node-viewer --disable-pricing
```

### Computed Labels

`eks-node-viewer` supports some custom label names that can be passed to the `--extra-labels` to display additional node information. 

- `eks-node-viewer/node-age` - Age of the node
- `eks-node-viewer/node-cpu-usage` - CPU usage (requests)
- `eks-node-viewer/node-memory-usage` - Memory usage (requests)
- `eks-node-viewer/node-pods-usage` - Pod usage (requests)
- `eks-node-viewer/node-ephemeral-storage-usage` - Ephemeral Storage usage (requests)

### Default Options
You can supply default options to `eks-node-viewer` by creating a file named `.eks-node-viewer` in your home directory and specifying
options there. The format is `option-name=value` where the option names are the command line flags:

**Unix/Linux/macOS:**
```text
# select only Karpenter managed nodes
node-selector=karpenter.sh/nodepool

# display both CPU and memory
resources=cpu,memory

# show the zone and nodepool name by default
extra-labels=topology.kubernetes.io/zone,karpenter.sh/nodepool

# sort so that the newest nodes are first
node-sort=creation=asc

# change default color style
style=#2E91D2,#ffff00,#D55E00
```

**Windows:**
The config file is located at `%USERPROFILE%\.eks-node-viewer` with the same format as above.

### Troubleshooting

#### NoCredentialProviders: no valid providers in chain. Deprecated.

This CLI relies on AWS credentials to access pricing data if you don't use the `--disable-pricing` option. You must have credentials configured via `~/aws/credentials`, `~/.aws/config`, environment variables, or some other credential provider chain.

See [credential provider documentation](https://docs.aws.amazon.com/sdk-for-go/api/aws/session/) for more.

#### I get an error of `creating client, exec plugin: invalid apiVersion "client.authentication.k8s.io/v1alpha1"`

Updating your AWS cli to the latest version and [updating your kubeconfig](https://docs.aws.amazon.com/cli/latest/reference/eks/update-kubeconfig.html) should resolve this issue.

## Development

### Building

**Unix/Linux/macOS:**
```shell
$ make build
```

**Windows (PowerShell):**
```powershell
.\build.ps1
```

**Cross-platform with GoReleaser:**
```shell
$ make goreleaser
```

### Windows Development

```powershell
# Full development build with tests
.\build.ps1 -Generate -Test

# Clean build
.\build.ps1 -Clean

# Release build
.\build.ps1 -Release
```

## 🤝 Contributing

This fork welcomes contributions, especially those that improve Windows compatibility and user experience. 

### Areas for Contribution:
- Windows-specific bug fixes
- PowerShell script improvements
- Windows Terminal enhancements
- Documentation improvements
- Cross-platform testing

## 📄 License

This project maintains the same Apache 2.0 license as the original AWS Labs project.

## 🙏 Acknowledgments

- Original project by [AWS Labs](https://github.com/awslabs/eks-node-viewer)
- Windows enhancements by [brunokktro](https://github.com/brunokktro)
- All contributors to the original project

## 🔗 Related Links

- [Original EKS Node Viewer](https://github.com/awslabs/eks-node-viewer)
- [Karpenter](https://karpenter.sh/)
- [Windows Documentation](WINDOWS.md)
