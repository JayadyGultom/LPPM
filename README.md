# 📚 LPPM REST API

Sebuah RESTful API untuk mengelola data LPPM yang dibangun menggunakan bahasa Go dan mengikuti prinsip **Clean Architecture**.

> 📝 *Dokumen ini disusun sebagai bagian dari tugas mata kuliah **Integrative Programming and Technology***.

---

## 🔧 Fitur Utama

* **Manajemen Profil**: Akses baca terhadap data visi, misi, dan struktur organisasi
* **Manajemen Penelitian**: Operasi CRUD (Create, Read, Update, Delete) penuh untuk berbagai kategori penelitian
* **Manajemen Pengabdian kepada Masyarakat (PKM)**: CRUD penuh untuk kegiatan pengabdian
* **Manajemen Jurnal**: CRUD penuh untuk berbagai jenis jurnal
* **Manajemen HKI (Hak Kekayaan Intelektual)**: CRUD penuh untuk data HKI dosen dan mahasiswa

---

## 🗃️ Struktur Tabel Database

### 🔹 Profil (Hanya Baca)

* `lppm_profile_visimisi` – Visi & Misi
* `lppm_profile_so` – Struktur Organisasi

### 🔹 Penelitian (CRUD Lengkap)

* `lppm_penelitian_bame`, `hpp`, `lp`, `p3`, `rrp`, `sk_r`, `stp`, `tcr`

### 🔹 PKM (CRUD Lengkap)

* `lppm_pkm_bame`, `hpp`, `lp`, `p3`, `rrp`, `sk_r`, `stp`, `tcr`

### 🔹 Jurnal (CRUD Lengkap)

* `lppm_jurnal_jk`, `js`, `kiat`, `tajb`, `teknois`, `tmjb`

### 🔹 HKI (CRUD Lengkap)

* `lppm_hki_dosen` – HKI Dosen
* `lppm_hki_mhs` – HKI Mahasiswa

---

## 🌐 Contoh Endpoint API

Struktur endpoint mengikuti pola standar REST, contohnya:

### 🔸 Penelitian BAME

* `GET /penelitian/bame` – Ambil semua data
* `POST /penelitian/bame` – Tambah data baru
* `GET /penelitian/bame/{id}` – Ambil data berdasarkan ID
* `PUT /penelitian/bame/{id}` – Perbarui data
* `DELETE /penelitian/bame/{id}` – Hapus data

Struktur endpoint serupa untuk PKM (`/pkm/...`), Jurnal (`/jurnal/...`), dan HKI (`/hki/...`)

---

## 💡 Struktur Data

```json
{
  "id": 1,
  "judul": "Judul Contoh",
  "konten": "Isi konten contoh"
}
```

---

## 🚀 Cara Menjalankan Aplikasi

### 1. 📦 Persyaratan

* Golang versi 1.19+
* MySQL versi 5.7+

### 2. 🗄️ Setup Database

* Buat database `lppm`
* Import file `lppm (4).sql`

### 3. ⚙️ Konfigurasi Environment

```bash
export DB_DSN="username:password@tcp(localhost:3306)/lppm?charset=utf8mb4&parseTime=True&loc=Local"
```

### 4. 🔧 Install Dependensi

```bash
go mod tidy
```

### 5. ▶️ Jalankan Aplikasi

```bash
go run cmd/main.go
```

Akses aplikasi di `http://localhost:8081`

---

## 🏗️ Arsitektur (Clean Architecture)

* **Entity Layer** – Definisi struktur data
* **Repository Layer** – Akses data dan database
* **Use Case Layer** – Logika aplikasi
* **Interface Layer** – Routing dan handler
* **Infrastructure Layer** – Koneksi database dan eksternal service

---

## 🧪 Contoh Penggunaan API

### Ambil Semua Visi Misi

```bash
curl -X GET http://localhost:8081/visimisi
```

### Tambah Data Penelitian

```bash
curl -X POST http://localhost:8081/penelitian/bame \
  -H "Content-Type: application/json" \
  -d '{"judul": "Penelitian Baru", "konten": "Isi konten"}'
```

### Update Data

```bash
curl -X PUT http://localhost:8081/penelitian/bame/1 \
  -H "Content-Type: application/json" \
  -d '{"judul": "Update", "konten": "Konten baru"}'
```

### Hapus Data

```bash
curl -X DELETE http://localhost:8081/penelitian/bame/1
```

---

## 📄 Lisensi

Proyek ini dibuat untuk **tujuan edukasi**.
Lisensi: **MIT License**

---

## 👨‍👩‍👧‍👦 Anggota Kelompok

* **Jayady Managam Gultom**
* **Agung Dwi Pradipta**
* **Firman Sulaiman**
* **Adrian Wahyuda Aditya P**
* **Ahmad Hardiansyah**
