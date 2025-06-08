# Seção para adicionar ao README.md principal

## Installation

### Homebrew (macOS/Linux)

```bash
brew tap aws/tap
brew install eks-node-viewer
```

### Windows

#### Pre-built Binary (Recommended)
1. Download the latest Windows binary from [GitHub Releases](https://github.com/awslabs/eks-node-viewer/releases)
2. Extract `eks-node-viewer.exe` to a directory in your PATH
3. Verify installation:
   ```cmd
   eks-node-viewer --version
   ```

#### PowerShell Build Script
```powershell
git clone https://github.com/awslabs/eks-node-viewer.git
cd eks-node-viewer
.\build.ps1
```

#### Go Install
```cmd
go install github.com/awslabs/eks-node-viewer/cmd/eks-node-viewer@latest
```

**Note for Windows users**: See [WINDOWS.md](WINDOWS.md) for detailed Windows-specific installation and usage instructions.

### Manual (All Platforms)
Please either fetch the latest [release](https://github.com/awslabs/eks-node-viewer/releases) or install manually using:
```shell
go install github.com/awslabs/eks-node-viewer/cmd/eks-node-viewer@latest
```

Note: This will install it to your `GOBIN` directory, typically `~/go/bin` on Unix systems or `%USERPROFILE%\go\bin` on Windows if it is unconfigured.
