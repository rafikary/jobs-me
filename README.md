# Job Application Tracker

Sistem tracking lamaran kerja yang dibuat dengan Golang. Aplikasi ini membantu mencatat dan memonitor status lamaran kerja ke berbagai perusahaan.

## 🚀 Features

- ✅ CRUD (Create, Read, Update, Delete) lamaran kerja
- ✅ Track status: Belum Apply, Applied, Interview HRD, Interview User, Ditolak, Diterima
- ✅ Multi-platform: LinkedIn, Glints, Jobstreet, Company Site, dll
- ✅ Dashboard dengan statistik real-time
- ✅ Filter berdasarkan status, platform, dan pencarian
- ✅ Simpan link lowongan dan catatan
- ✅ Responsive design (mobile-friendly)

## 🛠️ Tech Stack

- **Backend**: Go 1.22 + Fiber Framework
- **Database**: PostgreSQL (GORM)
- **Frontend**: HTML Templates + Tailwind CSS + Alpine.js + HTMX
- **Deployment**: Railway

## 📦 Installation

### Prerequisites
- Go 1.22 or higher
- PostgreSQL database

### Local Development

1. Clone repository:
```bash
git clone <repository-url>
cd jobs-me
```

2. Install dependencies:
```bash
go mod download
```

3. Setup environment variables:
```bash
cp .env.example .env
```

Edit `.env` file dengan database credentials Anda:
```env
DATABASE_URL=postgresql://username:password@localhost:5432/jobsme?sslmode=disable
PORT=3000
ENV=development
```

4. Run the application:
```bash
go run main.go
```

Aplikasi akan berjalan di `http://localhost:3000`

## 🚢 Deployment ke Railway

### 1. Setup Database di Railway

1. Buat project baru di [Railway.app](https://railway.app)
2. Tambahkan PostgreSQL service:
   - Click "New" → "Database" → "Add PostgreSQL"
3. Copy `DATABASE_URL` dari PostgreSQL variables

### 2. Deploy Aplikasi

**Opsi A: Deploy dari GitHub**

1. Push code ke GitHub repository
2. Di Railway, click "New" → "GitHub Repo"
3. Pilih repository Anda
4. Railway akan auto-detect Golang project

**Opsi B: Deploy dengan Railway CLI**

1. Install Railway CLI:
```bash
npm i -g @railway/cli
```

2. Login:
```bash
railway login
```

3. Initialize project:
```bash
railway init
```

4. Deploy:
```bash
railway up
```

### 3. Setup Environment Variables di Railway

1. Buka project di Railway dashboard
2. Masuk ke service yang baru dibuat
3. Klik tab "Variables"
4. Tambahkan:
   - `DATABASE_URL`: Copy dari PostgreSQL service
   - `PORT`: (optional, Railway auto-assign)
   - `ENV`: `production`

### 4. Generate Domain

1. Klik tab "Settings"
2. Scroll ke "Networking"
3. Click "Generate Domain"
4. Railway akan memberikan public URL

## 📁 Project Structure

```
jobs-me/
├── config/
│   └── database.go          # Database connection
├── handlers/
│   └── job_handler.go       # HTTP handlers
├── models/
│   └── job_application.go   # Database models
├── routes/
│   └── routes.go            # Route definitions
├── static/
│   └── app.js               # Frontend JavaScript
├── templates/
│   └── dashboard.html       # HTML templates
├── .env.example             # Environment template
├── .gitignore
├── go.mod                   # Go modules
├── main.go                  # Application entry point
├── railway.json             # Railway config
└── README.md
```

## 🎯 API Endpoints

### Web Routes
- `GET /` - Dashboard page

### API Routes
- `GET /api/jobs` - Get all jobs (with filters)
- `GET /api/jobs/:id` - Get single job
- `POST /api/jobs` - Create new job
- `PUT /api/jobs/:id` - Update job
- `DELETE /api/jobs/:id` - Delete job
- `GET /api/health` - Health check

## 📝 Usage

### Tambah Lamaran Baru
1. Click tombol "Tambah Lamaran" di header
2. Isi form dengan data lamaran:
   - Nama Perusahaan
   - Posisi yang dilamar
   - Platform (LinkedIn, Glints, Jobstreet, dll)
   - Status saat ini
   - Link lowongan (opsional)
   - Catatan (opsional)
3. Click "Simpan"

### Update Status Lamaran
1. Click icon edit (pensil) pada baris lamaran
2. Update status atau informasi lainnya
3. Click "Simpan"

### Filter dan Pencarian
- Gunakan field "Cari" untuk mencari berdasarkan nama perusahaan atau posisi
- Pilih status atau platform tertentu untuk filter
- Click tombol "Filter"

### Hapus Lamaran
- Click icon delete (trash) pada baris lamaran
- Confirm untuk menghapus

## 🔒 Security Notes

- Jangan commit file `.env` ke repository
- Gunakan PostgreSQL dengan SSL di production (`sslmode=require`)
- Update dependencies secara berkala

## 📊 Database Schema

```sql
job_applications (
    id SERIAL PRIMARY KEY,
    company VARCHAR NOT NULL,
    position VARCHAR NOT NULL,
    platform VARCHAR NOT NULL,
    status VARCHAR NOT NULL DEFAULT 'Belum Apply',
    job_link VARCHAR,
    notes TEXT,
    applied_date TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)
```

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 📄 License

MIT License

## 👤 Author

Dibuat dengan ❤️ untuk tracking lamaran kerja yang lebih efisien

---

**Happy Job Hunting! 🎯**
