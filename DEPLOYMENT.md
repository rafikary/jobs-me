# 🚀 Panduan Deployment ke Railway

## Prerequisites
- Akun Railway (https://railway.app)
- Akun GitHub (untuk deploy dari repository)
- Git installed

---

## 📋 Step by Step Deployment

### 1. Setup Database PostgreSQL di Railway

1. **Login ke Railway**
   - Buka https://railway.app
   - Login dengan akun GitHub

2. **Buat Project Baru**
   - Click "New Project"
   - Pilih "Deploy PostgreSQL"
   - Database akan otomatis ter-provision

3. **Copy Database URL**
   - Click PostgreSQL service yang baru dibuat
   - Buka tab "Variables"
   - Copy value dari `DATABASE_URL`
   - Format: `postgresql://username:password@host:port/database`

---

### 2. Push Code ke GitHub

```bash
# Initialize git (jika belum)
git init

# Add all files
git add .

# Commit
git commit -m "Initial commit: Job Application Tracker"

# Add remote (ganti dengan URL repository Anda)
git remote add origin https://github.com/username/jobs-me.git

# Push ke GitHub
git push -u origin main
```

---

### 3. Deploy Aplikasi ke Railway

#### Opsi A: Deploy dari GitHub (Recommended)

1. **Kembali ke Railway Dashboard**
   - Click "New" → "GitHub Repo"
   - Authorize Railway untuk akses repository
   - Pilih repository `jobs-me`

2. **Railway akan otomatis:**
   - Detect Golang project
   - Build aplikasi
   - Deploy

3. **Tunggu Build Process**
   - Monitor di tab "Deployments"
   - Build biasanya memakan waktu 2-5 menit

#### Opsi B: Deploy dengan Railway CLI

```bash
# Install Railway CLI
npm install -g @railway/cli

# Login
railway login

# Link ke project
railway link

# Deploy
railway up
```

---

### 4. Setup Environment Variables

1. **Buka Service yang baru di-deploy**
2. **Klik tab "Variables"**
3. **Tambahkan variable berikut:**

   ```
   DATABASE_URL=postgresql://user:pass@host:port/db?sslmode=require
   PORT=3000
   ENV=production
   ```

   > ⚠️ **Penting**: Copy `DATABASE_URL` dari PostgreSQL service Anda

4. **Click "Add Variable" untuk save**

---

### 5. Generate Public Domain

1. **Buka service aplikasi**
2. **Klik tab "Settings"**
3. **Scroll ke bagian "Networking"**
4. **Click "Generate Domain"**
5. **Railway akan berikan URL seperti:**
   ```
   https://jobs-me-production-xxxx.up.railway.app
   ```

6. **Buka URL tersebut di browser** ✅

---

### 6. Custom Domain (Opsional)

Jika Anda punya domain sendiri:

1. **Di Railway Settings → Networking**
2. **Click "Custom Domain"**
3. **Masukkan domain Anda** (misal: `jobs.yourdomain.com`)
4. **Copy CNAME record yang diberikan**
5. **Tambahkan CNAME di DNS provider Anda:**
   ```
   CNAME: jobs → your-app.up.railway.app
   ```
6. **Tunggu DNS propagation** (5-30 menit)

---

## 🔧 Troubleshooting

### Build Failed

**Cek logs di Railway:**
```
Tab "Deployments" → Click pada deployment → "View Logs"
```

**Common issues:**
- Missing dependencies: Run `go mod tidy` locally dan push lagi
- Wrong Go version: Pastikan go.mod menggunakan Go 1.22+

### Database Connection Error

**Cek:**
1. `DATABASE_URL` sudah benar di Variables
2. PostgreSQL service running
3. Connection string format: `postgresql://...?sslmode=require`

### Application Crash

**Cek logs:**
```bash
railway logs
```

**Common fixes:**
- Restart service: Tab "Settings" → "Restart"
- Re-deploy: Push commit baru ke GitHub

---

## 📊 Monitoring

### View Logs
```bash
# Dengan Railway CLI
railway logs

# Atau di Dashboard
Tab "Deployments" → "View Logs"
```

### Metrics
- Lihat di tab "Metrics"
- Monitor CPU, Memory, Network usage

---

## 🔄 Update Aplikasi

Setiap kali push ke GitHub, Railway akan otomatis:
1. Detect changes
2. Rebuild aplikasi
3. Deploy ulang
4. Zero-downtime deployment

```bash
# Buat perubahan
git add .
git commit -m "Update feature"
git push

# Railway akan auto-deploy dalam 2-5 menit
```

---

## 💰 Biaya

**Railway Free Tier:**
- $5 gratis per bulan (credit)
- Cukup untuk aplikasi kecil-menengah
- PostgreSQL included

**Estimasi usage:**
- Small app: ~$2-3/bulan
- Medium traffic: ~$5-10/bulan

---

## 🔐 Security Checklist

- [x] `.env` tidak di-commit ke Git
- [x] `DATABASE_URL` menggunakan SSL (`sslmode=require`)
- [x] Environment variables di Railway, bukan hardcoded
- [x] `.gitignore` sudah setup dengan benar

---

## 📞 Support

Jika ada masalah:
1. Cek Railway documentation: https://docs.railway.app
2. Railway Discord: https://discord.gg/railway
3. Cek logs untuk error messages

---

## ✅ Checklist Deployment

- [ ] PostgreSQL database sudah dibuat di Railway
- [ ] `DATABASE_URL` sudah di-copy
- [ ] Code sudah di-push ke GitHub
- [ ] Service sudah di-deploy dari GitHub repo
- [ ] Environment variables sudah di-set
- [ ] Public domain sudah di-generate
- [ ] Aplikasi bisa diakses via browser
- [ ] CRUD operations berfungsi normal

---

🎉 **Selamat! Aplikasi Anda sudah live di Railway!**
