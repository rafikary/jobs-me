# 🚀 CARA PAKAI APLIKASI INI

## ❗ PENTING!

**Kamu tadi buka `ui-preview.html` - itu cuma PREVIEW statis!**

**Untuk pakai aplikasi ASLI dengan FORM yang bisa input data:**

---

## 🎯 Cara Run Aplikasi (2 Pilihan)

### PILIHAN 1: Deploy ke Railway (RECOMMENDED - Paling Gampang!)

1. **Code sudah di GitHub:** ✅ https://github.com/rafikary/jobs-me.git

2. **Setup Railway:**
   - Buka https://railway.app
   - Login dengan GitHub
   - New Project → Deploy PostgreSQL
   - Copy `DATABASE_URL`

3. **Deploy Aplikasi:**
   - New → GitHub Repo
   - Pilih repository `jobs-me`
   - Add Variable: `DATABASE_URL` (paste dari PostgreSQL)
   - Settings → Generate Domain
   - **DONE! Buka URL-nya!** 🎉

**Baca panduan lengkap:** [QUICK_DEPLOY.md](QUICK_DEPLOY.md)

---

### PILIHAN 2: Run Local (Butuh PostgreSQL)

#### Windows:

1. **Install PostgreSQL:**
   - Download: https://www.postgresql.org/download/windows/
   - Setelah install, buat database:
   ```sql
   CREATE DATABASE jobsme;
   ```

2. **Edit file `.env`:**
   ```env
   DATABASE_URL=postgresql://postgres:yourpassword@localhost:5432/jobsme?sslmode=disable
   PORT=3000
   ```

3. **Run aplikasi:**
   ```powershell
   go run main.go
   ```

4. **Buka browser:**
   ```
   http://localhost:3000
   ```

---

## 📱 CARA PAKAI (Setelah Aplikasi Running)

### 1️⃣ **Tambah Lamaran Baru:**
- **Desktop:** Click tombol merah melayang di kanan bawah (FAB)
- **Mobile:** Tap tombol **+** merah di kanan bawah
- Isi form:
  - Company (wajib)
  - Position (wajib)
  - Platform (LinkedIn/Glints/dll)
  - Status (Pending/Applied/dll)
  - Job Link (opsional)
  - Notes (opsional)
- Click **Save**

### 2️⃣ **Edit Lamaran:**
- **Desktop:** Click icon **pensil** di kanan baris
- **Mobile:** Tap icon **pensil** di card

### 3️⃣ **Delete Lamaran:**
- **Desktop:** Click icon **sampah** merah
- **Mobile:** Tap icon **sampah**

### 4️⃣ **Filter & Search:**
- Gunakan form filter di atas table
- Search by company atau position
- Filter by status atau platform
- Click tombol **Filter**

---

## 🎨 FITUR DESIGN BARU!

### ✨ Yang Baru:
- **Dark Theme** - Background gelap dengan gradient unik
- **Floating Action Button (FAB)** - Tombol merah bulat melayang untuk tambah data
- **Mobile Responsive** - Card view di mobile, table di desktop
- **Animated Background** - Pattern yang bergerak halus
- **Unique Colors** - Tidak generic, punya personality
- **Smooth Animations** - Hover effects & transitions

### 📱 Mobile Features:
- 2 kolom grid untuk stats
- Card-based list view (bukan table)
- Touch-friendly buttons
- Optimized spacing
- Easy swipe & scroll

### 🖥️ Desktop Features:
- 6 kolom stats grid
- Full table dengan semua kolom
- Hover effects
- Spacious layout

---

## 🎯 STATUS WARNA:

- ⏳ **Pending** (Belum Apply) - Kuning
- 📤 **Applied** - Hijau
- 🤝 **Interview HRD** - Ungu
- 💬 **Interview User** - Biru
- ❌ **Rejected** (Ditolak) - Merah
- 🎉 **Accepted** (Diterima) - Cyan

---

## 📸 PREVIEW vs REAL APP

### `ui-preview.html` (PREVIEW):
- ❌ Tidak bisa input data
- ❌ Tidak connect ke database
- ❌ Cuma tampilan statis
- ✅ Untuk lihat design saja

### `http://localhost:3000` atau Railway URL (REAL):
- ✅ **Bisa tambah data dengan FAB!**
- ✅ **Form lengkap & working!**
- ✅ Edit, delete, filter works!
- ✅ Data tersimpan di database
- ✅ **INI YANG HARUS KAMU PAKAI!**

---

## 🔴 FLOATING ACTION BUTTON (FAB)

Tombol merah bulat di kanan bawah adalah **tombol utama untuk tambah data!**

**Kenapa FAB?**
- Selalu visible (tidak perlu scroll)
- Mobile-friendly (mudah dijangkau jempol)
- Animated (ada pulse effect)
- Clear action (+ icon = tambah)
- Modern UX pattern

**Cara pakai:**
1. Click/Tap FAB merah
2. Modal form muncul
3. Isi data
4. Save
5. Data muncul di list!

---

## 🎨 DESIGN PHILOSOPHY

**Kenapa Dark Theme?**
- Less eye strain (tidak silau)
- Modern & professional
- Fokus ke content
- Stand out (tidak generic)

**Kenapa Tidak "AI-Looking"?**
- Custom color palette (bukan template)
- Asymmetric elements
- Playful animations
- Personality (emoji, warna unik)
- Human touch

---

## 🚀 NEXT STEPS

### ✅ Sudah Selesai:
- [x] Code di GitHub
- [x] Dark theme unik
- [x] FAB untuk tambah data
- [x] Mobile responsive
- [x] Form lengkap & works

### 🎯 Sekarang Kamu:
1. **Deploy ke Railway** (15 menit)
   - Ikuti: [QUICK_DEPLOY.md](QUICK_DEPLOY.md)
   - Atau run local (butuh PostgreSQL)

2. **Mulai Pakai:**
   - Click FAB merah
   - Tambah lamaran
   - Track progress
   - Impress recruiters! 💼

---

## 💡 TIPS

### Desktop:
- Hover over cards untuk effect
- Use keyboard shortcuts (Tab untuk navigate)
- Click outside modal untuk close

### Mobile:
- Swipe untuk scroll horizontal table
- Tap & hold untuk select text
- Use native browser zoom

---

## 🆘 TROUBLESHOOTING

### "Tidak ada tombol tambah data!"
→ Kamu buka `ui-preview.html`, bukan aplikasi asli!
→ Run `go run main.go` atau deploy ke Railway

### "Form tidak bisa save!"
→ Pastikan database sudah setup
→ Cek `DATABASE_URL` di `.env` atau Railway variables

### "Tampilan berantakan di mobile!"
→ Refresh browser (Ctrl/Cmd + Shift + R)
→ Clear cache

### "FAB tidak muncul!"
→ Refresh page
→ Check console untuk errors (F12)

---

## 📞 Need Help?

- Read: [README.md](README.md)
- Deploy Guide: [QUICK_DEPLOY.md](QUICK_DEPLOY.md)
- Local Setup: [QUICKSTART.md](QUICKSTART.md)

---

**🎉 Sekarang GO! Deploy atau run local, terus pakai tombol merah untuk input data!**

**Design udah gak AI banget kan? Dark, unique, ada personality! 😎**
