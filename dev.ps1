# Development script with hot reload
# Run: .\dev.ps1

Write-Host "🔥 Job Application Tracker - Development Mode" -ForegroundColor Magenta
Write-Host "============================================`n" -ForegroundColor Magenta

# Check if Air is installed
$airInstalled = Get-Command air -ErrorAction SilentlyContinue
if (-not $airInstalled) {
    Write-Host "⚠️  Air (hot reload tool) not found. Installing..." -ForegroundColor Yellow
    go install github.com/cosmtrek/air@latest
    if ($LASTEXITCODE -ne 0) {
        Write-Host "❌ Failed to install Air" -ForegroundColor Red
        Write-Host "Running without hot reload instead..." -ForegroundColor Yellow
        go run main.go
        exit
    }
    Write-Host "✅ Air installed successfully" -ForegroundColor Green
}

# Check if .air.toml exists
if (-not (Test-Path ".air.toml")) {
    Write-Host "❌ .air.toml not found" -ForegroundColor Red
    exit 1
}

# Run with Air
Write-Host "🚀 Starting development server with hot reload..." -ForegroundColor Yellow
Write-Host "Application: http://localhost:3000" -ForegroundColor Cyan
Write-Host "Changes will auto-reload. Press Ctrl+C to stop`n" -ForegroundColor Gray

air
