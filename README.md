# 🚀 GIN SELF PROJECT

REST API self project menggunakan **Go + Gin + MySQL (GORM)** dengan arsitektur layered yang clean dan siap dikembangkan lebih lanjut.

---

## 🛠️ Tech Stack

| Layer | Library |
|---|---|
| HTTP Framework | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io) |
| Database | MySQL 8.x |
| Config / Env | [godotenv](https://github.com/joho/godotenv) |

---

## 📁 Struktur Project

```
gin-self-project/
├── cmd/
│   └── server/
│       └── main.go          # Entry point
├── internal/
│   ├── domain/              # Model & DTO (struct, tidak ada logic)
│   │   └── user.go
│   ├── repository/          # Query database (GORM)
│   │   └── user_repository.go
│   ├── service/             # Business logic
│   │   └── user_service.go
│   └── handler/             # HTTP handler & router (Gin)
│       ├── router.go
│       └── user_handler.go
├── config/
│   └── config.go            # Load env & koneksi database
├── migrations/
│   └── 001_create_users.sql # SQL migration manual
├── pkg/
│   └── response/            # Helper standard JSON response
│       └── response.go
├── .env.example             # Template environment variable
├── Makefile                 # Kumpulan perintah build & dev
└── README.md
```

---

## ⚡ Quick Start

### 1. Clone & masuk folder

```bash
git clone https://github.com/rakabgs27/gin-self-project.git
cd gin-self-project
```

### 2. Copy dan isi file env

```bash
cp .env.example .env
```

Edit `.env` sesuai konfigurasi MySQL kamu:

```env
APP_PORT=8080
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASS=
DB_NAME=mydb
```

### 3. Buat database di MySQL

```sql
CREATE DATABASE mydb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 4. Download dependencies

```bash
make tidy
# atau: go mod tidy
```

### 5. Jalankan server

```bash
make run
# atau: go run ./cmd/server/main.go
```

Server akan jalan di `http://localhost:8080`

> **AutoMigrate aktif** — tabel `users` akan otomatis dibuat saat server pertama kali dijalankan.

---

## 📡 API Endpoints

Base URL: `http://localhost:8080/api/v1`

| Method | Endpoint | Deskripsi |
|---|---|---|
| GET | `/ping` | Health check |
| GET | `/api/v1/users` | Ambil semua user |
| GET | `/api/v1/users/:id` | Ambil user by ID |
| POST | `/api/v1/users` | Buat user baru |
| PUT | `/api/v1/users/:id` | Update user |
| DELETE | `/api/v1/users/:id` | Hapus user |

### Contoh Request & Response

**POST /api/v1/users**

Request:
```json
{
  "name": "Raka Bagus",
  "email": "raka@example.com",
  "phone": "081234567890"
}
```

Response `201 Created`:
```json
{
  "success": true,
  "message": "User berhasil dibuat",
  "data": {
    "id": 1,
    "name": "Raka Bagus",
    "email": "raka@example.com",
    "phone": "081234567890",
    "created_at": "2026-04-28T11:00:00Z",
    "updated_at": "2026-04-28T11:00:00Z"
  }
}
```

**GET /api/v1/users**

Response `200 OK`:
```json
{
  "success": true,
  "message": "Berhasil mengambil data user",
  "data": [...]
}
```

**Response error (contoh: user tidak ditemukan)**
```json
{
  "success": false,
  "message": "user tidak ditemukan"
}
```

---

## 🧰 Makefile Commands

```bash
make help          # Tampilkan semua perintah
make run           # Jalankan server
make run-air       # Jalankan dengan hot-reload (butuh air)
make build         # Build binary ke ./bin/
make test          # Jalankan semua test
make test-cover    # Test + laporan coverage HTML
make fmt           # Format kode Go
make lint          # Jalankan linter (butuh golangci-lint)
make check         # fmt + lint + test sekaligus
make tidy          # go mod tidy
make clean         # Hapus build artifacts
make docker-up     # Jalankan MySQL via Docker
make docker-down   # Hentikan container Docker
```

---

## 🏗️ Arsitektur

```
Request → Handler → Service → Repository → Database
                ↑         ↑           ↑
           (Gin)    (Business     (GORM)
                      Logic)
```

- **Handler** — menerima HTTP request, validasi input, kirim response
- **Service** — business logic, tidak tahu apapun tentang HTTP atau DB
- **Repository** — satu-satunya yang boleh menyentuh database
- **Domain** — struct model dan DTO, tidak ada logic

---

## 🔧 Tools Opsional

| Tool | Fungsi | Install |
|---|---|---|
| [air](https://github.com/air-verse/air) | Hot reload saat development | `go install github.com/air-verse/air@latest` |
| [golangci-lint](https://golangci-lint.run) | Linter | `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` |

---

## 📝 Catatan Development

- Gunakan `make run` untuk development sehari-hari
- Tambah endpoint baru dengan urutan: `domain` → `repository` → `service` → `handler` → daftarkan di `router.go`
- Format kode sebelum commit: `make fmt`

---

## 📄 License

MIT License — bebas digunakan untuk belajar dan dikembangkan.