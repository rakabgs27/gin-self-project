# ================================================
#  Makefile — gin-self-project
#  Tested on Windows (Git Bash / WSL) & Linux/Mac
# ================================================

APP_NAME  = gin-self-project
BUILD_DIR = ./bin
MAIN_PATH = ./cmd/server/main.go

# Deteksi OS
ifeq ($(OS),Windows_NT)
  BIN_EXT = .exe
  RM      = if exist $(BUILD_DIR) rmdir /s /q $(BUILD_DIR)
else
  BIN_EXT =
  RM      = rm -rf $(BUILD_DIR)
endif

.PHONY: all setup run run-air build test test-cover fmt lint check \
        tidy clean migrate docker-up docker-down docker-build help

# ================================================
#  DEFAULT
# ================================================

all: tidy build  ## Download deps lalu build

# ================================================
#  SETUP — buat semua folder & file sekaligus
# ================================================

setup:  ## Buat struktur folder, file awal, dan install semua deps
	@echo ""
	@echo "=> Membuat struktur folder..."
	@mkdir -p cmd/server
	@mkdir -p internal/handler
	@mkdir -p internal/service
	@mkdir -p internal/repository
	@mkdir -p internal/domain
	@mkdir -p config
	@mkdir -p migrations
	@mkdir -p pkg/response
	@echo "=> Membuat .env dari .env.example (jika belum ada)..."
	@if [ ! -f .env ]; then cp .env.example .env && echo "   .env dibuat"; else echo "   .env sudah ada, skip"; fi
	@echo "=> Membuat .air.toml (jika belum ada)..."
	@if [ ! -f .air.toml ]; then \
		printf 'root = "."\ntmp_dir = "tmp"\n\n[build]\n  cmd = "go build -o ./tmp/main$(BIN_EXT) $(MAIN_PATH)"\n  bin = "./tmp/main$(BIN_EXT)"\n  delay = 1000\n  exclude_dir = ["tmp","vendor","bin"]\n  include_ext = ["go","env"]\n  kill_delay = "0s"\n\n[log]\n  time = true\n\n[color]\n  main = "magenta"\n  watcher = "cyan"\n  build = "yellow"\n  runner = "green"\n\n[misc]\n  clean_on_exit = true\n' > .air.toml && echo "   .air.toml dibuat"; \
	else echo "   .air.toml sudah ada, skip"; fi
	@echo "=> Menginstall dependencies..."
	@go get github.com/gin-gonic/gin
	@go get gorm.io/gorm
	@go get gorm.io/driver/mysql
	@go get github.com/joho/godotenv
	@go mod tidy
	@echo "=> Menginstall tools (air)..."
	@go install github.com/air-verse/air@latest
	@echo ""
	@echo "======================================"
	@echo "  Setup selesai!"
	@echo "  1. Edit .env sesuai MySQL kamu"
	@echo "  2. Jalankan: make run"
	@echo "  3. Hot reload: make run-air"
	@echo "======================================"
	@echo ""

# ================================================
#  DEVELOPMENT
# ================================================

run:  ## Jalankan server (go run)
	@go run $(MAIN_PATH)

run-air:  ## Jalankan server dengan hot-reload (air)
	@air

build:  ## Build binary ke ./bin/
	@echo "Building $(APP_NAME)..."
	@go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)$(BIN_EXT) $(MAIN_PATH)
	@echo "Binary siap: $(BUILD_DIR)/$(APP_NAME)$(BIN_EXT)"

# ================================================
#  TESTING & QUALITY
# ================================================

test:  ## Jalankan semua unit test
	@go test ./... -v -race -cover

test-cover:  ## Test + laporan coverage HTML
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Laporan coverage: coverage.html"

fmt:  ## Format semua file Go
	@gofmt -w .

lint:  ## Jalankan linter (butuh golangci-lint)
	@golangci-lint run ./...

check: fmt lint test  ## fmt + lint + test sekaligus

tidy:  ## Download & rapikan dependencies
	@go mod tidy

# ================================================
#  DATABASE
# ================================================

migrate:  ## Jalankan SQL migration manual
	@echo "Menjalankan migration..."
	@mysql -u$(DB_USER) -p$(DB_PASS) $(DB_NAME) < migrations/001_create_users.sql
	@echo "Migration selesai"

# ================================================
#  DOCKER (opsional)
# ================================================

docker-up:  ## Jalankan MySQL via Docker Compose
	@docker compose up -d

docker-down:  ## Hentikan semua container
	@docker compose down

docker-build:  ## Build Docker image aplikasi
	@docker build -t $(APP_NAME) .

# ================================================
#  MISC
# ================================================

clean:  ## Hapus build artifacts dan tmp
	@$(RM)
	@rm -rf tmp coverage.out coverage.html
	@echo "Clean selesai"

help:  ## Tampilkan semua perintah yang tersedia
	@echo ""
	@echo "  $(APP_NAME) — Available Commands"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36mmake %-15s\033[0m %s\n", $$1, $$2}'
	@echo ""