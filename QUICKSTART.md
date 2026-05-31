# 🚀 Quick Start Guide

## Local Development Setup

### 1. Install Prerequisites

**Windows:**
```powershell
# Install Go 1.22+
# Download dari: https://go.dev/dl/

# Install PostgreSQL
# Download dari: https://www.postgresql.org/download/windows/

# Install Git
# Download dari: https://git-scm.com/download/win

# Install Make (optional)
choco install make
```

**Mac:**
```bash
brew install go postgresql git
```

**Linux:**
```bash
sudo apt install golang-go postgresql git
```

---

### 2. Clone & Setup

```bash
# Clone repository
git clone <your-repo-url>
cd jobs-me

# Install dependencies
go mod download

# Copy environment file
cp .env.example .env
```

---

### 3. Setup Database

**Opsi A: PostgreSQL Lokal**

```bash
# Windows (PowerShell dengan PostgreSQL di PATH)
psql -U postgres
```

```sql
-- Buat database
CREATE DATABASE jobsme;

-- Buat user (optional)
CREATE USER jobsme_user WITH PASSWORD 'yourpassword';
GRANT ALL PRIVILEGES ON DATABASE jobsme TO jobsme_user;

-- Exit
\q
```

**Update `.env`:**
```env
DATABASE_URL=postgresql://postgres:yourpassword@localhost:5432/jobsme?sslmode=disable
```

**Opsi B: Docker PostgreSQL**

```bash
docker run -d \
  --name postgres-jobsme \
  -e POSTGRES_DB=jobsme \
  -e POSTGRES_PASSWORD=password \
  -p 5432:5432 \
  postgres:16-alpine
```

**Update `.env`:**
```env
DATABASE_URL=postgresql://postgres:password@localhost:5432/jobsme?sslmode=disable
```

---

### 4. Run Application

**Cara 1: Direct Run**
```bash
go run main.go
```

**Cara 2: Build & Run**
```bash
go build -o main.exe
./main.exe
```

**Cara 3: Dengan Make**
```bash
make run
```

**Cara 4: Hot Reload (Development)**
```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run dengan hot reload
air
# atau
make dev
```

---

### 5. Akses Aplikasi

Buka browser dan akses:
```
http://localhost:3000
```

---

## 🎯 Fitur yang Tersedia

### Dashboard
- Lihat semua lamaran
- Statistik real-time
- Filter berdasarkan status dan platform

### Tambah Lamaran
1. Click "Tambah Lamaran"
2. Isi form:
   - Nama Perusahaan
   - Posisi
   - Platform (LinkedIn, Glints, Jobstreet, dll)
   - Status
   - Link lowongan
   - Catatan
3. Save

### Edit Lamaran
- Click icon edit (✏️) pada baris lamaran
- Update informasi
- Save

### Hapus Lamaran
- Click icon delete (🗑️)
- Confirm

### Filter & Search
- Cari perusahaan atau posisi
- Filter by status
- Filter by platform

---

## 📁 Project Structure

```
jobs-me/
├── config/           # Database configuration
├── handlers/         # HTTP request handlers
├── models/           # Database models (GORM)
├── routes/           # Route definitions
├── static/           # JavaScript, CSS, images
├── templates/        # HTML templates
├── main.go           # Application entry point
├── .env              # Environment variables (don't commit!)
├── go.mod            # Go dependencies
└── README.md         # Documentation
```

---

## 🧪 Testing

```bash
# Run tests
go test ./...

# With coverage
go test -cover ./...

# Verbose
go test -v ./...
```

---

## 🔧 Development Tools

### Recommended VS Code Extensions
- Go (by Go Team at Google)
- PostgreSQL (by Chris Kolkman)
- Thunder Client (for API testing)
- Error Lens

### Useful Commands

```bash
# Format code
go fmt ./...

# Check for errors
go vet ./...

# Update dependencies
go get -u ./...
go mod tidy

# Build for production
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```

---

## 📡 API Endpoints

Test dengan curl atau Thunder Client:

### Get All Jobs
```bash
curl http://localhost:3000/api/jobs
```

### Create Job
```bash
curl -X POST http://localhost:3000/api/jobs \
  -H "Content-Type: application/json" \
  -d '{
    "company": "PT Tech Indonesia",
    "position": "Backend Developer",
    "platform": "LinkedIn",
    "status": "Applied",
    "job_link": "https://linkedin.com/jobs/123",
    "notes": "Interview next week"
  }'
```

### Update Job
```bash
curl -X PUT http://localhost:3000/api/jobs/1 \
  -H "Content-Type: application/json" \
  -d '{
    "status": "Interview HRD"
  }'
```

### Delete Job
```bash
curl -X DELETE http://localhost:3000/api/jobs/1
```

---

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Windows: Find and kill process
netstat -ano | findstr :3000
taskkill /PID <process_id> /F

# Linux/Mac
lsof -ti:3000 | xargs kill -9
```

### Database Connection Error
1. Pastikan PostgreSQL running
2. Cek credentials di `.env`
3. Test connection:
   ```bash
   psql -U postgres -d jobsme
   ```

### Module Not Found
```bash
go mod download
go mod tidy
```

---

## 🚀 Next Steps

1. ✅ Setup local development
2. ✅ Test semua fitur
3. ✅ Push code ke GitHub
4. 🚀 Deploy ke Railway (lihat DEPLOYMENT.md)

---

## 💡 Tips

- Gunakan **Air** untuk hot reload saat development
- Test API endpoints dengan **Thunder Client**
- Backup database secara berkala
- Commit changes frequently

---

## 📚 Resources

- [Go Documentation](https://go.dev/doc/)
- [Fiber Framework](https://docs.gofiber.io/)
- [GORM Documentation](https://gorm.io/docs/)
- [PostgreSQL Tutorial](https://www.postgresqltutorial.com/)

---

**Happy Coding! 🎉**
