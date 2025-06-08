#!/usr/bin/env pwsh
# PowerShell equivalent of gen_licenses.sh for Windows compatibility

Write-Host "Installing go-licenses..." -ForegroundColor Green
go install github.com/google/go-licenses@latest

Write-Host "Downloading dependencies..." -ForegroundColor Green
go mod download

Write-Host "Generating attribution..." -ForegroundColor Green
$env:GOROOT = go env GOROOT
go-licenses report ./... --template hack/attribution.tmpl | Out-File -FilePath ATTRIBUTION.md -Encoding UTF8

Write-Host "Attribution generated successfully!" -ForegroundColor Green
