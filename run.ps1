# PowerShell script untuk menjalankan Job Application Tracker
# Run: .\run.ps1

Write-Host "🚀 Job Application Tracker - Startup Script" -ForegroundColor Cyan
Write-Host "==========================================`n" -ForegroundColor Cyan

# Check if Go is installed
Write-Host "Checking prerequisites..." -ForegroundColor Yellow
$goVersion = go version 2>$null
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Go tidak terinstall. Download dari: https://go.dev/dl/" -ForegroundColor Red
    exit 1
}
Write-Host "✅ Go version: $goVersion" -ForegroundColor Green

# Check if .env file exists
if (-not (Test-Path ".env")) {
    Write-Host "⚠️  .env file tidak ditemukan. Membuat dari .env.example..." -ForegroundColor Yellow
    Copy-Item ".env.example" ".env"
    Write-Host "✅ .env file created. Silakan edit file .env dengan database credentials Anda." -ForegroundColor Green
    Write-Host "Press any key to continue..."
    $null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
}

# Install dependencies
Write-Host "`n📦 Installing dependencies..." -ForegroundColor Yellow
go mod download
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Failed to download dependencies" -ForegroundColor Red
    exit 1
}
Write-Host "✅ Dependencies installed" -ForegroundColor Green

# Build application
Write-Host "`n🔨 Building application..." -ForegroundColor Yellow
go build -o main.exe main.go
if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Build failed" -ForegroundColor Red
    exit 1
}
Write-Host "✅ Build successful" -ForegroundColor Green

# Run application
Write-Host "`n🎯 Starting server..." -ForegroundColor Yellow
Write-Host "Application will be available at: http://localhost:3000" -ForegroundColor Cyan
Write-Host "Press Ctrl+C to stop the server`n" -ForegroundColor Gray

.\main.exe
