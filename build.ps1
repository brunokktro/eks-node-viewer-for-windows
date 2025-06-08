#!/usr/bin/env pwsh
# Windows build script for EKS Node Viewer

param(
    [string]$Version = "local",
    [string]$BuiltBy = "PowerShell",
    [switch]$Release,
    [switch]$Clean,
    [switch]$Test,
    [switch]$Generate
)

$ErrorActionPreference = "Stop"

Write-Host "EKS Node Viewer Windows Build Script" -ForegroundColor Cyan
Write-Host "====================================" -ForegroundColor Cyan

if ($Clean) {
    Write-Host "Cleaning artifacts..." -ForegroundColor Yellow
    if (Test-Path "eks-node-viewer.exe") { Remove-Item "eks-node-viewer.exe" -Force }
    if (Test-Path "dist") { Remove-Item "dist" -Recurse -Force }
    Write-Host "Clean completed!" -ForegroundColor Green
}

if ($Generate) {
    Write-Host "Generating code..." -ForegroundColor Yellow
    
    # Generate attribution (run twice as per Makefile)
    go generate ./...
    & "./hack/gen_licenses.ps1"
    go generate ./...
    
    # Download pricing data
    Write-Host "Downloading pricing data..." -ForegroundColor Yellow
    $urls = @(
        "https://raw.githubusercontent.com/aws/karpenter-provider-aws/main/pkg/providers/pricing/zz_generated.pricing_aws.go",
        "https://raw.githubusercontent.com/aws/karpenter-provider-aws/main/pkg/providers/pricing/zz_generated.pricing_aws_cn.go",
        "https://raw.githubusercontent.com/aws/karpenter-provider-aws/main/pkg/providers/pricing/zz_generated.pricing_aws_us_gov.go"
    )
    
    $files = @(
        "./pkg/aws/zz_generated_aws.pricing.go",
        "./pkg/aws/zz_generated_aws_cn.pricing.go",
        "./pkg/aws/zz_generated_aws_us_gov.pricing.go"
    )
    
    for ($i = 0; $i -lt $urls.Length; $i++) {
        Invoke-WebRequest -Uri $urls[$i] -OutFile $files[$i]
        (Get-Content $files[$i]) -replace 'package pricing', 'package aws' | Set-Content $files[$i]
    }
    
    Write-Host "Code generation completed!" -ForegroundColor Green
}

if ($Test) {
    Write-Host "Running tests..." -ForegroundColor Yellow
    go test -v ./pkg/... ./cmd/...
    Write-Host "Tests completed!" -ForegroundColor Green
}

Write-Host "Building EKS Node Viewer..." -ForegroundColor Yellow

$ldflags = "-s -w -X main.version=$Version -X main.builtBy=$BuiltBy"

if ($Release) {
    Write-Host "Building release with GoReleaser..." -ForegroundColor Yellow
    goreleaser build --snapshot --clean
} else {
    Write-Host "Building local binary..." -ForegroundColor Yellow
    $env:CGO_ENABLED = "0"
    go build -ldflags="$ldflags" -o eks-node-viewer.exe ./cmd/eks-node-viewer
}

if ($LASTEXITCODE -eq 0) {
    Write-Host "Build completed successfully!" -ForegroundColor Green
    if (Test-Path "eks-node-viewer.exe") {
        $size = (Get-Item "eks-node-viewer.exe").Length / 1MB
        Write-Host "Binary size: $([math]::Round($size, 2)) MB" -ForegroundColor Cyan
    }
} else {
    Write-Host "Build failed!" -ForegroundColor Red
    exit $LASTEXITCODE
}
